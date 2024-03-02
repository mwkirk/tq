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
	wq := container.NewSliceQueue[*pb.JobSpec]()
	rq := container.NewSliceQueue[*pb.JobSpec]()
	dq := container.NewSliceQueue[*pb.JobSpec]()
	aws := container.NewSimpleMapStore[model.JobNumber, model.WorkerId]()
	jobMgr := NewSimpleJobMgr(wq, rq, dq, aws)
	orc := NewSimpleQueueOrchestrator(workerMgr, jobMgr)

	srv := grpc.NewServer()
	provider := newServer(orc)
	pb.RegisterTqWorkerServer(srv, provider)
	pb.RegisterTqJobServer(srv, provider)

	log.Printf("started server on %s", lis.Addr().String())
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("server exited with: %v", err)
	}
}
