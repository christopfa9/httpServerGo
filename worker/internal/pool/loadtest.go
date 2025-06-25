package pool

import (
	"worker/internal/commands"
)

type LoadTestJob struct {
	Tasks    int
	SleepSec int
	Resp     chan LoadTestResult
}

type LoadTestResult struct {
	Value string
	Err   error
}

var LoadTestJobs chan LoadTestJob

func StartLoadTestPool(workers, queueSize int) {
	LoadTestJobs = make(chan LoadTestJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range LoadTestJobs {
				res, err := commands.LoadTest(job.Tasks, job.SleepSec)
				job.Resp <- LoadTestResult{Value: res, Err: err}
			}
		}()
	}
}
