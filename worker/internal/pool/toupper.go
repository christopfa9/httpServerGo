package pool

import (
	"worker/internal/commands"
)

type ToUpperJob struct {
	Text string
	Resp chan ToUpperResult
}

type ToUpperResult struct {
	Value string
	Err   error
}

var ToUpperJobs chan ToUpperJob

func StartToUpperPool(workers, queueSize int) {
	ToUpperJobs = make(chan ToUpperJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range ToUpperJobs {
				res, err := commands.ToUpper(job.Text)
				job.Resp <- ToUpperResult{Value: res, Err: err}
			}
		}()
	}
}
