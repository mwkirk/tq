package main

import "tq/pb"
import "context"

type workerJob interface {
	cancel()
	run()
	updateStatus(pb.JobState, float32, string)
}

type workerJobImpl struct {
	jobSpec    *pb.JobSpec
	updates    chan<- *pb.JobStatus
	ctx        context.Context
	cancelFunc context.CancelFunc
}
