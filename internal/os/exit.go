package os

import (
	original "os"
	"sync"
)

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicated success, non-zero an error. The program
// terminates immediately; deferred functions are not run.
//
// For portability, the status code should be in the range [0, 125].
func Exit(code int) {
	exitHook(code)
}

var exitHook func(int) = original.Exit

// exitMutex locks access to changing the exitHook variable.
var exitMutex sync.Mutex

// MockAndLockExit mocks the exit function and returns its unlock hook.
func MockAndLockExit(mock func(int)) interface{ Unlock() } {
	exitMutex.Lock()
	exitHook = mock
	return &exitMutex
}
