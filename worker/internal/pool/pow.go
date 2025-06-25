package pool

import (
	"worker/internal/commands"
)

type PowJob struct {
	Prefix    string
	MaxTrials int
	Resp      chan PowResult
}

type PowResult struct {
	Value string
	Err   error
}

var PowJobs chan PowJob

func StartPowPool(workers, queueSize int) {
	PowJobs = make(chan PowJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range PowJobs {
				res, err := commands.Pow(job.Prefix, job.MaxTrials)
				job.Resp <- PowResult{Value: res, Err: err}
			}
		}()
	}
}
