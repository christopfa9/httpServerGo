package pool

import (
	"worker/internal/commands"
)

type HelpJob struct {
	Resp chan HelpResult
}

type HelpResult struct {
	Value string
	Err   error
}

var HelpJobs chan HelpJob

func StartHelpPool(workers, queueSize int) {
	HelpJobs = make(chan HelpJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range HelpJobs {
				res, err := commands.Help()
				job.Resp <- HelpResult{Value: res, Err: err}
			}
		}()
	}
}
