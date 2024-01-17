package main

import (
	"errors"
	"fmt"
	"log"
	"tq/internal/container"
	"tq/internal/model"
	"tq/pbuf"
)

type WorkerMgr interface {
	Register(label string) (model.WorkerId, error)
	Deregister(id model.WorkerId) error
	Status(id model.WorkerId, state pbuf.WorkerState, status *pbuf.JobStatus) error
}

type WorkerStore interface {
	Get(model.WorkerId) (*model.Worker, error)
	Add(model.WorkerId, *model.Worker) error
	Remove(model.WorkerId) error
}

type SimpleWorkerMgr struct {
	store WorkerStore
}

func NewSimpleWorkerMgr(ws WorkerStore) *SimpleWorkerMgr {
	return &SimpleWorkerMgr{
		store: ws,
	}
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

func (mgr *SimpleWorkerMgr) Status(id model.WorkerId, workerState pbuf.WorkerState, jobStatus *pbuf.JobStatus) error {
	w, err := mgr.store.Get(id)
	if err != nil {
		if errors.Is(err, container.ErrorNotFound) {
			log.Printf("status error: unknown worker [%s]", id)
		}
		return err
	}

	log.Printf("worker reported status [%s, %v]", id, workerState)

	switch workerState {
	case pbuf.WorkerState_WORKER_STATE_UNAVAILABLE:
		// no-op for now
	case pbuf.WorkerState_WORKER_STATE_AVAILABLE:
		log.Printf("worker available [%s]", id)

	case pbuf.WorkerState_WORKER_STATE_WORKING:
		log.Printf("worker working [%s, %v]", id, jobStatus)

	default:
		log.Printf("bad worker state [%s, %v]. THIS SHOULD NOT HAPPEN", w.Id, w.WorkerState)
	}

	return nil
}
