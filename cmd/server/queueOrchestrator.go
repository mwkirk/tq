package main

import (
	"errors"
	"fmt"
	"log"
	"tq/internal/container"
	"tq/internal/model"
	"tq/pb"
)

type QueueOrchestrator interface {
	Register(*pb.RegisterOptions) (*pb.RegisterResult, error)
	Deregister(*pb.DeregisterOptions) (*pb.DeregisterResult, error)
	Status(*pb.StatusOptions) (*pb.StatusResult, error)
	Submit(*pb.SubmitOptions) (*pb.SubmitResult, error)
	Cancel(*pb.CancelOptions) (*pb.CancelResult, error)
	List(*pb.ListOptions) (*pb.ListResult, error)
}

type SimpleQueueOrchestrator struct {
	workerMgr WorkerMgr
	jobMgr    JobMgr
}

func NewSimpleQueueOrchestrator(workerMgr WorkerMgr, jobMgr JobMgr) *SimpleQueueOrchestrator {
	return &SimpleQueueOrchestrator{
		workerMgr: workerMgr,
		jobMgr:    jobMgr,
	}
}

func (orc *SimpleQueueOrchestrator) Register(options *pb.RegisterOptions) (*pb.RegisterResult, error) {
	id, err := orc.workerMgr.Register(options.Label)
	return &pb.RegisterResult{
		Registered: err == nil,
		WorkerId:   string(id)}, err
}

func (orc *SimpleQueueOrchestrator) Deregister(options *pb.DeregisterOptions) (*pb.DeregisterResult, error) {
	err := orc.workerMgr.Deregister(model.WorkerId(options.WorkerId))
	return &pb.DeregisterResult{
		Deregistered: err == nil,
	}, err
}

func (orc *SimpleQueueOrchestrator) Status(options *pb.StatusOptions) (*pb.StatusResult, error) {
	log.Printf("worker %s status %v", options.WorkerId, options.WorkerState)

	if len(options.JobStatus) > 0 && options.JobStatus[0].JobNum != 0 {
		n := model.JobNumber(options.JobStatus[0].JobNum)
		err := orc.jobMgr.UpdateJobHistory(n, options.JobStatus)
		if err != nil {
			log.Printf("error updating job history: %s", err)
		}

		last := options.JobStatus[len(options.JobStatus)-1]
		switch last.JobState {
		case pb.JobState_JOB_STATE_DONE_OK:
			fallthrough
		case pb.JobState_JOB_STATE_DONE_ERR:
			fallthrough
		case pb.JobState_JOB_STATE_DONE_CANCEL:
			orc.finish(last)
		}
	}

	switch options.WorkerState {
	case pb.WorkerState_WORKER_STATE_UNAVAILABLE:
		return &pb.StatusResult{}, nil // no-op for now
	case pb.WorkerState_WORKER_STATE_AVAILABLE:
		return orc.dispatch(model.WorkerId(options.WorkerId))
	case pb.WorkerState_WORKER_STATE_WORKING:
		log.Printf("worker working [%s, %v]", options.WorkerId, options.JobStatus)

		// Check to see if the worker's job has been canceled
		jobNum, err := orc.workerMgr.GetAssignedJob(model.WorkerId(options.WorkerId))
		if err != nil {
			return &pb.StatusResult{}, err
		}

		shouldCancel, err := orc.jobMgr.MarkedForCancellation(jobNum)
		if err != nil {
			return &pb.StatusResult{}, err
		}

		if shouldCancel {
			err = orc.jobMgr.UnmarkForCancellation(jobNum)
			if err != nil {
				return &pb.StatusResult{}, err
			}

			log.Printf("canceling job %d for worker %s", jobNum, options.WorkerId)
			return &pb.StatusResult{JobControl: pb.JobControl_JOB_CONTROL_CANCEL}, nil
		} else {
			return &pb.StatusResult{JobControl: pb.JobControl_JOB_CONTROL_CONTINUE}, nil
		}
	default:
		return &pb.StatusResult{}, fmt.Errorf("bad worker state [%s, %v]. THIS SHOULD NOT HAPPEN", options.WorkerId,
			options.WorkerState)
	}
}

func (orc *SimpleQueueOrchestrator) Submit(options *pb.SubmitOptions) (*pb.SubmitResult, error) {
	return orc.jobMgr.Submit(options.JobSpec)
}

func (orc *SimpleQueueOrchestrator) Cancel(options *pb.CancelOptions) (*pb.CancelResult, error) {
	return orc.jobMgr.Cancel(options)
}

func (orc *SimpleQueueOrchestrator) List(options *pb.ListOptions) (*pb.ListResult, error) {
	return orc.jobMgr.List(options)
}

// ------------------------------------------------------------------
// Unexported methods
// ------------------------------------------------------------------

// Dispatch a job to an available worker
func (orc *SimpleQueueOrchestrator) dispatch(id model.WorkerId) (*pb.StatusResult, error) {
	// guard - do we know about this worker?
	_, err := orc.workerMgr.Exists(id)
	if err != nil {
		if errors.Is(err, container.ErrorNotFound) {
			log.Printf("status error: unknown worker [%s]", id)
		}
		return &pb.StatusResult{}, err
	}

	// Dequeue a job from the wait queue (safe) to prevent another goroutine from grabbing it.
	// If the wait queue is empty, then respond with job control none
	j, err := orc.jobMgr.DequeueWait()
	if err != nil {
		if errors.Is(err, container.ErrorQueueEmpty) {
			// log.Printf("no jobs waiting")
			return &pb.StatusResult{JobControl: pb.JobControl_JOB_CONTROL_NONE}, nil
		}
		return &pb.StatusResult{JobControl: pb.JobControl_JOB_CONTROL_NONE}, err
	}

	// Assign job to worker
	err = orc.workerMgr.AssignJob(id, model.JobNumber(j.JobNum))
	if err != nil {
		// If we can't assign the job to the worker, then try to put the job back in the wait queue.
		// This could fail as well, of course, but we'll just swallow the error here since it's just for fun.
		// A production system would need to be transactional.
		_ = orc.jobMgr.EnqueueWait(j)
		return &pb.StatusResult{}, fmt.Errorf("error assigning job %d to worker [%s]: %w", j.JobNum, id, err)
	}

	// Assign worker to job
	err = orc.jobMgr.AssignWorker(j, id)
	if err != nil {
		// Again, only so much we can do to roll back to a correct state w/o a more transactional design
		log.Printf("failed to assign job %d to worker [%s] and move to run queue", j.JobNum, id)
		_ = orc.workerMgr.Reset(id)
		return &pb.StatusResult{}, fmt.Errorf("error moving job %d to run queue and assigning worker: %w", j.JobNum,
			err)
	}

	// Build the final response
	log.Printf("dispatching job %d to worker [%s]", j.JobNum, id)
	return &pb.StatusResult{
		JobControl: pb.JobControl_JOB_CONTROL_NEW,
		Job:        j,
	}, nil
}

func (orc *SimpleQueueOrchestrator) finish(status *pb.JobStatus) error {
	id, err := orc.jobMgr.Finish(model.JobNumber(status.JobNum))
	if err != nil {
		return err
	}

	err = orc.workerMgr.Reset(id)
	if err != nil {
		return err
	}

	return nil
}
