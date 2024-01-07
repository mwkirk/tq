package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
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
	log.Printf("worker registered: %v, id: %v", rr.GetRegistered(), rr.GetId())

	dr, err := c.Deregister(ctx, &pbuf.DeregisterRequest{})
	if err != nil {
		log.Fatalf("failed to degister: %v", err)
	}
	log.Printf("worker deregistered: %v", dr.GetRegistered())
}
