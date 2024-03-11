package main

import (
	"fmt"
	"log"
	"sync"
	"tq/internal/container"
	"tq/pb"
)
import "tq/internal/model"

type JobMgr interface {
	Submit(*pb.JobSpec) error
	Cancel(model.JobNumber) error
	List(*pb.ListRequest) (*pb.ListResponse, error)
	AssignWorker(*pb.JobSpec, model.WorkerId) error
	UpdateJobHistory(model.JobNumber, []*pb.JobStatus) error
	EnqueueWait(*pb.JobSpec) error
	DequeueWait() (*pb.JobSpec, error)
}

type JobQueue container.Queue[*pb.JobSpec]
type JobStore container.Store[model.JobNumber, *pb.JobSpec]
type AssignedWorkerStore container.Store[model.JobNumber, model.WorkerId]
type JobHistoryStore container.Store[model.JobNumber, []*pb.JobStatus]

type SimpleJobMgr struct {
	l               sync.Mutex
	jobNum          model.JobNumber // next JobNumber
	wait            JobQueue        // queues of JobSpecs
	run             JobStore
	done            JobQueue
	assignedWorkers AssignedWorkerStore // JobNumber -> WorkerId
	jobHistory      JobHistoryStore     // JobNumber -> slices of JobStatus
}

func NewSimpleJobMgr(waitQueue JobQueue, runStore JobStore, doneQueue JobQueue,
	assignedWorkerStore AssignedWorkerStore, jobHistoryStore JobHistoryStore) *SimpleJobMgr {
	return &SimpleJobMgr{
		wait:            waitQueue,
		run:             runStore,
		done:            doneQueue,
		assignedWorkers: assignedWorkerStore,
		jobHistory:      jobHistoryStore,
	}
}

// Submit assigns a new job number to a job and adds it to the wait queue
func (mgr *SimpleJobMgr) Submit(job *pb.JobSpec) error {
	job.JobNum = uint32(mgr.newJobNumber())
	log.Printf("submitted job %v", job)
	return mgr.wait.Enqueue(job)
}

func (mgr *SimpleJobMgr) Cancel(jobNum model.JobNumber) error {
	// TODO implement me
	panic("implement me")
}

// List lists jobs filtered by the job state, job kind, and job numbers
func (mgr *SimpleJobMgr) List(req *pb.ListRequest) (*pb.ListResponse, error) {
	resp := &pb.ListResponse{
		Wait: nil,
		Run:  nil,
		Done: nil,
	}

	if req.JobStateFilter&int32(pb.JobState_JOB_STATE_WAIT) != 0 {
		resp.Wait = mgr.getJobList(req, mgr.wait)
	}

	if req.JobStateFilter&int32(pb.JobState_JOB_STATE_RUN) != 0 {
		resp.Run = mgr.getJobList(req, mgr.run)
	}

	if req.JobStateFilter&int32(pb.JobState_JOB_STATE_DONE_OK|pb.JobState_JOB_STATE_DONE_ERR|pb.JobState_JOB_STATE_DONE_CANCEL) != 0 {
		resp.Done = mgr.getJobList(req, mgr.done)
	}

	fmt.Printf("List: %v\n", resp)
	return resp, nil
}

func (mgr *SimpleJobMgr) AssignWorker(job *pb.JobSpec, id model.WorkerId) error {
	err := mgr.run.Add(model.JobNumber(job.JobNum), job)
	if err != nil {
		return err
	}

	// With our crude implementation, there's not much that can be done to move the job back to the correct
	// queue is this fails
	return mgr.assignedWorkers.Add(model.JobNumber(job.JobNum), id)
}

func (mgr *SimpleJobMgr) UpdateJobHistory(jobNum model.JobNumber, jobStatus []*pb.JobStatus) error {
	return mgr.jobHistory.Update(jobNum, func(v []*pb.JobStatus) []*pb.JobStatus {
		return append(v, jobStatus...)
	})
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

func (mgr *SimpleJobMgr) getAssignedWorkerId(jobNum model.JobNumber) model.WorkerId {
	v, err := mgr.assignedWorkers.Get(jobNum)
	if err != nil {
		return model.NullWorkerId
	}
	return v
}

func (mgr *SimpleJobMgr) getLatestJobStatus(jobNum model.JobNumber) *pb.JobStatus {
	h, err := mgr.jobHistory.Get(jobNum)
	if err != nil || len(h) == 0 {
		return &pb.JobStatus{
			JobState: pb.JobState_JOB_STATE_NONE,
			JobNum:   uint32(jobNum),
			Progress: 0.0,
			Msg:      nil,
		}
	}

	return h[len(h)-1]
}

func (mgr *SimpleJobMgr) getJobList(req *pb.ListRequest, c container.Filterable[*pb.JobSpec]) []*pb.JobListItem {
	jobSpecs := c.Filter(func(spec *pb.JobSpec) bool {
		pred := func() bool {
			if len(req.JobNums) == 0 {
				return true
			}

			for _, n := range req.JobNums {
				if spec.JobNum == n {
					return true
				}
			}
			return false
		}

		if req.JobKindFilter&int32(spec.Kind) != 0 && pred() {
			return true
		}
		return false
	})

	jobList := make([]*pb.JobListItem, 0, len(jobSpecs))
	for _, spec := range jobSpecs {
		jobNum := model.JobNumber(spec.JobNum)
		js := mgr.getLatestJobStatus(jobNum)
		item := &pb.JobListItem{
			JobNum:   spec.JobNum,
			JobState: js.JobState,
			Kind:     spec.Kind,
			Progress: js.Progress,
			Name:     spec.Name,
			Worker:   mgr.getAssignedWorkerId(model.JobNumber(spec.JobNum)).String(),
			Msg:      js.Msg,
		}

		jobList = append(jobList, item)
	}
	return jobList
}
