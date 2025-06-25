package pool

import (
	"worker/internal/commands"
)

type ComputePiJob struct {
	Iters int
	Resp  chan ComputePiResult
}

type ComputePiResult struct {
	Value string
	Err   error
}

var ComputePiJobs chan ComputePiJob

func StartComputePiPool(workers, queueSize int) {
	ComputePiJobs = make(chan ComputePiJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range ComputePiJobs {
				res, err := commands.ComputePi(job.Iters)
				job.Resp <- ComputePiResult{Value: res, Err: err}
			}
		}()
	}
}
