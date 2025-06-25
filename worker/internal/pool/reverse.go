package pool

import (
	"worker/internal/commands"
)

type ReverseJob struct {
	Text string
	Resp chan ReverseResult
}

type ReverseResult struct {
	Value string
	Err   error
}

var ReverseJobs chan ReverseJob

func StartReversePool(workers, queueSize int) {
	ReverseJobs = make(chan ReverseJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range ReverseJobs {
				res, err := commands.Reverse(job.Text)
				job.Resp <- ReverseResult{Value: res, Err: err}
			}
		}()
	}
}
