package status_test

import (
	"encoding/json"
	"strings"
	"testing"

	"httpServerGo/internal/status"
)

func TestMetrics_IncrementConnections(t *testing.T) {
	m := status.InitMetrics()

	initial := m.TotalConnections
	status.IncTotalConnections()
	if m.TotalConnections != initial+1 {
		t.Errorf("TotalConnections expected %d, got %d", initial+1, m.TotalConnections)
	}
}

func TestMetrics_ActiveHandlers(t *testing.T) {
	m := status.InitMetrics()

	initial := m.ActiveHandlers
	status.IncActiveHandlers()
	if m.ActiveHandlers != initial+1 {
		t.Errorf("ActiveHandlers expected %d, got %d", initial+1, m.ActiveHandlers)
	}
	status.DecActiveHandlers()
	if m.ActiveHandlers != initial {
		t.Errorf("ActiveHandlers expected %d after decrement, got %d", initial, m.ActiveHandlers)
	}
}

func TestMetrics_ProcessRegisterAndStatus(t *testing.T) {
	m := status.InitMetrics()

	pid := 1234
	cmd := "testcmd"

	status.RegisterProcess(pid, cmd)
	// Check if process registered
	proc, ok := m.ActiveProcesses[pid]
	if !ok {
		t.Fatal("Process not registered")
	}
	if proc.Command != cmd || proc.Status != "idle" {
		t.Errorf("Process fields mismatch. Got %+v", proc)
	}

	// Change status
	status.SetProcessStatus(pid, "busy")
	proc2 := m.ActiveProcesses[pid]
	if proc2.Status != "busy" {
		t.Errorf("Expected status busy, got %s", proc2.Status)
	}
}

func TestMetrics_MarshalJSON(t *testing.T) {
	status.InitMetrics()
	status.IncTotalConnections()
	status.IncActiveHandlers()
	status.RegisterProcess(1, "cmd1")

	jsonBytes, err := status.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	jsonStr := string(jsonBytes)
	if !strings.Contains(jsonStr, `"total_connections": 1`) {
		t.Errorf("JSON missing total_connections")
	}
	if !strings.Contains(jsonStr, `"active_handlers": 1`) {
		t.Errorf("JSON missing active_handlers")
	}
	if !strings.Contains(jsonStr, `"command": "cmd1"`) {
		t.Errorf("JSON missing process command")
	}

	// Verify valid JSON
	var payload map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &payload); err != nil {
		t.Errorf("JSON invalid: %v", err)
	}
}
