package main

import (
	"context"
	"strconv"
	"time"
	"tq/pb"
)

type workerSleepJob struct {
	jobMsg  *pb.JobSpec
	updates chan<- *pb.JobStatus
}

func newWorkerSleepJob(jobMsg *pb.JobSpec, updates chan<- *pb.JobStatus) *workerSleepJob {
	return &workerSleepJob{
		jobMsg:  jobMsg,
		updates: updates,
	}
}

func (j *workerSleepJob) run(ctx context.Context) {
	duration, ok := j.jobMsg.Parms["duration"]
	if !ok {
		s := &pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_ERR,
			JobNum:   j.jobMsg.JobNum,
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
			JobNum:   j.jobMsg.JobNum,
			Progress: 0,
			Msg:      []string{"bad duration specified for sleep job"},
		}
		j.updates <- s
		return
	}

	go func() {
		chunk := d / 10
		for i := 0; i < d; i += chunk {
			select {
			case <-ctx.Done():
				s := &pb.JobStatus{
					JobState: pb.JobState_JOB_STATE_DONE_CANCEL,
					JobNum:   j.jobMsg.JobNum,
					Progress: float32(i) / float32(d),
					Msg:      []string{"sleep job cancelling"},
				}
				j.updates <- s
				return
			default:
				s := &pb.JobStatus{
					JobState: pb.JobState_JOB_STATE_RUN,
					JobNum:   j.jobMsg.JobNum,
					Progress: float32(i) / float32(d),
					Msg:      []string{"sleep job running"},
				}
				j.updates <- s
				time.Sleep(time.Duration(chunk) * time.Second)
			}
		}

		s := &pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_OK,
			JobNum:   j.jobMsg.JobNum,
			Progress: 1.0,
			Msg:      []string{"sleep job completed successfully"},
		}
		j.updates <- s
	}()
}
