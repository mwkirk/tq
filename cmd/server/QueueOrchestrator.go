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
	Status(id model.WorkerId, state pb.WorkerState, status *pb.JobStatus) (StatusResponse, error)
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
	jobStatus *pb.JobStatus) (StatusResponse, error) {
	log.Printf("worker reported status [%s, %v]", id, workerState)

	switch workerState {
	case pb.WorkerState_WORKER_STATE_UNAVAILABLE:
		return StatusResponse{}, nil // no-op for now
	case pb.WorkerState_WORKER_STATE_AVAILABLE:
		log.Printf("worker available [%s]", id)
		return orc.dispatch(id)
	case pb.WorkerState_WORKER_STATE_WORKING:
		log.Printf("worker working [%s, %v]", id, jobStatus)
		return StatusResponse{JobControl: pb.JobControl_JOB_CONTROL_CONTINUE}, nil
	default:
		return StatusResponse{}, fmt.Errorf("bad worker state [%s, %v]. THIS SHOULD NOT HAPPEN", id, workerState)
	}
}

// ------------------------------------------------------------------
// Unexported methods
// ------------------------------------------------------------------

// Dispatch a job to an available worker
func (orc *SimpleQueueOrchestrator) dispatch(id model.WorkerId) (StatusResponse, error) {
	// guard - do we know about this worker?
	_, err := orc.workerMgr.Exists(id)
	if err != nil {
		if errors.Is(err, container.ErrorNotFound) {
			log.Printf("status error: unknown worker [%s]", id)
		}
		return StatusResponse{}, err
	}
	log.Printf("here 1")
	// Dequeue a job from the wait queue (safe) to prevent another goroutine from grabbing it.
	// If the wait queue is empty, then respond with job control none
	j, err := orc.jobMgr.DequeueWait()
	if err != nil {
		if errors.Is(err, container.ErrorQueueEmpty) {
			log.Printf("no jobs waiting")
			return StatusResponse{JobControl: pb.JobControl_JOB_CONTROL_NONE}, nil
		}
		return StatusResponse{JobControl: pb.JobControl_JOB_CONTROL_NONE}, err
	}

	log.Printf("here 2")
	// Assign job to worker
	err = orc.workerMgr.AssignJob(id, j.Num)
	if err != nil {
		// If we can't assign the job to the worker, then try to put the job back in the wait queue.
		// This could fail as well, of course, but we'll just swallow the error here since it's just for fun.
		// A production system would need to be transactional.
		_ = orc.jobMgr.EnqueueWait(j)
		return StatusResponse{}, fmt.Errorf("error assigning job %d to worker [%s]: %w", j.Num, id, err)
	}
	log.Printf("here 3")

	// Move job to the run queue
	err = orc.jobMgr.EnqueueRun(j)
	if err != nil {
		// Again, only so much we can do to roll back to a correct state w/o a more transactional design
		log.Printf("failed to move job [%d] to run queue", j.Num)
		_ = orc.workerMgr.Reset(id)
		return StatusResponse{}, fmt.Errorf("error moving job %d to run queue: %w", j.Num, err)
	}

	log.Printf("here 4")
	// Assign the worker's ID to the job
	j.AssignedWorkerId = id

	// Build the final response
	log.Printf("dispatching job %d to worker [%s]", j.Num, id)
	return StatusResponse{
		JobControl: pb.JobControl_JOB_CONTROL_NEW,
		Job:        *j,
	}, nil
}
