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
	Submit(*pb.JobSpec) (*pb.SubmitResult, error)
	Cancel(*pb.CancelOptions) (*pb.CancelResult, error)
	List(*pb.ListOptions) (*pb.ListResult, error)
	AssignWorker(*pb.JobSpec, model.WorkerId) error
	UpdateJobHistory(model.JobNumber, []*pb.JobStatus) error
	EnqueueWait(*pb.JobSpec) error
	DequeueWait() (*pb.JobSpec, error)
	Finish(model.JobNumber) (model.WorkerId, error)
	Requeue(model.JobNumber) (model.WorkerId, error)
	MarkedForCancellation(model.JobNumber) (bool, error)
	UnmarkForCancellation(number model.JobNumber) error
}

type JobQueue container.Queue[*pb.JobSpec]
type JobStore container.KVStore[model.JobNumber, *pb.JobSpec]
type CancelStore container.KVStore[model.JobNumber, bool]
type AssignedWorkerStore container.KVStore[model.JobNumber, model.WorkerId]
type JobHistoryStore container.KVStore[model.JobNumber, []*pb.JobStatus]

type SimpleJobMgr struct {
	l               sync.Mutex
	jobNum          model.JobNumber // next JobNumber
	wait            JobQueue        // queues of JobSpecs
	run             JobStore
	done            JobQueue
	cancel          CancelStore
	assignedWorkers AssignedWorkerStore // JobNumber -> WorkerId
	jobHistory      JobHistoryStore     // JobNumber -> slices of JobStatus
}

// NewSimpleJobMgr creates a new instance of SimpleJobMgr with the provided dependencies and returns a pointer to it.
// The SimpleJobMgr is responsible for managing jobs, including submitting, canceling, listing, assigning workers, updating job history, and handling job queues and stores.
// Parameters:
// - waitQueue: The job queue that holds waiting jobs (type: JobQueue)
// - runStore: The key-value store that holds running jobs (type: JobStore)
// - doneQueue: The job queue that holds completed jobs (type: JobQueue)
// - assignedWorkerStore: The key-value store that holds assigned workers for jobs (type: AssignedWorkerStore)
// - jobHistoryStore: The key-value store that holds job history information (type: JobHistoryStore)
// Returns:
// - *SimpleJobMgr: A pointer to the created SimpleJobMgr instance
func NewSimpleJobMgr(waitQueue JobQueue, runStore JobStore, doneQueue JobQueue, cancelStore CancelStore,
	assignedWorkerStore AssignedWorkerStore, jobHistoryStore JobHistoryStore) *SimpleJobMgr {
	return &SimpleJobMgr{
		wait:            waitQueue,
		run:             runStore,
		done:            doneQueue,
		cancel:          cancelStore,
		assignedWorkers: assignedWorkerStore,
		jobHistory:      jobHistoryStore,
	}
}

// Submit assigns a new job number to a job and adds it to the wait queue
func (mgr *SimpleJobMgr) Submit(job *pb.JobSpec) (*pb.SubmitResult, error) {
	result := &pb.SubmitResult{
		Accepted: false,
	}
	jobNum := mgr.newJobNumber()
	job.JobNum = uint32(jobNum)
	err := mgr.wait.Enqueue(job)
	if err != nil {
		return result, err
	}

	// create an initial JobStatus entry
	status := &pb.JobStatus{
		JobState: pb.JobState_JOB_STATE_WAIT,
		JobNum:   uint32(jobNum),
		Progress: 0.0,
		Msg:      []string{"accepted"},
	}

	err = mgr.jobHistory.Put(jobNum, []*pb.JobStatus{status})
	if err != nil {
		return result, err
	}

	result.Accepted = true
	result.JobStatus = status
	log.Printf("accepted job %v", result)
	return result, err
}

