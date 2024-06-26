package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"
	"tq/pb"
)

type workerSleepJob struct {
	workerJobImpl
}

func newWorkerSleepJob(pctx context.Context, jobMsg *pb.JobSpec, updates chan<- *pb.JobStatus) *workerSleepJob {
	ctx, cancel := context.WithCancel(pctx)
	return &workerSleepJob{
		workerJobImpl{
			jobSpec:    jobMsg,
			updates:    updates,
			ctx:        ctx,
			cancelFunc: cancel,
		},
	}
}

func (j *workerSleepJob) cancel() {
	j.cancelFunc()
}

func (j *workerSleepJob) updateStatus(jobState pb.JobState, progress float32, msg string) {
	s := &pb.JobStatus{
		JobState: jobState,
		JobNum:   j.jobSpec.JobNum,
		Progress: progress,
		Msg:      []string{msg},
	}
	j.updates <- s
}

func (j *workerSleepJob) run() {
	duration, ok := j.jobSpec.Parms["duration"]
	if !ok {
		j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "no duration specified for sleep job")
		return
	}

	d, err := strconv.Atoi(duration)
	if err != nil {
		j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "bad duration specified for sleep job")
		return
	}

	shouldFail := false
	failRateStr, ok := j.jobSpec.Parms["failRate"]
	if ok {
		failRate, err := strconv.ParseFloat(failRateStr, 64)
		if err != nil {
			j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "bad failRate specified for sleep job")
			return
		}

		if failRate < 0 || failRate > 1 {
			j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "failRate must be between 0 and 1")
			return
		}

		if rand.Float64() < failRate {
			shouldFail = true
		}
	}

	go func() {
		chunk := max(d/10, 1)
		failTime := rand.Intn(d)
		if shouldFail {
			log.Printf("failTime: %d", failTime)
		}

		for i := 0; i < d; i += chunk {
			select {
			case <-j.ctx.Done():
				j.updateStatus(pb.JobState_JOB_STATE_DONE_CANCEL, float32(i)/float32(d), "sleep job canceled")
				return
			default:
				if shouldFail && failTime <= i {
					j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, float32(i)/float32(d), "sleep job failed")
					return
				}
				j.updateStatus(pb.JobState_JOB_STATE_RUN, float32(i)/float32(d), "sleep job running")
				time.Sleep(time.Duration(chunk) * time.Second)
			}
		}

		j.updateStatus(pb.JobState_JOB_STATE_DONE_OK, 1.0, "sleep job completed")
	}()
}
