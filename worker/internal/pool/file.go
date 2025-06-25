package pool

import (
	"worker/internal/commands"
)

type CreateFileJob struct {
	Name    string
	Content string
	Repeat  int
	Resp    chan CreateFileResult
}
type CreateFileResult struct {
	Value string
	Err   error
}

var CreateFileJobs chan CreateFileJob

func StartCreateFilePool(workers, queueSize int) {
	CreateFileJobs = make(chan CreateFileJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range CreateFileJobs {
				res, err := commands.CreateFile(job.Name, job.Content, job.Repeat)
				job.Resp <- CreateFileResult{Value: res, Err: err}
			}
		}()
	}
}

type DeleteFileJob struct {
	Name string
	Resp chan DeleteFileResult
}
type DeleteFileResult struct {
	Value string
	Err   error
}

var DeleteFileJobs chan DeleteFileJob

func StartDeleteFilePool(workers, queueSize int) {
	DeleteFileJobs = make(chan DeleteFileJob, queueSize)
	for i := 0; i < workers; i++ {
		go func() {
			for job := range DeleteFileJobs {
				res, err := commands.DeleteFile(job.Name)
				job.Resp <- DeleteFileResult{Value: res, Err: err}
			}
		}()
	}
}
