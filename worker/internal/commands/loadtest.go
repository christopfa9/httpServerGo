package commands

import (
	"fmt"
	"sync"
	"time"
)

// LoadTest launches 'tasks' concurrent goroutines that each sleep for 'sleepSec' seconds.
// Waits for all of them to finish before returning a confirmation message.
func LoadTest(tasks, sleepSec int) (string, error) {
	// 1) Validate parameters
	if tasks <= 0 {
		return "", fmt.Errorf("the 'tasks' parameter must be > 0, received %d", tasks)
	}
	if sleepSec < 0 {
		return "", fmt.Errorf("the 'sleep' parameter must be >= 0, received %d", sleepSec)
	}

	// 2) Launch concurrent goroutines with WaitGroup
	var wg sync.WaitGroup
	wg.Add(tasks)
	for i := 0; i < tasks; i++ {
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(sleepSec) * time.Second)
		}(i)
	}

	// 3) Wait for all to finish
	wg.Wait()

	// 4) Build confirmation message
	msg := fmt.Sprintf("Executed %d concurrent tasks sleeping %d seconds each", tasks, sleepSec)
	return msg, nil
}
