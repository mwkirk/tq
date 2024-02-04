package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"tq/internal/container"
	"tq/internal/model"
	"tq/pb"
)

const fixtureId = "e0df68b9-3d33-4262-8af6-71516108ea2d"

func testingServer(ctx context.Context) (pb.TqWorkerClient, func()) {
	bufSize := 1024 * 1024
	lis := bufconn.Listen(bufSize)

	// wire up dependencies
	ws := container.NewSimpleMapStore[model.WorkerId, *model.Worker]()
	ws.Add(fixtureId, &model.Worker{
		Registered:  true,
		Id:          fixtureId,
		Label:       "fixture worker",
		WorkerState: 0,
	})
	workerMgr := NewSimpleWorkerMgr(ws)
	wq := container.NewSliceQueue[*pb.Job]()
	rq := container.NewSliceQueue[*pb.Job]()
	dq := container.NewSliceQueue[*pb.Job]()
	aws := container.NewSimpleMapStore[int64, model.WorkerId]()
	jobMgr := NewSimpleJobMgr(wq, rq, dq, aws)
	orc := NewSimpleQueueOrchestrator(workerMgr, jobMgr)

	srv := grpc.NewServer()
	pb.RegisterTqWorkerServer(srv, newServer(orc))
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Printf("error serving: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		ctx,
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}))
	if err != nil {
		log.Printf("error dialing server: %v", err)
	}

	closer := func() {
		if err := lis.Close(); err != nil {
			log.Printf("error closing listener: %v", err)
		}
		srv.Stop()
	}

	c := pb.NewTqWorkerClient(conn)
	return c, closer
}

func TestTqServer_Register(t *testing.T) {
	ctx := context.Background()

	c, closer := testingServer(ctx)
	defer closer()

	type expectation struct {
		out *pb.RegisterResponse
		err error
	}

	tests := map[string]struct {
		in       *pb.RegisterRequest
		expected expectation
	}{
		"Successful_Register": {
			in: &pb.RegisterRequest{
				Label: "test worker",
			},
			expected: expectation{
				out: &pb.RegisterResponse{
					Registered: true,
					Id:         "testId",
				},
				err: nil,
			},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			out, err := c.Register(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Registered != out.Registered {
					t.Errorf("Out -> \nWant: %q\nGot: %q\n", tt.expected.out, out)
				}
			}
		})
	}
}

func TestTqServer_Deregister(t *testing.T) {
	ctx := context.Background()

	c, closer := testingServer(ctx)
	defer closer()

	type expectation struct {
		out *pb.DeregisterResponse
		err error
	}

	tests := map[string]struct {
		in       *pb.DeregisterRequest
		expected expectation
	}{
		"Successful_Deregister": {
			in: &pb.DeregisterRequest{
				Id: fixtureId,
			},
			expected: expectation{
				out: &pb.DeregisterResponse{
					Registered: false,
				},
				err: nil,
			},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			out, err := c.Deregister(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Registered != out.Registered {
					t.Errorf("Out -> \nWant: %q\nGot: %q\n", tt.expected.out, out)
				}
			}
		})
	}
}
