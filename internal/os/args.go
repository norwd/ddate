package os

import (
	original "os"
	"sync"
)

// Args hold the command-line arguments, starting with the program name.
var Args []string = original.Args

// argsMutex locks access to changing the Args variable.
var argsMutex sync.Mutex

// MockAndLockArgs mocks the argument list and returns its unlock hook.
func MockAndLockArgs(mockProgramName string, mockArguments []string) interface{ Unlock() } {
	argsMutex.Lock()
	Args = append([]string{mockProgramName}, mockArguments...)
	return &argsMutex
}
