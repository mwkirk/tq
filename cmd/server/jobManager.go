package main

import "tq/internal/container"
import "tq/internal/model"

type JobMgr interface {
	Submit(*model.Job) error
	Cancel() error
	List() error
	EnqueueWait(*model.Job) error
	DequeueWait() (*model.Job, error)
	EnqueueRun(*model.Job) error
}

type JobQueue container.Queue[*model.Job]

type SimpleJobMgr struct {
	wait JobQueue
	run  JobQueue
	done JobQueue
}

func NewSimpleJobMgr(waitQueue JobQueue, runQueue JobQueue, doneQueue JobQueue) *SimpleJobMgr {
	return &SimpleJobMgr{
		wait: waitQueue,
		run:  runQueue,
		done: doneQueue,
	}
}

func (mgr *SimpleJobMgr) Submit(job *model.Job) error {
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

func (mgr *SimpleJobMgr) EnqueueWait(job *model.Job) error {
	return mgr.wait.Enqueue(job)
}

func (mgr *SimpleJobMgr) DequeueWait() (*model.Job, error) {
	return mgr.wait.Dequeue()
}

func (mgr *SimpleJobMgr) EnqueueRun(job *model.Job) error {
	return mgr.run.Enqueue(job)
}
