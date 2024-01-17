package main

import (
	"context"
	"fmt"
	"tq/internal/model"
	"tq/pbuf"
)

type server struct {
	pbuf.UnimplementedTqServer
	mgr WorkerMgr
}

func NewServer(mgr WorkerMgr) *server {
	s := &server{
		mgr: mgr,
	}
	return s
}

func (s *server) Register(ctx context.Context, request *pbuf.RegisterRequest) (*pbuf.RegisterResponse, error) {
	id, err := s.mgr.Register(request.Label)
	if err != nil {
		return nil, fmt.Errorf("failed to register worker: %w", err)
	}

	return &pbuf.RegisterResponse{
		Registered: true,
		Id:         id.String(),
	}, nil
}

func (s *server) Deregister(ctx context.Context, request *pbuf.DeregisterRequest) (*pbuf.DeregisterResponse, error) {
	err := s.mgr.Deregister(model.WorkerId(request.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to deregister worker: %w", err)
	}
	return &pbuf.DeregisterResponse{Registered: false}, nil
}

func (s *server) Status(ctx context.Context, request *pbuf.StatusRequest) (*pbuf.StatusResponse, error) {
	err := s.mgr.Status(model.WorkerId(request.Id), request.WorkerState, request.JobStatus)
	if err != nil {
		return nil, fmt.Errorf("status update failed for worker: %w", err)
	}

	return &pbuf.StatusResponse{
		JobControl: pbuf.JobControl_JOB_CONTROL_NONE,
	}, nil
}
