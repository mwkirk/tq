package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	"tq/internal/model"
	"tq/pbuf"
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
	c := pbuf.NewTqClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	rr, err := c.Register(ctx, &pbuf.RegisterRequest{Label: label})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	log.Printf("worker registered: %v, id: %v", rr.Registered, rr.Id)

	done := make(chan struct{})
	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()

	go func() {
		w := model.Worker{
			Registered:  true,
			Id:          model.WorkerId(rr.Id),
			Label:       label,
			WorkerState: pbuf.WorkerState_WORKER_STATE_AVAILABLE,
		}

		for {
			select {
			case <-done:
				log.Printf("exiting status goroutine")
				return
			case <-ticker.C:
				log.Printf("tick")
				sr, err := c.Status(ctx, &pbuf.StatusRequest{
					Id:          w.Id.String(),
					WorkerState: w.WorkerState,
					JobStatus:   &w.JobStatus,
				})

				if err != nil {
					log.Printf("error received from status request: %v", err)
				} else {
					log.Printf("status response, job control = %v", sr.JobControl)
				}
			}
		}
	}()

	<-done
	dr, err := c.Deregister(ctx, &pbuf.DeregisterRequest{})
	if err != nil {
		log.Fatalf("failed to degister: %v", err)
	}
	log.Printf("worker deregistered: %v", dr.GetRegistered())
}
