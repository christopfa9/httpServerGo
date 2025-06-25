package pool

import (
	"worker/internal/commands"
)

type RandomJob struct {
	Count int
	Min   int
	Max   int
	Resp  chan RandomResult
}

type RandomResult struct {
	Value []int
	Err   error
}

var RandomJobs chan RandomJob

func StartRandomPool(workers, queueSize int) {
	RandomJobs = make(chan RandomJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range RandomJobs {
				res, err := commands.Random(job.Count, job.Min, job.Max)
				job.Resp <- RandomResult{Value: res, Err: err}
			}
		}()
	}
}
