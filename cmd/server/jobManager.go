package main

import (
	"log"
	"sync"
	"tq/internal/container"
	"tq/pb"
)
import "tq/internal/model"

type JobMgr interface {
	Submit(*pb.JobSpec) error
	Cancel(number model.JobNumber) error
	List() ([]*pb.JobStatus, error)
	AssignWorker(*pb.JobSpec, model.WorkerId) error
	EnqueueWait(*pb.JobSpec) error
	DequeueWait() (*pb.JobSpec, error)
}

type JobQueue container.Queue[*pb.JobSpec]
type AssignedWorkerStore container.Store[model.JobNumber, model.WorkerId]

type SimpleJobMgr struct {
	l               sync.Mutex
	jobNum          model.JobNumber
	wait            JobQueue
	run             JobQueue
	done            JobQueue
	assignedWorkers AssignedWorkerStore
}

func NewSimpleJobMgr(waitQueue JobQueue, runQueue JobQueue, doneQueue JobQueue,
	store AssignedWorkerStore) *SimpleJobMgr {
	return &SimpleJobMgr{
		wait:            waitQueue,
		run:             runQueue,
		done:            doneQueue,
		assignedWorkers: store,
	}
}

// Submit assigns a new job number to a job and adds it to the wait queue
func (mgr *SimpleJobMgr) Submit(job *pb.JobSpec) error {
	job.Num = uint32(mgr.newJobNumber())
	log.Printf("submitted job %v", job)
	return mgr.wait.Enqueue(job)
}

func (mgr *SimpleJobMgr) Cancel(jobNum model.JobNumber) error {
	// TODO implement me
	panic("implement me")
}

func (mgr *SimpleJobMgr) List() ([]*pb.JobStatus, error) {
	// TODO implement me
	panic("implement me")
}

func (mgr *SimpleJobMgr) AssignWorker(job *pb.JobSpec, id model.WorkerId) error {
	err := mgr.run.Enqueue(job)
	if err != nil {
		return err
	}

	// With our crude implementation, there's not much that can be done to move the job back to the correct
	// queue is this fails
	return mgr.assignedWorkers.Add(model.JobNumber(job.Num), id)
}

// EnqueueWait adds a job to the wait queue without assigning a job number
func (mgr *SimpleJobMgr) EnqueueWait(job *pb.JobSpec) error {
	return mgr.wait.Enqueue(job)
}

func (mgr *SimpleJobMgr) DequeueWait() (*pb.JobSpec, error) {
	return mgr.wait.Dequeue()
}

// ------------------------------------------------------------------
// Unexported methods
// ------------------------------------------------------------------

// todo: Should initialize and persist job number between runs of server
func (mgr *SimpleJobMgr) newJobNumber() model.JobNumber {
	mgr.l.Lock()
	defer mgr.l.Unlock()
	mgr.jobNum++
	return mgr.jobNum
}
