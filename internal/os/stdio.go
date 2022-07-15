package os

import (
	"io"
	original "os"
	"sync"
)

// Stdin, Stdout, and Stderr are open Files pointing to the standard input,
// standard output, and standard error file descriptors.
//
// Note that the Go runtime writes to standard error for panics and crashes;
// closing Stderr may cause those messages to go elsewhere, perhaps to a file
// opened later.
var (
	Stdin  io.Reader = original.Stdin
	Stdout io.Writer = original.Stdout
	Stderr io.Writer = original.Stderr
)

// stdinMutex locks access to changing the Args variable.
var stdinMutex, stderrMutex, stdoutMutex sync.Mutex

// MockAndLockStdin mocks the input stream and returns its unlock hook.
func MockAndLockStdin(mock io.Reader) interface{ Unlock() } {
	stdinMutex.Lock()
	Stdin = mock
	return &stdinMutex
}

// MockAndLockStderr mocks the error stream and returns its unlock hook.
func MockAndLockStderr(mock io.Writer) interface{ Unlock() } {
	stderrMutex.Lock()
	Stderr = mock
	return &stderrMutex
}

// MockAndLockStdout mocks the output stream and returns its unlock hook.
func MockAndLockStdout(mock io.Writer) interface{ Unlock() } {
	stdoutMutex.Lock()
	Stdout = mock
	return &stdoutMutex
}
