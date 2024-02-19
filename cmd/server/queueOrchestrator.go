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
	Register(label string) (model.WorkerId, error)
	Deregister(id model.WorkerId) error
	Status(id model.WorkerId, state pb.WorkerState, status []*pb.JobStatus) (pb.StatusResponse, error)
	Submit(job *pb.Job) error
	Cancel(jobNum int64) error
	List() error
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

func (orc *SimpleQueueOrchestrator) Register(label string) (model.WorkerId, error) {
	// simple delegation for now
	return orc.workerMgr.Register(label)
}

func (orc *SimpleQueueOrchestrator) Deregister(id model.WorkerId) error {
	// simple delegation for now
	return orc.workerMgr.Deregister(id)
}

func (orc *SimpleQueueOrchestrator) Status(id model.WorkerId, workerState pb.WorkerState,
	jobStatus []*pb.JobStatus) (pb.StatusResponse, error) {
	log.Printf("worker reported status [%s, %v]", id, workerState)

	switch workerState {
	case pb.WorkerState_WORKER_STATE_UNAVAILABLE:
		return pb.StatusResponse{}, nil // no-op for now
	case pb.WorkerState_WORKER_STATE_AVAILABLE:
		return orc.dispatch(id)
	case pb.WorkerState_WORKER_STATE_WORKING:
		log.Printf("worker working [%s, %v]", id, jobStatus)
		return pb.StatusResponse{JobControl: pb.JobControl_JOB_CONTROL_CONTINUE}, nil
	default:
		return pb.StatusResponse{}, fmt.Errorf("bad worker state [%s, %v]. THIS SHOULD NOT HAPPEN", id, workerState)
	}
}

func (orc *SimpleQueueOrchestrator) Submit(job *pb.Job) error {
	return orc.jobMgr.Submit(job)
}

func (orc *SimpleQueueOrchestrator) Cancel(jobNum int64) error {
	// TODO implement me
	panic("implement me")
}

func (orc *SimpleQueueOrchestrator) List() error {
	// TODO implement me
	panic("implement me")
}

// ------------------------------------------------------------------
// Unexported methods
// ------------------------------------------------------------------

// Dispatch a job to an available worker
func (orc *SimpleQueueOrchestrator) dispatch(id model.WorkerId) (pb.StatusResponse, error) {
	// guard - do we know about this worker?
	_, err := orc.workerMgr.Exists(id)
	if err != nil {
		if errors.Is(err, container.ErrorNotFound) {
			log.Printf("status error: unknown worker [%s]", id)
		}
		return pb.StatusResponse{}, err
	}

	// Dequeue a job from the wait queue (safe) to prevent another goroutine from grabbing it.
	// If the wait queue is empty, then respond with job control none
	j, err := orc.jobMgr.DequeueWait()
	if err != nil {
		if errors.Is(err, container.ErrorQueueEmpty) {
			log.Printf("no jobs waiting")
			return pb.StatusResponse{JobControl: pb.JobControl_JOB_CONTROL_NONE}, nil
		}
		return pb.StatusResponse{JobControl: pb.JobControl_JOB_CONTROL_NONE}, err
	}

	// Assign job to worker
	err = orc.workerMgr.AssignJob(id, model.JobNumber(j.Num))
	if err != nil {
		// If we can't assign the job to the worker, then try to put the job back in the wait queue.
		// This could fail as well, of course, but we'll just swallow the error here since it's just for fun.
		// A production system would need to be transactional.
		_ = orc.jobMgr.EnqueueWait(j)
		return pb.StatusResponse{}, fmt.Errorf("error assigning job %d to worker [%s]: %w", j.Num, id, err)
	}

	// Assign worker to job
	err = orc.jobMgr.AssignWorker(j, id)
	if err != nil {
		// Again, only so much we can do to roll back to a correct state w/o a more transactional design
		log.Printf("failed to assign job %d to worker [%s] and move to run queue", j.Num, id)
		_ = orc.workerMgr.Reset(id)
		return pb.StatusResponse{}, fmt.Errorf("error moving job %d to run queue and assigning worker: %w", j.Num, err)
	}

	// Build the final response
	log.Printf("dispatching job %d to worker [%s]", j.Num, id)
	return pb.StatusResponse{
		JobControl: pb.JobControl_JOB_CONTROL_NEW,
		Job:        j,
	}, nil
}
