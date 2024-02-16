package main

import (
	"errors"
	"fmt"
	"log"
	"tq/internal/container"
	"tq/internal/model"
	"tq/pb"
)

type WorkerMgr interface {
	Exists(id model.WorkerId) (bool, error)
	Register(label string) (model.WorkerId, error)
	Deregister(id model.WorkerId) error
	AssignJob(id model.WorkerId, jobNum model.JobNumber) error
	Reset(id model.WorkerId) error
}

type WorkerStore container.Store[model.WorkerId, *model.Worker]

type SimpleWorkerMgr struct {
	store WorkerStore
}

func NewSimpleWorkerMgr(ws WorkerStore) *SimpleWorkerMgr {
	return &SimpleWorkerMgr{
		store: ws,
	}
}

func (mgr *SimpleWorkerMgr) Exists(id model.WorkerId) (bool, error) {
	return mgr.store.Exists(id)
}

func (mgr *SimpleWorkerMgr) Register(label string) (model.WorkerId, error) {
	id, err := model.NewWorkerId()
	if err != nil {
		return id, fmt.Errorf("failed to register worker: %w", err)
	}

	w := &model.Worker{
		Registered: true,
		Id:         id,
		Label:      label,
	}

	err = mgr.store.Add(id, w)
	if err != nil {
		return id, fmt.Errorf("failed to register worker: %w", err)
	}

	log.Printf("registered new worker [%s, %s]", w.Id, w.Label)
	return id, nil
}

func (mgr *SimpleWorkerMgr) Deregister(id model.WorkerId) error {
	w, err := mgr.store.Get(id)
	if err != nil {
		if errors.Is(err, container.ErrorNotFound) {
			log.Printf("deregistration error: unknown worker [%s]", id)
		}
		return err
	}

	w.Registered = false

	log.Printf("deregistered worker [%s, %s]", w.Id, w.Label)
	return nil
}

func (mgr *SimpleWorkerMgr) AssignJob(id model.WorkerId, jobNum model.JobNumber) error {
	return mgr.store.Update(id, func(w *model.Worker) *model.Worker {
		w.JobNum = jobNum
		w.WorkerState = pb.WorkerState_WORKER_STATE_WORKING
		return w
	})
}

func (mgr *SimpleWorkerMgr) Reset(id model.WorkerId) error {
	return mgr.store.Update(id, func(w *model.Worker) *model.Worker {
		w.JobNum = 0
		w.WorkerState = pb.WorkerState_WORKER_STATE_UNAVAILABLE
		return w
	})
}
