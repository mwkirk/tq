package model

import "tq/pbuf"

type Job struct {
	Kind             pbuf.JobKind
	Num              int64
	Name             string
	Parms            map[string]string
	AssignedWorkerId WorkerId
}
