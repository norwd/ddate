package main // import "github.com/norwd/ddate"

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// defaultFormat is used if no format is explicitly given.
const defaultFormat = "%A, %B %d, %Y YOLD"

// Injectable dependencies that can be swapped out in testing.
var (
	// backend is the date formatting implementation.
	backend = func(string, time.Time) (string, error) {
		panic("backend not defined")
	}

	// exit is the exit call hook.
	exit = os.Exit

	// self is the invocation name.
	self = filepath.Base(os.Args[0])

	// args are the arguments passed to the system.
	args = os.Args[1:]

	// stderr is a writeable error stream.
	stderr io.Writer = os.Stderr

	// stdout is a writeable output stream.
	stdout io.Writer = os.Stdout
)

// errorf prints the formatted error message to stderr and exits with error 1.
func errorf(format string, args ...interface{}) {
	err := fmt.Sprintf(format, args...)

	fmt.Fprintln(stderr, err)

	exit(1)
}

// println prints a line to the given output stream.
func println(line string) {
	fmt.Fprintln(stdout, line)
}

// parseDDMMYYYY parses strings representing a day, month, and year as a time.
//
// Note that the time returned will normalise the day, month, and year values if
// they are outside their allowed range. E.g. Oct 32 becomes Nov 1.
func parseDDMMYYYY(dayStr, monthStr, yearStr string) (t time.Time, err error) {
	var day, month, year, hour, min, sec, nsec int

	if day, err = strconv.Atoi(dayStr); err != nil {
		return
	}

	if month, err = strconv.Atoi(monthStr); err != nil {
		return
	}

	if year, err = strconv.Atoi(yearStr); err != nil {
		return
	}

	t = time.Date(year, time.Month(month), day, hour, min, sec, nsec, time.Local)

	return
}

func main() {
	// Get the default values
	format, date := defaultFormat, time.Now()

	// Determine date format
	if argc := len(args); argc > 1 && strings.HasPrefix("+", args[0]) {
		// Trim the plus sing from the format
		format = strings.TrimPrefix(args[0], "+")

		// Reslice arguments to skip the format arguments
		args = args[1:]
	}

	// Determine date to use
	if argc := len(args); argc == 3 {
		var err error

		if date, err = parseDDMMYYYY(args[0], args[1], args[2]); err != nil {
			errorf("%s: %s", self, err)
		}
	} else if argc > 3 {
		errorf("%s: too many arguments for DD MM YYYY", self)
	} else if argc < 0 {
		errorf("%s: not enough arguments for DD MM YYYY", self)
	}

	// Format the date conversion
	if date, err := backend(format, date); err != nil {
		errorf("%s: %s", self, err)
	} else {
		fmt.Println(date)
	}
}
