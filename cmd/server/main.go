package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"tq/internal/container"
	"tq/internal/model"
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

	ws := container.NewSimpleMapStore[model.WorkerId, *model.Worker]()
	mgr := NewSimpleWorkerMgr(&ws)
	srv := grpc.NewServer()
	pbuf.RegisterTqServer(srv, NewServer(mgr))

	log.Printf("started server on %s", lis.Addr().String())
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("server exited with: %v", err)
	}
}
