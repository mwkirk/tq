package main

import (
	"context"
	"fmt"
	"tq/internal/model"
	"tq/pb"
)

type server struct {
	pb.UnimplementedTqWorkerServer
	pb.UnimplementedTqJobServer
	orc QueueOrchestrator
}

func newServer(orc QueueOrchestrator) *server {
	s := &server{
		orc: orc,
	}
	return s
}

// ------------------------------------------------------------------
// TqWorker
// ------------------------------------------------------------------

func (s *server) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := s.orc.Register(request.Label)
	if err != nil {
		return nil, fmt.Errorf("failed to register worker: %w", err)
	}

	return &pb.RegisterResponse{
		Registered: true,
		Id:         id.String(),
	}, nil
}

func (s *server) Deregister(ctx context.Context, request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	err := s.orc.Deregister(model.WorkerId(request.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to deregister worker: %w", err)
	}
	return &pb.DeregisterResponse{Registered: false}, nil
}

func (s *server) Status(ctx context.Context, request *pb.StatusRequest) (*pb.StatusResponse, error) {
	sr, err := s.orc.Status(model.WorkerId(request.Id), request.WorkerState, request.JobStatus)
	if err != nil {
		return nil, fmt.Errorf("status update failed for worker: %w", err)
	}
	return &sr, nil
}

// ------------------------------------------------------------------
// TqJob
// ------------------------------------------------------------------

func (s *server) Submit(ctx context.Context, request *pb.SubmitRequest) (*pb.SubmitResponse, error) {
	err := s.orc.Submit(request.Job)
	if err != nil {
		return &pb.SubmitResponse{Accepted: false}, err
	}
	return &pb.SubmitResponse{Accepted: true, JobNum: request.Job.Num}, nil
}
