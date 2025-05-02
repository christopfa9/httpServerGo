package status

// TODO: Implement metrics.go (Tracks server status and runtime metrics)
//
// [ ] Import necessary packages:
//     - sync, time, os
//
// [ ] Define a struct ServerMetrics with fields:
//     - StartTime time.Time
//     - TotalConnections int
//     - Mutex sync.Mutex
//     - ActiveProcesses map[int]ProcessInfo
//
// [ ] Define a struct ProcessInfo with fields:
//     - PID int
//     - Command string
//     - Status string ("busy", "idle")
//
// [ ] Initialize metrics at server startup
//
// [ ] Implement functions:
//     - InitMetrics() *ServerMetrics
//     - IncrementConnections()
//     - RegisterProcess(pid int, command string)
//     - SetProcessStatus(pid int, status string)
//     - GetStatusJSON() ([]byte, error)
//
// [ ] Ensure thread-safe access using sync.Mutex
//
// [ ] Use GetStatusJSON() in /status handler to return JSON-formatted status

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// ProcessInfo almacena información sobre un proceso hijo.
type ProcessInfo struct {
	PID     int    `json:"pid"`
	Command string `json:"command"`
	Status  string `json:"status"` // "busy" o "idle"
}

// ServerMetrics mantiene las métricas del servidor en ejecución.
type ServerMetrics struct {
	StartTime        time.Time           `json:"start_time"`
	TotalConnections int                 `json:"total_connections"`
	ActiveHandlers   int                 `json:"active_handlers"`
	ActiveProcesses  map[int]ProcessInfo `json:"-"`
	mutex            sync.Mutex          // protege todos los campos anteriores
}

var metrics *ServerMetrics

// InitMetrics crea y devuelve el singleton de métricas.
// Se llama automáticamente en init().
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

// IncTotalConnections incrementa el contador de conexiones totales.
func IncTotalConnections() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.TotalConnections++
}

// IncActiveHandlers incrementa el contador de handlers activos.
func IncActiveHandlers() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.ActiveHandlers++
}

// DecActiveHandlers decrementa el contador de handlers activos.
func DecActiveHandlers() {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.ActiveHandlers--
}

// RegisterProcess añade un proceso con estado "idle".
func RegisterProcess(pid int, command string) {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	metrics.ActiveProcesses[pid] = ProcessInfo{
		PID:     pid,
		Command: command,
		Status:  "idle",
	}
}

// SetProcessStatus actualiza el estado ("busy"/"idle") de un proceso ya registrado.
func SetProcessStatus(pid int, statusStr string) {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()
	if p, ok := metrics.ActiveProcesses[pid]; ok {
		p.Status = statusStr
		metrics.ActiveProcesses[pid] = p
	}
}

// Marshal serializa todas las métricas a JSON, incluyendo el nombre de host.
func Marshal() ([]byte, error) {
	metrics.mutex.Lock()
	defer metrics.mutex.Unlock()

	// Convertimos el map a slice para JSON
	procs := make([]ProcessInfo, 0, len(metrics.ActiveProcesses))
	for _, pi := range metrics.ActiveProcesses {
		procs = append(procs, pi)
	}

	// Obtenemos hostname del servidor
	host, _ := os.Hostname()

	// Payload con la forma que queremos exponer
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
