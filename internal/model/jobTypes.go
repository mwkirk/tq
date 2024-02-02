package model

import "tq/pb"

type Job struct {
	Kind             pb.JobKind
	Num              int64
	Name             string
	Parms            map[string]string
	AssignedWorkerId WorkerId
}
