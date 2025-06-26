package status

import (
	"encoding/json"
	"os"
	"sync"
	"time"

	"worker/internal/pool"
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

	// Calcula la carga (suma de la longitud de todas las colas de trabajo)
	queueLength := 0
	queueLength += len(pool.FibJobs)
	queueLength += len(pool.CreateFileJobs)
	queueLength += len(pool.DeleteFileJobs)
	queueLength += len(pool.ReverseJobs)
	queueLength += len(pool.ToUpperJobs)
	queueLength += len(pool.RandomJobs)
	queueLength += len(pool.TimestampJobs)
	queueLength += len(pool.HashJobs)
	queueLength += len(pool.SimulateJobs)
	queueLength += len(pool.SleepJobs)
	queueLength += len(pool.LoadTestJobs)
	queueLength += len(pool.ComputePiJobs)
	queueLength += len(pool.PowJobs)
	queueLength += len(pool.HelpJobs)

	// Payload in the format we want to expose
	payload := struct {
		Hostname         string        `json:"hostname"`
		StartTime        time.Time     `json:"start_time"`
		TotalConnections int           `json:"total_connections"`
		ActiveHandlers   int           `json:"active_handlers"`
		Processes        []ProcessInfo `json:"processes"`
		CompletedTasks   int           `json:"completed_tasks"`
		QueueLength      int           `json:"queue_length"`
	}{
		Hostname:         host,
		StartTime:        metrics.StartTime,
		TotalConnections: metrics.TotalConnections,
		ActiveHandlers:   metrics.ActiveHandlers,
		Processes:        procs,
		CompletedTasks:   completedTasks,
		QueueLength:      queueLength,
	}

	return json.MarshalIndent(payload, "", "  ")
}

// completedTasks cuenta el total de tareas completadas por el worker.
var completedTasks int

// IncCompletedTasks incrementa el contador de tareas completadas.
func IncCompletedTasks() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	completedTasks++
}

// GetCompletedTasks devuelve el total de tareas completadas.
func GetCompletedTasks() int {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	return completedTasks
}
