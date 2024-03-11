package pb

import (
	"fmt"
	"log"
)

func (x *JobStatus) LogString() string {
	return fmt.Sprintf("%s jobNum: %d progress: %.0f%% msg: %s\n", x.JobState, x.JobNum, x.Progress*100.0, x.Msg)
}

func (x JobState) ShortDesc() string {
	var s string
	switch x {
	case JobState_JOB_STATE_NONE:
		s = "none"
	case JobState_JOB_STATE_WAIT:
		s = "wait"
	case JobState_JOB_STATE_RUN:
		s = "run"
	case JobState_JOB_STATE_DONE_OK:
		s = "ok"
	case JobState_JOB_STATE_DONE_ERR:
		s = "err"
	case JobState_JOB_STATE_DONE_CANCEL:
		s = "cancel"
	default:
		log.Fatalf("Unexpected job state: %v", x)
	}
	return s
}

func (x JobKind) ShortDesc() string {
	var s string
	switch x {
	case JobKind_JOB_KIND_NULL:
		s = "null"
	case JobKind_JOB_KIND_TEST:
		s = "test"
	case JobKind_JOB_KIND_SLEEP:
		s = "sleep"
	case JobKind_JOB_KIND_FFMPEG:
		s = "ffmpeg"
	default:
		log.Fatalf("Unexpected job kind: %v", x)
	}
	return s
}
