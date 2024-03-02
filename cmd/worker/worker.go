package main

import (
	"context"
	"fmt"
	"log"
	"tq/internal/model"
	"tq/pb"
)

func handleStatusResponse(ctx context.Context, sr *pb.StatusResponse, w *model.Worker,
	updates chan<- *pb.JobStatus) error {
	switch sr.JobControl {
	case pb.JobControl_JOB_CONTROL_NONE:
		log.Printf("no job available")
	case pb.JobControl_JOB_CONTROL_CONTINUE:
		log.Printf("continue current job")
	case pb.JobControl_JOB_CONTROL_NEW:
		log.Printf("new job")
		err := startJob(ctx, sr.Job, updates)
		if err != nil {
			return err
		}
		w.WorkerState = pb.WorkerState_WORKER_STATE_WORKING
	case pb.JobControl_JOB_CONTROL_CANCEL:
		log.Printf("cancel current job")
		w.WorkerState = pb.WorkerState_WORKER_STATE_AVAILABLE
	default:
		return fmt.Errorf("received unexpect job control message")
	}

	return nil
}

func startJob(ctx context.Context, job *pb.JobSpec, updates chan<- *pb.JobStatus) error {
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
		workerJob := newWorkerSleepJob(job, updates)
		workerJob.run(ctx)
	case pb.JobKind_JOB_KIND_FFMPEG:
		log.Printf("FFMPEG job received")
	default:
		return fmt.Errorf("unexpected job kind")
	}

	return nil
}
