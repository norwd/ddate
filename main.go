package main // import "github.com/norwd/ddate"

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// defaultFormat is used if no format is explicitly given.
const defaultFormat = "%A, %B %d, %Y YOLD"

// backend is a runtime swappable date formatting implementation.
var backend = func(string, time.Time) (string, error) {
	panic("backend not defined")
}

// errorf prints the formatted error message to stderr and exits with error 1.
func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
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

	// Get the invocation
	self := os.Args[0]
	args := os.Args[1:]

	// Trim the leading directory
	self = filepath.Base(self)

	// Debug
	fmt.Printf("self: %q, args: %#v\n", self, args)

	// Determine arg actions

	// Debug
	fmt.Printf("format: %q, date: %#v", format, date)

	// Format the date conversion
	if date, err := backend(format, date); err != nil {
		errorf("%s: %s", self, err)
	} else {
		fmt.Println(date)
	}
}
