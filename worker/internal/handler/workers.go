package handler

import (
	"net"
	"encoding/json"
	"worker/internal/pool"
	"worker/internal/utils"
)

// WorkerStatus holds info about a worker pool for /workers endpoint
type WorkerStatus struct {
	Name         string `json:"name"`
	NumWorkers   int    `json:"num_workers"`
	QueueLength  int    `json:"queue_length"`
}

// HandleWorkers returns the status of all worker pools
func HandleWorkers(params map[string]string, conn net.Conn) {
	workers := []WorkerStatus{
		{"fibonacci", 5, len(pool.FibJobs)},
		{"createfile", 3, len(pool.CreateFileJobs)},
		{"deletefile", 3, len(pool.DeleteFileJobs)},
		{"reverse", 3, len(pool.ReverseJobs)},
		{"toupper", 3, len(pool.ToUpperJobs)},
		{"random", 3, len(pool.RandomJobs)},
		{"timestamp", 2, len(pool.TimestampJobs)},
		{"hash", 3, len(pool.HashJobs)},
		{"simulate", 3, len(pool.SimulateJobs)},
		{"sleep", 3, len(pool.SleepJobs)},
		{"loadtest", 3, len(pool.LoadTestJobs)},
		{"computepi", 2, len(pool.ComputePiJobs)},
		{"pow", 2, len(pool.PowJobs)},
		{"help", 1, len(pool.HelpJobs)},
	}
	data, _ := json.MarshalIndent(workers, "", "  ")
	utils.WriteHTTPResponse(conn, 200, "application/json", string(data))
}
