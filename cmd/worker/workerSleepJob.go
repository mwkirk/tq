package main

import (
	"context"
	"log"
	"strconv"
	"time"
	"tq/pb"
)

type job interface {
	run()
	cancel()
}

type workerSleepJob struct {
	jobSpec    *pb.JobSpec
	updates    chan<- *pb.JobStatus
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func newWorkerSleepJob(ctx context.Context, jobMsg *pb.JobSpec, updates chan<- *pb.JobStatus) *workerSleepJob {
	ctx, cancel := context.WithCancel(ctx)
	return &workerSleepJob{
		jobSpec:    jobMsg,
		updates:    updates,
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (j *workerSleepJob) cancel() {
	j.cancelFunc()
}

func (j *workerSleepJob) run() {
	duration, ok := j.jobSpec.Parms["duration"]
	if !ok {
		s := &pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_ERR,
			JobNum:   j.jobSpec.JobNum,
			Progress: 0,
			Msg:      []string{"no duration specified for sleep job"},
		}
		j.updates <- s
		return
	}

	d, err := strconv.Atoi(duration)
	if err != nil {
		s := &pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_ERR,
			JobNum:   j.jobSpec.JobNum,
			Progress: 0,
			Msg:      []string{"bad duration specified for sleep job"},
		}
		j.updates <- s
		return
	}

	go func() {
		chunk := max(d/10, 1)
		for i := 0; i < d; i += chunk {
			select {
			case <-j.ctx.Done():
				log.Printf("*************** cancelled **********************")
				s := &pb.JobStatus{
					JobState: pb.JobState_JOB_STATE_DONE_CANCEL,
					JobNum:   j.jobSpec.JobNum,
					Progress: float32(i) / float32(d),
					Msg:      []string{"sleep job canceled"},
				}
				j.updates <- s
				return
			default:
				s := &pb.JobStatus{
					JobState: pb.JobState_JOB_STATE_RUN,
					JobNum:   j.jobSpec.JobNum,
					Progress: float32(i) / float32(d),
					Msg:      []string{"sleep job running"},
				}
				j.updates <- s
				time.Sleep(time.Duration(chunk) * time.Second)
			}
		}

		s := &pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_OK,
			JobNum:   j.jobSpec.JobNum,
			Progress: 1.0,
			Msg:      []string{"sleep job completed"},
		}
		j.updates <- s
	}()
}
