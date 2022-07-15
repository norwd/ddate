package main

import (
	"sync"
	"time"
)

// Lock the injectable dependencies only allow sequential access.
var backendMutex sync.Mutex

// mockAndLockBackend mocks the backend function and returns its unlock hook.
func mockAndLockBackend(mock func(string, time.Time) (string, error)) interface{ Unlock() } {
	backendMutex.Lock()
	backend = mock
	return &backendMutex
}
