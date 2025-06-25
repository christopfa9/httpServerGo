package pool

import (
	"worker/internal/commands"
)

type SimulateJob struct {
	Seconds int
	Task    string
	Resp    chan SimulateResult
}

type SimulateResult struct {
	Value string
	Err   error
}

var SimulateJobs chan SimulateJob

func StartSimulatePool(workers, queueSize int) {
	SimulateJobs = make(chan SimulateJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range SimulateJobs {
				res, err := commands.Simulate(job.Seconds, job.Task)
				job.Resp <- SimulateResult{Value: res, Err: err}
			}
		}()
	}
}
