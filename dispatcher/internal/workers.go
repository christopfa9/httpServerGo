package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type StatusPayload struct {
	Hostname         string        `json:"hostname"`
	StartTime        string        `json:"start_time"`
	TotalConnections int           `json:"total_connections"`
	ActiveHandlers   int           `json:"active_handlers"`
	Processes        []interface{} `json:"processes"`
	CompletedTasks   int           `json:"completed_tasks"`
	QueueLength      int           `json:"queue_length"`
}

type WorkerStatusReport struct {
	Worker  string        `json:"worker"`
	Status  *StatusPayload `json:"status,omitempty"`
	Error   string        `json:"error,omitempty"`
}

// GetWorkersStatus contacts all workers and aggregates their /status endpoint
func GetWorkersStatus() ([]WorkerStatusReport, error) {
	workersEnv := os.Getenv("WORKERS")
	if workersEnv == "" {
		return nil, fmt.Errorf("WORKERS env variable not set")
	}
	workerAddrs := strings.Split(workersEnv, ",")
	var reports []WorkerStatusReport
	for _, addr := range workerAddrs {
		url := fmt.Sprintf("http://%s/status", addr)
		resp, err := http.Get(url)
		if err != nil {
			reports = append(reports, WorkerStatusReport{Worker: addr, Error: err.Error()})
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			reports = append(reports, WorkerStatusReport{Worker: addr, Error: err.Error()})
			continue
		}
		var status StatusPayload
		if err := json.Unmarshal(body, &status); err != nil {
			reports = append(reports, WorkerStatusReport{Worker: addr, Error: err.Error()})
			continue
		}
		reports = append(reports, WorkerStatusReport{Worker: addr, Status: &status})
	}
	return reports, nil
}
