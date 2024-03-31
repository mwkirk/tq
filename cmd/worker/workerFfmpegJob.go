package main

import (
	"context"
	"fmt"
	"github.com/xfrr/goffmpeg/transcoder"
	"tq/pb"
)

type workerFfmpegJob struct {
	workerJobImpl
}

func newWorkerFfmpegJob(pctx context.Context, jobMsg *pb.JobSpec, updates chan<- *pb.JobStatus) *workerFfmpegJob {
	ctx, cancel := context.WithCancel(pctx)
	return &workerFfmpegJob{
		workerJobImpl{
			jobSpec:    jobMsg,
			updates:    updates,
			ctx:        ctx,
			cancelFunc: cancel,
		},
	}
}

func (j *workerFfmpegJob) cancel() {
	j.cancelFunc()
}

func (j *workerFfmpegJob) updateStatus(jobState pb.JobState, progress float32, msg string) {
	s := &pb.JobStatus{
		JobState: jobState,
		JobNum:   j.jobSpec.JobNum,
		Progress: progress,
		Msg:      []string{msg},
	}
	j.updates <- s
}

func (j *workerFfmpegJob) run() {
	inputPath, ok := j.jobSpec.Parms["inputPath"]
	if !ok {
		j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "inputPath not specified for ffmpeg job")
		return
	}

	outputPath, ok := j.jobSpec.Parms["outputPath"]
	if !ok {
		j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "outputPath not specified for ffmpeg job")
		return
	}

	// keyInfoPath, ok := j.jobSpec.Parms["keyInfoPath"]
	// if !ok {
	// 	j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "keyInfoPath not specified for ffmpeg job")
	// 	return
	// }

	trans := new(transcoder.Transcoder)

	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "error initializing ffmpeg")
		return
	}

	go func() {
		trans.MediaFile().SetVideoCodec("libx264")
		trans.MediaFile().SetHlsSegmentDuration(4)
		// trans.MediaFile().SetEncryptionKey(keyInfoPath)

		done := trans.Run(true)
		progress := trans.Output()

		for {
			select {
			case p, ok := <-progress:
				if ok {
					msg := fmt.Sprintf("ffmpeg running: %v", p)
					j.updateStatus(pb.JobState_JOB_STATE_RUN, float32(p.Progress)/100.0, msg)
				}
			case <-j.ctx.Done():
				err := trans.Stop()
				if err != nil {
					j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "error stopping ffmpeg")
					return
				} else {
					j.updateStatus(pb.JobState_JOB_STATE_DONE_CANCEL, 0, "ffmpeg job cancelled")
					return
				}
			case err := <-done:
				if err != nil {
					j.updateStatus(pb.JobState_JOB_STATE_DONE_ERR, 0, "error running ffmpeg")
					return
				} else {
					j.updateStatus(pb.JobState_JOB_STATE_DONE_OK, 1, "ffmpeg job completed")
					return
				}
			}
		}
	}()
}
