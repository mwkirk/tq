package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"tq/pbuf"
)

const (
	host = "localhost"
	port = "8000"
)

func main() {
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pbuf.RegisterTqServer(srv, NewServer())

	log.Printf("started server on %s", lis.Addr().String())
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("server exited with: %v", err)
	}
}
