package pool

import (
	"worker/internal/commands"
)

type SleepJob struct {
	Seconds int
	Resp    chan SleepResult
}

type SleepResult struct {
	Value string
	Err   error
}

var SleepJobs chan SleepJob

func StartSleepPool(workers, queueSize int) {
	SleepJobs = make(chan SleepJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range SleepJobs {
				res, err := commands.Sleep(job.Seconds)
				job.Resp <- SleepResult{Value: res, Err: err}
			}
		}()
	}
}