// Cancel marks a job for cancellation by placing it into the cancellation store if it's currently
// in the run queue. If it's in the wait queue, we'll cancel it by moving it to the done queue.
func (mgr *SimpleJobMgr) Cancel(options *pb.CancelOptions) (*pb.CancelResult, error) {
	jobNum := model.JobNumber(options.JobNum)
	result := &pb.CancelResult{
		Canceled: false,
		// todo: set the job status. What was I thinking here? The last status?
		// JobStatus:
	}

	markCanceled := func(jobNum model.JobNumber) (*pb.CancelResult, error) {
		err := mgr.cancel.Put(jobNum, true)
		if err != nil {
			return result, fmt.Errorf("unable to mark job %d for cancellation: %w", options.JobNum, err)
		}
		result.Canceled = true
		return result, nil
	}

	// Check if the job is in the run queue, and move it to the cancellation store if it is
	exists, err := mgr.run.Exists(jobNum)
	if err != nil {
		return result, fmt.Errorf("unable to mark job %d for cancellation: %w", options.JobNum, err)
	} else if exists {
		return markCanceled(jobNum)
	}

	// Check the wait queue, and cancel the job by moving it to the done queue if it's there
	matchesJobNum := func(jobSpec *pb.JobSpec) bool {
		if jobSpec.JobNum == uint32(jobNum) {
			return true
		}
		return false
	}

	jobSpec, exists := mgr.wait.FindFirstAndDelete(matchesJobNum)
	if exists {
		err := mgr.UpdateJobHistory(jobNum, []*pb.JobStatus{
			{
				JobState: pb.JobState_JOB_STATE_DONE_CANCEL,
				JobNum:   uint32(jobNum),
				Progress: 0.0,
				Msg:      []string{"canceled"},
			},
		})
		if err != nil {
			return result, fmt.Errorf("unable to cancel job %d: %w", options.JobNum, err)
		}

		err = mgr.done.Enqueue(jobSpec)
		if err != nil {
			return result, fmt.Errorf("unable to cancel job %d: %w", options.JobNum, err)
		}
		result.Canceled = true
		return result, nil
	}

	// Check the done queue so that the user can be informed
	_, exists = mgr.done.FindFirst(matchesJobNum)
	if exists {
		return result, fmt.Errorf("job %d has already completed", jobNum)
	}

	return result, fmt.Errorf("job %d cannot be found", jobNum)
}

// List lists jobs filtered by the job state, job kind, and job numbers
func (mgr *SimpleJobMgr) List(options *pb.ListOptions) (*pb.ListResult, error) {
	resp := &pb.ListResult{
		Wait: nil,
		Run:  nil,
		Done: nil,
	}

	if options.JobFilter.JobStateFilter&int32(pb.JobState_JOB_STATE_WAIT) != 0 {
		resp.Wait = mgr.getJobList(options, mgr.wait)
	}

	if options.JobFilter.JobStateFilter&int32(pb.JobState_JOB_STATE_RUN) != 0 {
		resp.Run = mgr.getJobList(options, mgr.run)
	}

	if options.JobFilter.JobStateFilter&int32(pb.JobState_JOB_STATE_DONE_OK|pb.JobState_JOB_STATE_DONE_ERR|pb.
		JobState_JOB_STATE_DONE_CANCEL) != 0 {
		resp.Done = mgr.getJobList(options, mgr.done)
	}

	fmt.Printf("List: %v\n", resp)
	return resp, nil
}

// AssignWorker assigns a worker to a job and updates the run queue and assigned workers map
func (mgr *SimpleJobMgr) AssignWorker(job *pb.JobSpec, id model.WorkerId) error {
	err := mgr.run.Put(model.JobNumber(job.JobNum), job)
	if err != nil {
		return err
	}

	// With our crude implementation, there's not much that can be done to move the job back to the correct
	// queue is this fails
	return mgr.assignedWorkers.Put(model.JobNumber(job.JobNum), id)
}

