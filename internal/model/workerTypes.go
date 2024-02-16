package model

import (
	"github.com/google/uuid"
	"tq/pb"
)

type WorkerId string

const NullWorkerId WorkerId = ""

func NewWorkerId() (WorkerId, error) {
	u, err := uuid.NewRandom()
	return WorkerId(u.String()), err
}

func (id WorkerId) String() string {
	return string(id)
}

type Worker struct {
	Registered  bool
	Id          WorkerId
	Label       string
	WorkerState pb.WorkerState
	JobNum      JobNumber
}
