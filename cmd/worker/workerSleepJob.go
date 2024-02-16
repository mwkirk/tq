package main

import (
	"strconv"
	"time"
	"tq/pb"
)

type workerSleepJob struct {
	jobMsg  *pb.Job
	updates chan<- pb.JobStatus
}

func newWorkerSleepJob(jobMsg *pb.Job, updates chan<- pb.JobStatus) *workerSleepJob {
	return &workerSleepJob{
		jobMsg:  jobMsg,
		updates: updates,
	}
}

func (j *workerSleepJob) run() {
	duration, ok := j.jobMsg.Parms["duration"]
	if !ok {
		s := pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_ERR,
			Num:      j.jobMsg.Num,
			Progress: 0,
			Msg:      []string{"no duration specified for sleep job"},
		}
		j.updates <- s
		return
	}

	d, err := strconv.Atoi(duration)
	if err != nil {
		s := pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_ERR,
			Num:      j.jobMsg.Num,
			Progress: 0,
			Msg:      []string{"bad duration specified for sleep job"},
		}
		j.updates <- s
		return
	}

	go func() {
		chunk := d / 10
		for i := 0; i < d; i += chunk {
			s := pb.JobStatus{
				JobState: pb.JobState_JOB_STATE_RUN,
				Num:      j.jobMsg.Num,
				Progress: float32(i / d),
				Msg:      []string{"sleep job running"},
			}
			j.updates <- s
			time.Sleep(time.Duration(chunk) * time.Second)
		}

		s := pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_DONE_OK,
			Num:      j.jobMsg.Num,
			Progress: 1.0,
			Msg:      []string{"sleep job completed successfully"},
		}
		j.updates <- s
	}()
}