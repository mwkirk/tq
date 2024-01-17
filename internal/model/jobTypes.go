package model

type Job struct {
	num   int64
	kind  int
	name  string
	parms map[string]string
}
