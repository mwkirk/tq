package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os/signal"
	"syscall"
	"time"
	"tq/internal"
	"tq/internal/model"
	"tq/pb"
)

var (
	addr              = flag.String("addr", "localhost:8000", "server address")
	label             = "foo"
	heartbeatInterval = 2 * time.Second
	timeout           = 30 * time.Second
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTqWorkerClient(conn)
	pctx := context.Background()

	// Register
	regCtx, regCancel := context.WithTimeout(pctx, timeout)
	defer regCancel()
	rr, err := c.Register(regCtx, &pb.RegisterRequest{Options: &pb.RegisterOptions{Label: label}})
	if err != nil {
		// todo: need retry mechanism
		log.Fatalf("failed to connect: %v", err)
	}
	log.Printf("worker registered: %v, id: %v", rr.Result.Registered, rr.Result.WorkerId)

	// Update status
	writeUpdates, readUpdates := internal.MakeNonblockingChanPair[*pb.JobStatus]()
	heartbeat := time.NewTicker(heartbeatInterval)
	defer heartbeat.Stop()
	sigCtx, stop := signal.NotifyContext(pctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	statusLoopCtx, statusLoopCancel := context.WithCancel(sigCtx)
	defer statusLoopCancel()

	go func() {
		currStatus := &pb.JobStatus{}
		w := &worker{
			Worker: model.Worker{
				Registered:  rr.Result.Registered,
				Id:          model.WorkerId(rr.Result.WorkerId),
				Label:       label,
				WorkerState: pb.WorkerState_WORKER_STATE_AVAILABLE,
			},
		}

		writeUpdates <- &pb.JobStatus{}

		for {
			select {
			case <-statusLoopCtx.Done():
				log.Printf("exiting status goroutine")
				return
			case <-heartbeat.C:
				// On each heartbeat, update the server with the current status.
				// This will be the status quo if no job is running or no new
				// progress from a job has been posted. If a job is running, there
				// could be several updates queued.

				updates := []*pb.JobStatus{}

			loop:
				for {
					select {
					case currStatus = <-readUpdates:
						log.Printf("job status updated: %s", currStatus.LogString())
						updates = append(updates, currStatus)
					default:
						break loop
					}
				}

				if len(updates) == 0 {
					updates = append(updates, currStatus)
				}

				// If job completes for any reason, set worker's state to available
				last := updates[len(updates)-1]
				switch last.JobState {
				case pb.JobState_JOB_STATE_DONE_OK:
					fallthrough
				case pb.JobState_JOB_STATE_DONE_ERR:
					fallthrough
				case pb.JobState_JOB_STATE_DONE_CANCEL:
					w.WorkerState = pb.WorkerState_WORKER_STATE_AVAILABLE
					w.job = nil
				}

				statusCtx, statusCancel := context.WithTimeout(statusLoopCtx, timeout)
				defer statusCancel()
				sr, err := c.Status(statusCtx, &pb.StatusRequest{
					Options: &pb.StatusOptions{
						WorkerId:    w.Id.String(),
						WorkerState: w.WorkerState,
						JobStatus:   updates,
					}})

				if err != nil {
					log.Printf("error received from status request: %v", err)
				} else {
					err := handleStatusResponse(statusLoopCtx, sr, w, writeUpdates)
					if err != nil {
						log.Printf("%s", err)
					}
				}
			}
		}
	}()

	<-statusLoopCtx.Done()

	// Deregister
	deregCtx, deregCancel := context.WithTimeout(pctx, timeout)
	defer deregCancel()
	dr, err := c.Deregister(deregCtx,
		&pb.DeregisterRequest{Options: &pb.DeregisterOptions{WorkerId: rr.Result.WorkerId}})
	if err != nil {
		log.Fatalf("failed to degister: %v", err)
	}
	log.Printf("worker deregistered: %v", dr.Result.Deregistered)
}

func handleStatusResponse(ctx context.Context, sr *pb.StatusResponse, w *worker, updates chan<- *pb.JobStatus) error {
	switch sr.Result.JobControl {
	case pb.JobControl_JOB_CONTROL_NONE:
		// log.Printf("no job available")
	case pb.JobControl_JOB_CONTROL_CONTINUE:
		// log.Printf("continue current job")
	case pb.JobControl_JOB_CONTROL_NEW:
		log.Printf("starting new job")
		err := w.startJob(ctx, sr.Result.Job, updates)
		if err != nil {
			return err
		}
		w.WorkerState = pb.WorkerState_WORKER_STATE_WORKING
	case pb.JobControl_JOB_CONTROL_CANCEL:
		log.Printf("canceling current job")
		w.job.cancel()
	default:
		return fmt.Errorf("received unexpected job control message")
	}

	return nil
}
