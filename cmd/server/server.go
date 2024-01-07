package main

import (
	"context"
	"tq/pbuf"
)

type server struct {
	pbuf.UnimplementedTqServer
}

func NewServer() pbuf.TqServer {
	s := &server{}
	return s
}

func (s *server) Register(ctx context.Context, request *pbuf.RegisterRequest) (*pbuf.RegisterResponse, error) {
	return &pbuf.RegisterResponse{
		Registered: true,
		Id:         "fakeId",
	}, nil
}

func (s *server) Deregister(ctx context.Context, request *pbuf.DeregisterRequest) (*pbuf.DeregisterResponse, error) {
	return &pbuf.DeregisterResponse{Registered: false}, nil
}

func (s *server) Status(ctx context.Context, request *pbuf.StatusRequest) (*pbuf.StatusResponse, error) {
	// TODO implement me
	panic("implement me")
}
