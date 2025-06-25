package pool

import (
	"worker/internal/commands"
)

type TimestampJob struct {
	Resp chan TimestampResult
}

type TimestampResult struct {
	Value string
	Err   error
}

var TimestampJobs chan TimestampJob

func StartTimestampPool(workers, queueSize int) {
	TimestampJobs = make(chan TimestampJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range TimestampJobs {
				res, err := commands.Timestamp()
				job.Resp <- TimestampResult{Value: res, Err: err}
			}
		}()
	}
}
