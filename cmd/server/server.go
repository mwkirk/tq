package main

import (
	"context"
	"fmt"
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
	result, err := s.orc.Register(request.Options)
	if err != nil {
		err = fmt.Errorf("failed to register worker: %w", err)
	}

	return &pb.RegisterResponse{Result: result}, err
}

func (s *server) Deregister(ctx context.Context, request *pb.DeregisterRequest) (*pb.DeregisterResponse, error) {
	result, err := s.orc.Deregister(request.Options)
	if err != nil {
		err = fmt.Errorf("failed to deregister worker: %w", err)
	}
	return &pb.DeregisterResponse{Result: result}, err
}

func (s *server) Status(ctx context.Context, request *pb.StatusRequest) (*pb.StatusResponse, error) {
	result, err := s.orc.Status(request.Options)
	if err != nil {
		err = fmt.Errorf("status update failed for worker: %w", err)
	}
	return &pb.StatusResponse{Result: result}, err
}

// ------------------------------------------------------------------
// TqJob
// ------------------------------------------------------------------

func (s *server) Submit(ctx context.Context, request *pb.SubmitRequest) (*pb.SubmitResponse, error) {
	result, err := s.orc.Submit(request.Options)
	if err != nil {
		err = fmt.Errorf("job submission failed: %w", err)
	}
	return &pb.SubmitResponse{Result: result}, err
}

func (s *server) Cancel(ctx context.Context, request *pb.CancelRequest) (*pb.CancelResponse, error) {
	result, err := s.orc.Cancel(request.Options)
	if err != nil {
		err = fmt.Errorf("job cancellation failed: %w", err)
	}
	return &pb.CancelResponse{Result: result}, err
}

func (s *server) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	result, err := s.orc.List(request.Options)
	if err != nil {
		err = fmt.Errorf("job list failed: %w", err)
	}
	return &pb.ListResponse{Result: result}, err
}
