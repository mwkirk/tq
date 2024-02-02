package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"tq/internal/container"
	"tq/internal/model"
	"tq/pb"
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

	// wire up dependencies
	ws := container.NewSimpleMapStore[model.WorkerId, *model.Worker]()
	workerMgr := NewSimpleWorkerMgr(ws)
	wq := container.NewSliceQueue[*model.Job]()
	rq := container.NewSliceQueue[*model.Job]()
	dq := container.NewSliceQueue[*model.Job]()
	jobMgr := NewSimpleJobMgr(wq, rq, dq)
	orc := NewSimpleQueueOrchestrator(workerMgr, jobMgr)

	srv := grpc.NewServer()
	pb.RegisterTqWorkerServer(srv, newServer(orc))

	log.Printf("started server on %s", lis.Addr().String())
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("server exited with: %v", err)
	}
}
