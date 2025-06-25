package pool

import (
	"worker/internal/commands"
)

type FibJob struct {
	N    int
	Resp chan FibResult
}
type FibResult struct {
	Value string
	Err   error
}

var FibJobs chan FibJob

func StartFibPool(workers, queueSize int) {
	FibJobs = make(chan FibJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range FibJobs {
				res, err := commands.Fibonacci(job.N)
				job.Resp <- FibResult{Value: res, Err: err}
			}
		}()
	}
}
