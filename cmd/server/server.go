package main

import (
	"context"
	"fmt"
	"tq/internal/model"
	"tq/pbuf"
)

type server struct {
	pbuf.UnimplementedTqServer
	orc QueueOrchestrator
}

func newServer(orc QueueOrchestrator) *server {
	s := &server{
		orc: orc,
	}
	return s
}

func (s *server) Register(ctx context.Context, request *pbuf.RegisterRequest) (*pbuf.RegisterResponse, error) {
	id, err := s.orc.Register(request.Label)
	if err != nil {
		return nil, fmt.Errorf("failed to register worker: %w", err)
	}

	return &pbuf.RegisterResponse{
		Registered: true,
		Id:         id.String(),
	}, nil
}

func (s *server) Deregister(ctx context.Context, request *pbuf.DeregisterRequest) (*pbuf.DeregisterResponse, error) {
	err := s.orc.Deregister(model.WorkerId(request.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to deregister worker: %w", err)
	}
	return &pbuf.DeregisterResponse{Registered: false}, nil
}

func (s *server) Status(ctx context.Context, request *pbuf.StatusRequest) (*pbuf.StatusResponse, error) {
	sr, err := s.orc.Status(model.WorkerId(request.Id), request.WorkerState, request.JobStatus)
	if err != nil {
		return nil, fmt.Errorf("status update failed for worker: %w", err)
	}

	return &pbuf.StatusResponse{
		JobControl: sr.JobControl,
		Job: &pbuf.Job{
			Kind:  sr.Kind,
			Num:   sr.Num,
			Name:  sr.Name,
			Parms: sr.Parms,
		},
	}, nil
}
