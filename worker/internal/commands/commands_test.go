package commands_test

import (
	"strings"
	"testing"
	"time"

	"worker/internal/commands"
)

// --- Fibonacci ---

func TestFibonacci_Valid(t *testing.T) {
	res, err := commands.Fibonacci(10)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if res != "55" {
		t.Errorf("Expected '55', got %q", res)
	}
}

func TestFibonacci_Negative(t *testing.T) {
	_, err := commands.Fibonacci(-1)
	if err == nil {
		t.Error("Expected error for negative input")
	}
}

// --- CreateFile ---

func TestCreateFile_Valid(t *testing.T) {
	msg, err := commands.CreateFile("testfile.txt", "hello", 2)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !strings.Contains(msg, "creado/truncado con éxito") {
		t.Errorf("Unexpected message: %s", msg)
	}
}

func TestCreateFile_InvalidName(t *testing.T) {
	_, err := commands.CreateFile("../badname.txt", "hi", 1)
	if err == nil {
		t.Error("Expected error for invalid filename")
	}
}

func TestCreateFile_RepeatZero(t *testing.T) {
	msg, err := commands.CreateFile("testfile2.txt", "x", 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !strings.Contains(msg, "creado/truncado con éxito") {
		t.Errorf("Unexpected message: %s", msg)
	}
}

// --- DeleteFile ---

func TestDeleteFile_Valid(t *testing.T) {
	// Create a file first to delete
	_, err := commands.CreateFile("todelete.txt", "data", 1)
	if err != nil {
		t.Fatalf("Setup error creating file: %v", err)
	}

	msg, err := commands.DeleteFile("todelete.txt")
	if err != nil {
		t.Fatalf("Unexpected error deleting file: %v", err)
	}
	if !strings.Contains(msg, "eliminado con éxito") {
		t.Errorf("Unexpected message: %s", msg)
	}
}

func TestDeleteFile_NotExist(t *testing.T) {
	_, err := commands.DeleteFile("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}

func TestDeleteFile_InvalidName(t *testing.T) {
	_, err := commands.DeleteFile("../badname.txt")
	if err == nil {
		t.Error("Expected error for invalid filename")
	}
}

// --- Hash ---

func TestHash_Valid(t *testing.T) {
	res, err := commands.Hash("test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	expectedPrefix := "9f86d081"
	if !strings.HasPrefix(res, expectedPrefix) {
		t.Errorf("Unexpected hash result %s", res)
	}
}

func TestHash_Empty(t *testing.T) {
	_, err := commands.Hash("")
	if err == nil {
		t.Error("Expected error for empty input")
	}
}

// --- Help ---

func TestHelp(t *testing.T) {
	res, err := commands.Help()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !strings.Contains(res, "/fibonacci?num={N}") {
		t.Error("Help text missing /fibonacci command")
	}
}

// --- LoadTest ---

func TestLoadTest_Valid(t *testing.T) {
	msg, err := commands.LoadTest(2, 1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !strings.Contains(msg, "Executed 2 concurrent tasks sleeping 1 seconds each") {
		t.Errorf("Unexpected message: %s", msg)
	}
}

func TestLoadTest_InvalidTasks(t *testing.T) {
	_, err := commands.LoadTest(0, 1)
	if err == nil {
		t.Error("Expected error for zero tasks")
	}
}

func TestLoadTest_InvalidSleep(t *testing.T) {
	_, err := commands.LoadTest(1, -1)
	if err == nil {
		t.Error("Expected error for negative sleep")
	}
}

// --- Random ---

func TestRandom_Valid(t *testing.T) {
	nums, err := commands.Random(5, 1, 10)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(nums) != 5 {
		t.Errorf("Expected 5 numbers, got %d", len(nums))
	}
	for _, v := range nums {
		if v < 1 || v > 10 {
			t.Errorf("Number %d out of range", v)
		}
	}
}

func TestRandom_InvalidCount(t *testing.T) {
	_, err := commands.Random(0, 1, 10)
	if err == nil {
		t.Error("Expected error for count=0")
	}
}

func TestRandom_MinGreaterThanMax(t *testing.T) {
	_, err := commands.Random(5, 10, 1)
	if err == nil {
		t.Error("Expected error for min > max")
	}
}

// --- Reverse ---

func TestReverse(t *testing.T) {
	out, err := commands.Reverse("abcd")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if out != "dcba" {
		t.Errorf("Expected 'dcba', got '%s'", out)
	}
}

// --- Simulate ---

func TestSimulate_Valid(t *testing.T) {
	msg, err := commands.Simulate(1, "task1")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !strings.Contains(msg, "task1") {
		t.Errorf("Expected message to contain 'task1', got '%s'", msg)
	}
}

func TestSimulate_InvalidSeconds(t *testing.T) {
	_, err := commands.Simulate(-1, "task1")
	if err == nil {
		t.Error("Expected error for negative seconds")
	}
}

// --- Sleep ---

func TestSleep_Valid(t *testing.T) {
	msg, err := commands.Sleep(1)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !strings.Contains(msg, "Slept for 1 seconds") {
		t.Errorf("Unexpected message: %s", msg)
	}
}

func TestSleep_InvalidSeconds(t *testing.T) {
	_, err := commands.Sleep(-1)
	if err == nil {
		t.Error("Expected error for negative seconds")
	}
}

// --- Timestamp ---

func TestTimestamp(t *testing.T) {
	ts, err := commands.Timestamp()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	// Example ISO8601 format: 2006-01-02T15:04:05Z07:00
	if len(ts) < len("2006-01-02T15:04:05Z") {
		t.Errorf("Timestamp too short: %s", ts)
	}
	// Optional: parse timestamp to verify format correctness
	if _, err := time.Parse(time.RFC3339, ts); err != nil {
		t.Errorf("Timestamp format invalid: %v", err)
	}
}

// --- ToUpper ---

func TestToUpper_Valid(t *testing.T) {
	out, err := commands.ToUpper("hello")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if out != "HELLO" {
		t.Errorf("Expected 'HELLO', got '%s'", out)
	}
}

func TestToUpper_Empty(t *testing.T) {
	out, err := commands.ToUpper("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if out != "" {
		t.Errorf("Expected empty string, got '%s'", out)
	}
}
