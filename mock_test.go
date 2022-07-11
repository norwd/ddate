package main

import (
	"io"
	"sync"
	"time"
)

// Lock the injectable dependencies only allow sequential access.
var backendMutex, exitMutex, selfMutex, argsMutex, stderrMutex, stdoutMutex sync.Mutex

// mockAndLockBackend mocks the backend function and returns its unlock hook.
func mockAndLockBackend(mock func(string, time.Time) (string, error)) interface{ Unlock() } {
	backendMutex.Lock()
	backend = mock
	return &backendMutex
}

// mockAndLockExit mocks the exit function and returns its unlock hook.
func mockAndLockExit(mock func(int)) interface{ Unlock() } {
	exitMutex.Lock()
	exit = mock
	return &exitMutex
}

// mockAndLockSelf mocks the self name and returns its unlock hook.
func mockAndLockSelf(mock string) interface{ Unlock() } {
	selfMutex.Lock()
	self = mock
	return &selfMutex
}

// mockAndLockArgs mocks the argument list and returns its unlock hook.
func mockAndLockArgs(mock []string) interface{ Unlock() } {
	argsMutex.Lock()
	args = mock
	return &argsMutex
}

// mockAndLockStderr mocks the error stream and returns its unlock hook.
func mockAndLockStderr(mock io.Writer) interface{ Unlock() } {
	stderrMutex.Lock()
	stderr = mock
	return &stderrMutex
}

// mockAndLockStdout mocks the output stream and returns its unlock hook.
func mockAndLockStdout(mock io.Writer) interface{ Unlock() } {
	stdoutMutex.Lock()
	stdout = mock
	return &stdoutMutex
}