// UpdateJobHistory updates the job history with the given job number and job status.
// It appends the job status to the existing job status list for the specified job.
// The updated job status list is then stored in the job history store.
func (mgr *SimpleJobMgr) UpdateJobHistory(jobNum model.JobNumber, jobStatus []*pb.JobStatus) error {
	return mgr.jobHistory.Update(jobNum, func(v []*pb.JobStatus) []*pb.JobStatus {
		return append(v, jobStatus...)
	})
}

// EnqueueWait adds a job to the wait queue
func (mgr *SimpleJobMgr) EnqueueWait(job *pb.JobSpec) error {
	return mgr.wait.Enqueue(job)
}

// DequeueWait retrieves and removes a job from the wait queue
func (mgr *SimpleJobMgr) DequeueWait() (*pb.JobSpec, error) {
	return mgr.wait.Dequeue()
}

// Finish completes a job by performing the following steps:
// - Retrieves the assigned worker ID for the job number
// - Removes the job from the run store
// - Enqueues the job into the done queue
// - Deletes the job number from the assigned workers map
// - Returns the assigned worker ID and any error that occurred
func (mgr *SimpleJobMgr) Finish(jobNum model.JobNumber) (model.WorkerId, error) {
	id := mgr.getAssignedWorkerId(jobNum)

	// remove job from run store
	j, err := mgr.run.GetAndDelete(jobNum)
	if err != nil {
		return id, fmt.Errorf("failed to finish job: %w", err)
	}

	// and move to done queue
	err = mgr.done.Enqueue(j)
	if err != nil {
		return id, fmt.Errorf("failed to finish job: %w", err)
	}

	// remove job number from assigned workers map
	err = mgr.assignedWorkers.Delete(jobNum)
	if err != nil {
		return id, fmt.Errorf("failed to finish job: %w", err)
	}

	return id, nil
}

func (mgr *SimpleJobMgr) Requeue(jobNum model.JobNumber) (model.WorkerId, error) {
	id := mgr.getAssignedWorkerId(jobNum)

	// remove job from run store
	j, err := mgr.run.GetAndDelete(jobNum)
	if err != nil {
		return id, fmt.Errorf("failed to requeue job: %w", err)
	}

	// and move to wait queue
	err = mgr.wait.Enqueue(j)
	if err != nil {
		return id, fmt.Errorf("failed to requeue job: %w", err)
	}

	// remove job number from assigned workers map
	err = mgr.assignedWorkers.Delete(jobNum)
	if err != nil {
		return id, fmt.Errorf("failed to requeue job: %w", err)
	}

	return id, nil
}

// MarkedForCancellation checks if a job is marked for cancellation
func (mgr *SimpleJobMgr) MarkedForCancellation(jobNum model.JobNumber) (bool, error) {
	marked, err := mgr.cancel.Exists(jobNum)
	if err != nil {
		return false, fmt.Errorf("error checking for job %d existence: %w", jobNum, err)
	}
	return marked, nil
}

// UnmarkForCancellation removes a job from the cancellation store
func (mgr *SimpleJobMgr) UnmarkForCancellation(jobNum model.JobNumber) error {
	err := mgr.cancel.Delete(jobNum)
	if err != nil {
		return fmt.Errorf("error unmarking job %d for cancellation", jobNum)
	}
	return nil
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

func (mgr *SimpleJobMgr) getJobList(options *pb.ListOptions, c container.Filterable[*pb.JobSpec]) *pb.JobList {
	jobSpecs := c.Filter(func(spec *pb.JobSpec) bool {
		pred := func() bool {
			if len(options.JobFilter.JobNums) == 0 {
				return true
			}

			for _, n := range options.JobFilter.JobNums {
				if spec.JobNum == n {
					return true
				}
			}
			return false
		}

		if options.JobFilter.JobKindFilter&int32(spec.Kind) != 0 && pred() {
			return true
		}
		return false
	})

	items := make([]*pb.JobListItem, 0, len(jobSpecs))
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

		items = append(items, item)
	}
	return &pb.JobList{Items: items}
}
