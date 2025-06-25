package server

import "worker/internal/pool"

// InitWorkerPools arranca todos los pools de trabajo con su configuración.
func InitWorkerPools() {
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
		computePiWorkers = 2
		powWorkers       = 2
		helpWorkers      = 1
		queueSize        = 50
	)

	pool.StartFibPool(fibWorkers, queueSize)
	pool.StartCreateFilePool(createWorkers, queueSize)
	pool.StartDeleteFilePool(deleteWorkers, queueSize)
	pool.StartReversePool(reverseWorkers, queueSize)
	pool.StartToUpperPool(toUpperWorkers, queueSize)
	pool.StartRandomPool(randomWorkers, queueSize)
	pool.StartTimestampPool(timestampWorkers, queueSize)
	pool.StartHashPool(hashWorkers, queueSize)
	pool.StartSimulatePool(simulateWorkers, queueSize)
	pool.StartSleepPool(sleepWorkers, queueSize)
	pool.StartLoadTestPool(loadTestWorkers, queueSize)
	pool.StartComputePiPool(computePiWorkers, queueSize)
	pool.StartPowPool(powWorkers, queueSize)
	pool.StartHelpPool(helpWorkers, queueSize)
}
