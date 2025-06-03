// worker/internal/server/workerPools.go

package server

import (
	"worker/internal/commands"
)

// ——————————————————————————
// Estructuras y canales para cada tipo de trabajo
// ——————————————————————————

// Fibonacci
type fibJob struct {
	n    int
	resp chan fibResult
}
type fibResult struct {
	value string
	err   error
}

var fibJobs chan fibJob

// CreateFile
type createFileJob struct {
	name    string
	content string
	repeat  int
	resp    chan createFileResult
}
type createFileResult struct {
	value string
	err   error
}

var createFileJobs chan createFileJob

// DeleteFile
type deleteFileJob struct {
	name string
	resp chan deleteFileResult
}
type deleteFileResult struct {
	value string
	err   error
}

var deleteFileJobs chan deleteFileJob

// Reverse
type reverseJob struct {
	text string
	resp chan reverseResult
}
type reverseResult struct {
	value string
	err   error
}

var reverseJobs chan reverseJob

// ToUpper
type toUpperJob struct {
	text string
	resp chan toUpperResult
}
type toUpperResult struct {
	value string
	err   error
}

var toUpperJobs chan toUpperJob

// Random
type randomJob struct {
	count int
	min   int
	max   int
	resp  chan randomResult
}
type randomResult struct {
	value []int
	err   error
}

var randomJobs chan randomJob

// Timestamp
type timestampJob struct {
	resp chan timestampResult
}
type timestampResult struct {
	value string
	err   error
}

var timestampJobs chan timestampJob

// Hash
type hashJob struct {
	text string
	resp chan hashResult
}
type hashResult struct {
	value string
	err   error
}

var hashJobs chan hashJob

// Simulate
type simulateJob struct {
	seconds int
	task    string
	resp    chan simulateResult
}
type simulateResult struct {
	value string
	err   error
}

var simulateJobs chan simulateJob

// Sleep
type sleepJob struct {
	seconds int
	resp    chan sleepResult
}
type sleepResult struct {
	value string
	err   error
}

var sleepJobs chan sleepJob

// LoadTest
type loadTestJob struct {
	tasks    int
	sleepSec int
	resp     chan loadTestResult
}
type loadTestResult struct {
	value string
	err   error
}

var loadTestJobs chan loadTestJob

// ComputePi (nuevo)
type computePiJob struct {
	iters int
	resp  chan computePiResult
}
type computePiResult struct {
	value string
	err   error
}

var computePiJobs chan computePiJob

// Pow (nuevo)
type powJob struct {
	prefix    string
	maxTrials int
	resp      chan powResult
}
type powResult struct {
	value string
	err   error
}

var powJobs chan powJob

// Help
type helpJob struct {
	resp chan helpResult
}
type helpResult struct {
	value string
	err   error
}

var helpJobs chan helpJob

