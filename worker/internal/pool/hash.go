package pool

import (
	"worker/internal/commands"
)

type HashJob struct {
	Text string
	Resp chan HashResult
}

type HashResult struct {
	Value string
	Err   error
}

var HashJobs chan HashJob

func StartHashPool(workers, queueSize int) {
	HashJobs = make(chan HashJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range HashJobs {
				res, err := commands.Hash(job.Text)
				job.Resp <- HashResult{Value: res, Err: err}
			}
		}()
	}
}
