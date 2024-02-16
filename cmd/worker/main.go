package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"tq/internal/model"
	"tq/pb"
)

var (
	addr  = flag.String("addr", "localhost:8000", "server address")
	label = "foo"
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTqWorkerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	rr, err := c.Register(ctx, &pb.RegisterRequest{Label: label})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	log.Printf("worker registered: %v, id: %v", rr.Registered, rr.Id)

	done := make(chan struct{})
	updates := make(chan pb.JobStatus)
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	go func() {
		w := model.Worker{
			Registered:  rr.Registered,
			Id:          model.WorkerId(rr.Id),
			Label:       label,
			WorkerState: pb.WorkerState_WORKER_STATE_AVAILABLE,
		}

		// fixme: this needs to be non-blocking
		updates <- pb.JobStatus{}

		for {
			var j pb.JobStatus

			select {
			case <-done:
				log.Printf("exiting status goroutine")
				return
			case j = <-updates:
				log.Printf("job status updated: %v", j)
			case <-ticker.C:
				sr, err := c.Status(ctx, &pb.StatusRequest{
					Id:          w.Id.String(),
					WorkerState: w.WorkerState,
					JobStatus:   &j,
				})

				if err != nil {
					log.Printf("error received from status request: %v", err)
				} else {
					log.Printf("handleStatusResponse")
					err := handleStatusResponse(sr, &w, updates)
					if err != nil {
						log.Printf("%s", err)
					}
				}
			}
		}
	}()

	<-done
	dr, err := c.Deregister(ctx, &pb.DeregisterRequest{})
	if err != nil {
		log.Fatalf("failed to degister: %v", err)
	}
	log.Printf("worker deregistered: %v", dr.GetRegistered())
}

func handleStatusResponse(sr *pb.StatusResponse, w *model.Worker, updates chan<- pb.JobStatus) error {
	switch sr.JobControl {
	case pb.JobControl_JOB_CONTROL_NONE:
		log.Printf("no job available")
	case pb.JobControl_JOB_CONTROL_CONTINUE:
		log.Printf("continue current job")
	case pb.JobControl_JOB_CONTROL_NEW:
		log.Printf("new job")
		err := startJob(sr.Job, updates)
		if err != nil {
			return err
		}
		w.WorkerState = pb.WorkerState_WORKER_STATE_WORKING
	case pb.JobControl_JOB_CONTROL_CANCEL:
		log.Printf("cancel current job")
		w.WorkerState = pb.WorkerState_WORKER_STATE_AVAILABLE
	default:
		return fmt.Errorf("received unexpect job control message")
	}

	return nil
}

func startJob(job *pb.Job, updates chan<- pb.JobStatus) error {
	// guard
	if job == nil {
		return fmt.Errorf("no job data")
	}

	switch job.Kind {
	case pb.JobKind_JOB_KIND_NULL:
		return fmt.Errorf("new null job received")
	case pb.JobKind_JOB_KIND_TEST:
		log.Printf("test job received")
	case pb.JobKind_JOB_KIND_SLEEP:
		workerJob := newWorkerSleepJob(job, updates)
		workerJob.run()
	case pb.JobKind_JOB_KIND_FFMPEG:
		log.Printf("FFMPEG job received")
	default:
		return fmt.Errorf("unexpected job kind")
	}

	return nil
}
