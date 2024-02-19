package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"tq/internal"
	"tq/internal/model"
	"tq/pb"
)

var (
	addr              = flag.String("addr", "localhost:8000", "server address")
	label             = "foo"
	heartbeatInterval = 2 * time.Second
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

	writeUpdates, readUpdates := internal.MakeNonblockingChanPair[*pb.JobStatus]()
	done := make(chan struct{})
	heartbeat := time.NewTicker(heartbeatInterval)
	defer heartbeat.Stop()

	go func() {
		w := model.Worker{
			Registered:  rr.Registered,
			Id:          model.WorkerId(rr.Id),
			Label:       label,
			WorkerState: pb.WorkerState_WORKER_STATE_AVAILABLE,
		}

		writeUpdates <- &pb.JobStatus{}

		for {
			var currStatus *pb.JobStatus

			select {
			case <-done:
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

				sr, err := c.Status(ctx, &pb.StatusRequest{
					Id:          w.Id.String(),
					WorkerState: w.WorkerState,
					JobStatus:   updates,
				})

				if err != nil {
					log.Printf("error received from status request: %v", err)
				} else {
					err := handleStatusResponse(sr, &w, writeUpdates)
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
