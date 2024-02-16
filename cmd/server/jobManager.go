package main

import (
	"log"
	"sync"
	"tq/internal/container"
	"tq/pb"
)
import "tq/internal/model"

type JobMgr interface {
	Submit(*pb.Job) error
	Cancel() error
	List() error
	AssignWorker(*pb.Job, model.WorkerId) error
	EnqueueWait(*pb.Job) error
	DequeueWait() (*pb.Job, error)
}

type JobQueue container.Queue[*pb.Job]
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
func (mgr *SimpleJobMgr) Submit(job *pb.Job) error {
	job.Num = uint32(mgr.newJobNumber())
	log.Printf("submitted job %v", job)
	return mgr.wait.Enqueue(job)
}

func (mgr *SimpleJobMgr) Cancel() error {
	// TODO implement me
	panic("implement me")
}

func (mgr *SimpleJobMgr) List() error {
	// TODO implement me
	panic("implement me")
}

func (mgr *SimpleJobMgr) AssignWorker(job *pb.Job, id model.WorkerId) error {
	err := mgr.run.Enqueue(job)
	if err != nil {
		return err
	}

	// With our crude implementation, there's not much that can be done to move the job back to the correct
	// queue is this fails
	return mgr.assignedWorkers.Add(model.JobNumber(job.Num), id)
}

// EnqueueWait adds a job to the wait queue without assigning a job number
func (mgr *SimpleJobMgr) EnqueueWait(job *pb.Job) error {
	return mgr.wait.Enqueue(job)
}

func (mgr *SimpleJobMgr) DequeueWait() (*pb.Job, error) {
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