// ——————————————————————————
// InitWorkerPools arranca todos los pools de trabajo
// ——————————————————————————
func InitWorkerPools() {
	// Configura número de workers y tamaño de cola para cada pool
	const (
		fibWorkers       = 5
		createWorkers    = 3
		deleteWorkers    = 3
		reverseWorkers   = 3
		toUpperWorkers   = 3
		randomWorkers    = 3
		timestampWorkers = 2
		hashWorkers      = 3
		simulateWorkers  = 3
		sleepWorkers     = 3
		loadTestWorkers  = 3
		computePiWorkers = 2 // nuevos
		powWorkers       = 2 // nuevos

		queueSize = 50
	)

	// Inicializar canales de buffered jobs
	fibJobs = make(chan fibJob, queueSize)
	createFileJobs = make(chan createFileJob, queueSize)
	deleteFileJobs = make(chan deleteFileJob, queueSize)
	reverseJobs = make(chan reverseJob, queueSize)
	toUpperJobs = make(chan toUpperJob, queueSize)
	randomJobs = make(chan randomJob, queueSize)
	timestampJobs = make(chan timestampJob, queueSize)
	hashJobs = make(chan hashJob, queueSize)
	simulateJobs = make(chan simulateJob, queueSize)
	sleepJobs = make(chan sleepJob, queueSize)
	loadTestJobs = make(chan loadTestJob, queueSize)
	computePiJobs = make(chan computePiJob, queueSize)
	powJobs = make(chan powJob, queueSize)
	helpJobs = make(chan helpJob, queueSize)

	// Lanzar goroutines para cada pool

	// Fibonacci pool
	for i := 0; i < fibWorkers; i++ {
		go func() {
			for job := range fibJobs {
				res, err := commands.Fibonacci(job.n)
				job.resp <- fibResult{value: res, err: err}
			}
		}()
	}

	// CreateFile pool
	for i := 0; i < createWorkers; i++ {
		go func() {
			for job := range createFileJobs {
				res, err := commands.CreateFile(job.name, job.content, job.repeat)
				job.resp <- createFileResult{value: res, err: err}
			}
		}()
	}

	// DeleteFile pool
	for i := 0; i < deleteWorkers; i++ {
		go func() {
			for job := range deleteFileJobs {
				res, err := commands.DeleteFile(job.name)
				job.resp <- deleteFileResult{value: res, err: err}
			}
		}()
	}

	// Reverse pool
	for i := 0; i < reverseWorkers; i++ {
		go func() {
			for job := range reverseJobs {
				res, err := commands.Reverse(job.text)
				job.resp <- reverseResult{value: res, err: err}
			}
		}()
	}

	// ToUpper pool
	for i := 0; i < toUpperWorkers; i++ {
		go func() {
			for job := range toUpperJobs {
				res, err := commands.ToUpper(job.text)
				job.resp <- toUpperResult{value: res, err: err}
			}
		}()
	}

	// Random pool
	for i := 0; i < randomWorkers; i++ {
		go func() {
			for job := range randomJobs {
				res, err := commands.Random(job.count, job.min, job.max)
				job.resp <- randomResult{value: res, err: err}
			}
		}()
	}

	// Timestamp pool
	for i := 0; i < timestampWorkers; i++ {
		go func() {
			for job := range timestampJobs {
				res, err := commands.Timestamp()
				job.resp <- timestampResult{value: res, err: err}
			}
		}()
	}

	// Hash pool (corregido)
	for i := 0; i < hashWorkers; i++ {
		go func() {
			for job := range hashJobs {
				res, err := commands.Hash(job.text)
				job.resp <- hashResult{value: res, err: err}
			}
		}()
	}

	// Simulate pool
	for i := 0; i < simulateWorkers; i++ {
		go func() {
			for job := range simulateJobs {
				res, err := commands.Simulate(job.seconds, job.task)
				job.resp <- simulateResult{value: res, err: err}
			}
		}()
	}

	// Sleep pool
	for i := 0; i < sleepWorkers; i++ {
		go func() {
			for job := range sleepJobs {
				res, err := commands.Sleep(job.seconds)
				job.resp <- sleepResult{value: res, err: err}
			}
		}()
	}

	// LoadTest pool
	for i := 0; i < loadTestWorkers; i++ {
		go func() {
			for job := range loadTestJobs {
				res, err := commands.LoadTest(job.tasks, job.sleepSec)
				job.resp <- loadTestResult{value: res, err: err}
			}
		}()
	}

	// ComputePi pool (nuevo)
	for i := 0; i < computePiWorkers; i++ {
		go func() {
			for job := range computePiJobs {
				res, err := commands.ComputePi(job.iters)
				job.resp <- computePiResult{value: res, err: err}
			}
		}()
	}

	// Pow pool (nuevo)
	for i := 0; i < powWorkers; i++ {
		go func() {
			for job := range powJobs {
				res, err := commands.Pow(job.prefix, job.maxTrials)
				job.resp <- powResult{value: res, err: err}
			}
		}()
	}

	// Help pool
	for i := 0; i < 1; i++ {
		go func() {
			for job := range helpJobs {
				res, err := commands.Help()
				job.resp <- helpResult{value: res, err: err}
			}
		}()
	}
}
