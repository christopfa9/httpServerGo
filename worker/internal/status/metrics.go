package status

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// ProcessInfo stores information about a child process.
type ProcessInfo struct {
	PID     int    `json:"pid"`
	Command string `json:"command"`
	Status  string `json:"status"` // "busy" or "idle"
}

// ServerMetrics holds runtime metrics for the server.
type ServerMetrics struct {
	StartTime        time.Time           `json:"start_time"`
	TotalConnections int                 `json:"total_connections"`
	ActiveHandlers   int                 `json:"active_handlers"`
	ActiveProcesses  map[int]ProcessInfo `json:"-"`
	mutex            sync.Mutex          // protects all the above fields
}

var metrics *ServerMetrics

// InitMetrics creates and returns the singleton for server metrics.
// Automatically called in init().
func InitMetrics() *ServerMetrics {
	m := &ServerMetrics{
		StartTime:       time.Now(),
		ActiveProcesses: make(map[int]ProcessInfo),
	}
	metrics = m
	return m
}

func init() {
	InitMetrics()
}

// IncTotalConnections increments the total connections counter.
func IncTotalConnections() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.TotalConnections++
}

// IncActiveHandlers increments the active handlers counter.
func IncActiveHandlers() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.ActiveHandlers++
}

// DecActiveHandlers decrements the active handlers counter.
func DecActiveHandlers() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.ActiveHandlers--
}

// RegisterProcess adds a process with "idle" status.
func RegisterProcess(pid int, command string) {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.ActiveProcesses[pid] = ProcessInfo{
		PID:     pid,
		Command: command,
		Status:  "idle",
	}
}

// SetProcessStatus updates the status ("busy"/"idle") of a registered process.
func SetProcessStatus(pid int, statusStr string) {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	if p, ok := metrics.ActiveProcesses[pid]; ok {
		p.Status = statusStr
		metrics.ActiveProcesses[pid] = p
	}
}

// Marshal serializes all metrics to JSON, including the hostname.
func Marshal() ([]byte, error) {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()

	// Convert the map to a slice for JSON
	procs := make([]ProcessInfo, 0, len(metrics.ActiveProcesses))
	for _, pi := range metrics.ActiveProcesses {
		procs = append(procs, pi)
	}

	// Get the server hostname
	host, _ := os.Hostname()

	// Payload in the format we want to expose
	payload := struct {
		Hostname         string        `json:"hostname"`
		StartTime        time.Time     `json:"start_time"`
		TotalConnections int           `json:"total_connections"`
		ActiveHandlers   int           `json:"active_handlers"`
		Processes        []ProcessInfo `json:"processes"`
	}{
		Hostname:         host,
		StartTime:        metrics.StartTime,
		TotalConnections: metrics.TotalConnections,
		ActiveHandlers:   metrics.ActiveHandlers,
		Processes:        procs,
	}

	return json.MarshalIndent(payload, "", "  ")
}
