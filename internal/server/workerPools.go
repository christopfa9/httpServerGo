package server

import (
	"httpServerGo/internal/commands"
)

// --- Fibonacci worker pool ---

type fibJob struct {
	n    int
	resp chan fibResult
}

type fibResult struct {
	value string
	err   error
}

var fibJobs chan fibJob

func initFibPool(workerCount, queueSize int) {
	fibJobs = make(chan fibJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range fibJobs {
				res, err := commands.Fibonacci(job.n)
				job.resp <- fibResult{value: res, err: err}
			}
		}()
	}
}

// --- CreateFile worker pool ---

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

func initCreateFilePool(workerCount, queueSize int) {
	createFileJobs = make(chan createFileJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range createFileJobs {
				res, err := commands.CreateFile(job.name, job.content, job.repeat)
				job.resp <- createFileResult{value: res, err: err}
			}
		}()
	}
}

// --- DeleteFile worker pool ---

type deleteFileJob struct {
	name string
	resp chan deleteFileResult
}

type deleteFileResult struct {
	value string
	err   error
}

var deleteFileJobs chan deleteFileJob

func initDeleteFilePool(workerCount, queueSize int) {
	deleteFileJobs = make(chan deleteFileJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range deleteFileJobs {
				res, err := commands.DeleteFile(job.name)
				job.resp <- deleteFileResult{value: res, err: err}
			}
		}()
	}
}

// --- Reverse worker pool ---

type reverseJob struct {
	text string
	resp chan reverseResult
}

type reverseResult struct {
	value string
	err   error
}

var reverseJobs chan reverseJob

func initReversePool(workerCount, queueSize int) {
	reverseJobs = make(chan reverseJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range reverseJobs {
				res, err := commands.Reverse(job.text)
				job.resp <- reverseResult{value: res, err: err}
			}
		}()
	}
}

// --- ToUpper worker pool ---

type toUpperJob struct {
	text string
	resp chan toUpperResult
}

type toUpperResult struct {
	value string
	err   error
}

var toUpperJobs chan toUpperJob

func initToUpperPool(workerCount, queueSize int) {
	toUpperJobs = make(chan toUpperJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range toUpperJobs {
				res, err := commands.ToUpper(job.text)
				job.resp <- toUpperResult{value: res, err: err}
			}
		}()
	}
}

// --- Random worker pool ---

type randomJob struct {
	count, min, max int
	resp            chan randomResult
}

type randomResult struct {
	value []int
	err   error
}

var randomJobs chan randomJob

func initRandomPool(workerCount, queueSize int) {
	randomJobs = make(chan randomJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range randomJobs {
				res, err := commands.Random(job.count, job.min, job.max)
				job.resp <- randomResult{value: res, err: err}
			}
		}()
	}
}

// --- Timestamp worker pool ---

type timestampJob struct {
	resp chan timestampResult
}

type timestampResult struct {
	value string
	err   error
}

var timestampJobs chan timestampJob

func initTimestampPool(workerCount, queueSize int) {
	timestampJobs = make(chan timestampJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range timestampJobs {
				res, err := commands.Timestamp()
				job.resp <- timestampResult{value: res, err: err}
			}
		}()
	}
}

// --- Hash worker pool ---

type hashJob struct {
	text string
	resp chan hashResult
}

type hashResult struct {
	value string
	err   error
}

var hashJobs chan hashJob

func initHashPool(workerCount, queueSize int) {
	hashJobs = make(chan hashJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range hashJobs {
				res, err := commands.Hash(job.text)
				job.resp <- hashResult{value: res, err: err}
			}
		}()
	}
}

// --- Simulate worker pool ---

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

func initSimulatePool(workerCount, queueSize int) {
	simulateJobs = make(chan simulateJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range simulateJobs {
				res, err := commands.Simulate(job.seconds, job.task)
				job.resp <- simulateResult{value: res, err: err}
			}
		}()
	}
}

// --- Sleep worker pool ---

type sleepJob struct {
	seconds int
	resp    chan sleepResult
}

type sleepResult struct {
	value string
	err   error
}

var sleepJobs chan sleepJob

func initSleepPool(workerCount, queueSize int) {
	sleepJobs = make(chan sleepJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range sleepJobs {
				res, err := commands.Sleep(job.seconds)
				job.resp <- sleepResult{value: res, err: err}
			}
		}()
	}
}

// --- LoadTest worker pool ---

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

func initLoadTestPool(workerCount, queueSize int) {
	loadTestJobs = make(chan loadTestJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range loadTestJobs {
				res, err := commands.LoadTest(job.tasks, job.sleepSec)
				job.resp <- loadTestResult{value: res, err: err}
			}
		}()
	}
}

// --- Help worker pool ---

type helpJob struct {
	resp chan helpResult
}

type helpResult struct {
	value string
	err   error
}

var helpJobs chan helpJob

func initHelpPool(workerCount, queueSize int) {
	helpJobs = make(chan helpJob, queueSize)
	for i := 0; i < workerCount; i++ {
		go func() {
			for job := range helpJobs {
				res, err := commands.Help()
				job.resp <- helpResult{value: res, err: err}
			}
		}()
	}
}

// InitWorkerPools arranca todos los pools con configuración por defecto:
// - 4 workers para comandos CPU-bound
// - 2 workers para comandos I/O-bound o triviales
// - cola de tamaño 100
func InitWorkerPools() {
	initFibPool(4, 100)
	initCreateFilePool(2, 100)
	initDeleteFilePool(2, 100)
	initReversePool(2, 100)
	initToUpperPool(2, 100)
	initRandomPool(2, 100)
	initTimestampPool(2, 100)
	initHashPool(2, 100)
	initSimulatePool(2, 100)
	initSleepPool(2, 100)
	initLoadTestPool(4, 100)
	initHelpPool(1, 100)
}
