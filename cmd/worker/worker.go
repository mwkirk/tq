package main

import (
	"context"
	"fmt"
	"log"
	"tq/internal/model"
	"tq/pb"
)

type worker struct {
	model.Worker
	job job
}

func (w *worker) startJob(ctx context.Context, job *pb.JobSpec, updates chan<- *pb.JobStatus) error {
	// guard
	if job == nil {
		return fmt.Errorf("no job data")
	}

	switch job.Kind {
	case pb.JobKind_JOB_KIND_NULL:
		return fmt.Errorf("new null job received")
	case pb.JobKind_JOB_KIND_TEST:
		log.Printf("test job received")
	case pb.JobKind_JOB_KIND_SLEEP:
		w.job = newWorkerSleepJob(ctx, job, updates)
		w.job.run()
	case pb.JobKind_JOB_KIND_FFMPEG:
		log.Printf("FFMPEG job received")
	default:
		return fmt.Errorf("unexpected job kind")
	}

	return nil
}
