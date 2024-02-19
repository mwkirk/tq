package pb

import "fmt"

func (js *JobStatus) LogString() string {
	return fmt.Sprintf("%s jobNum: %d progress: %.0f%% msg: %s\n", js.JobState, js.Num, js.Progress*100.0, js.Msg)
}
