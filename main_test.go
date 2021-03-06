package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/norwd/ddate/internal/os"
)

func TestErrorf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string // name of the test case
		want string // expected error message
		have string // input message
		args []any  // format arguments
	}{
		{
			name: "Empty Error Message",
		},
		{
			name: "Unformatted Error Message",
			want: "Expected error message",
			have: "Expected error message",
		},
		{
			name: "Formatted Error Message",
			want: "Expected error message",
			have: "Expected %s message",
			args: []any{"error"},
		},
	}

	for _, test := range tests {
		// shadow loop var to prevent nasty bugs
		test := test

		// trim whitespace from name of test case
		name := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}

			return r
		}, test.name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var buf bytes.Buffer
			var exit int

			defer os.MockAndLockStderr(&buf).Unlock()
			defer os.MockAndLockExit(func(code int) { exit = code }).Unlock()

			// Act
			errorf(test.have, test.args...)

			// Assert
			if have, want := buf.String(), fmt.Sprintln(test.want); have != want {
				t.Errorf("error message: have %q, want %q", have, want)
			}

			if have, want := exit, 1; have != want {
				t.Errorf("exit code: have %d, want %d", have, want)
			}
		})
	}
}

func TestPrintln(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string // name of the test case
		want string // expected message
		have string // input message
	}{
		{
			name: "Empty Message",
		},
		{
			name: "Simple Message",
			want: "This is a simple message",
			have: "This is a simple message",
		},
	}

	for _, test := range tests {
		// shadow loop var to prevent nasty bugs
		test := test

		// trim whitespace from name of test case
		name := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}

			return r
		}, test.name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var buf bytes.Buffer

			defer os.MockAndLockStdout(&buf).Unlock()

			// Act
			println(test.have)

			// Assert
			if have, want := buf.String(), fmt.Sprintln(test.want); have != want {
				t.Errorf("error message: have %q, want %q", have, want)
			}
		})
	}
}

func TestParseDDMMYYYY(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string    // name of the test case
		have [3]string // input in DD, MM, and YYYY
		want [3]int    // expected day, month, and year
	}{
		{
			name: "All Empty",
			have: [3]string{"", "", ""},
		},
		{
			name: "All Zeros",
			have: [3]string{"0", "0", "0"},
			want: [3]int{30, 11, -1},
		},
		{
			name: "Valid Date",
			have: [3]string{"6", "8", "1999"},
			want: [3]int{6, 8, 1999},
		},
		{
			name: "Valid Date Leading Zeros",
			have: [3]string{"06", "08", "01999"},
			want: [3]int{6, 8, 1999},
		},
		{
			name: "Invalid Day",
			have: [3]string{"_6", "8", "1999"},
			want: [3]int{0, 0, 0},
		},
		{
			name: "Invalid Month",
			have: [3]string{"6", "_8", "1999"},
			want: [3]int{0, 0, 0},
		},
		{
			name: "Invalid Year",
			have: [3]string{"6", "8", "_1999"},
			want: [3]int{0, 0, 0},
		},
		{
			name: "Invalid Date",
			have: [3]string{"32", "10", "2022"},
			want: [3]int{1, 11, 2022},
		},
	}

	for _, test := range tests {
		// shadow loop var to prevent nasty bugs
		test := test

		// trim whitespace from name of test case
		name := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}

			return r
		}, test.name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			day, month, year := test.have[0], test.have[1], test.have[2]

			// Act
			date, err := parseDDMMYYYY(day, month, year)

			// Assert
			if _, want := strconv.Atoi(day); want != nil {
				if want := want.Error(); err == nil {
					t.Fatalf("error: have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("error: have %q, want %q", have, want)
				}

				return // don't keep testing, expected failure detected
			} else if have, want := date.Day(), test.want[0]; have != want && err == nil {
				t.Fatalf("day: have %d, want %d", have, want)
			}

			if _, want := strconv.Atoi(month); want != nil {
				if want := want.Error(); err == nil {
					t.Fatalf("error: have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("error: have %q, want %q", have, want)
				}

				return // don't keep testing, expected failure detected
			} else if have, want := int(date.Month()), test.want[1]; have != want && err == nil {
				t.Fatalf("want: have %d, want %d", have, want)
			}

			if _, want := strconv.Atoi(year); want != nil {
				if want := want.Error(); err == nil {
					t.Fatalf("error: have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("error: have %q, want %q", have, want)
				}

				return // don't keep testing, expected failure detected
			} else if have, want := date.Year(), test.want[2]; have != want && err == nil {
				t.Fatalf("year: have %d, want %d", have, want)
			}

			if err != nil {
				t.Fatalf("error: have %q, want nil", err)
			}
		})
	}
}

func TestMain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string    // name of the test case
		self        string    // name of the application
		args        []string  // arguments to pass to main
		date        string    // date to return from the backend (if empty expect err)
		want        string    // expected output
		ptrn        string    // expected format pattern
		time        time.Time // expected time to pass to the backend
		exit        int       // expected error code (signals where output is expected)
		callBackend bool      // should the backend expect to be called?
	}{
		{
			name:        "No Args",
			self:        "ddate",
			args:        []string{},
			date:        "Today's discordian date",
			want:        "Today's discordian date",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        0,
			callBackend: true,
		},
		{
			name:        "Only Format",
			self:        "ddate",
			args:        []string{"+Some fancy format string"},
			date:        "Today's discordian date",
			want:        "Today's discordian date",
			ptrn:        "Some fancy format string",
			time:        time.Now(),
			exit:        0,
			callBackend: true,
		},
		{
			name:        "Only DD MM YYYY",
			self:        "ddate",
			args:        []string{"10", "11", "1999"},
			date:        "The discordian date for 1999-11-10",
			want:        "The discordian date for 1999-11-10",
			ptrn:        defaultFormat,
			time:        time.Date(1999, 11, 10, 0, 0, 0, 0, time.Local),
			exit:        0,
			callBackend: true,
		},
		{
			name:        "Format And DD MM YYYY",
			self:        "ddate",
			args:        []string{"+Some fancy format string", "10", "11", "1999"},
			date:        "The discordian date for 1999-11-10",
			want:        "The discordian date for 1999-11-10",
			ptrn:        "Some fancy format string",
			time:        time.Date(1999, 11, 10, 0, 0, 0, 0, time.Local),
			exit:        0,
			callBackend: true,
		},
		{
			name:        "Only Argument Is Invalid Format",
			self:        "ddate",
			args:        []string{"Some non-format string"},
			date:        "",
			want:        "ddate: not enough arguments for DD MM YYYY",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And DD",
			self:        "ddate",
			args:        []string{"+Some format string", "10"},
			date:        "",
			want:        "ddate: not enough arguments for DD MM YYYY",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And DD MM",
			self:        "ddate",
			args:        []string{"+Some format string", "10", "11"},
			date:        "",
			want:        "ddate: not enough arguments for DD MM YYYY",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And DD MM YY space YY",
			self:        "ddate",
			args:        []string{"+Some format string", "10", "11", "19", "99"},
			date:        "",
			want:        "ddate: too many arguments for DD MM YYYY",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And Invalid DD",
			self:        "ddate",
			args:        []string{"+Some format string", "1_0", "11", "1999"},
			date:        "",
			want:        "ddate: strconv.Atoi: parsing \"1_0\": invalid syntax",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And Invalid MM",
			self:        "ddate",
			args:        []string{"+Some format string", "10", "1_1", "1999"},
			date:        "",
			want:        "ddate: strconv.Atoi: parsing \"1_1\": invalid syntax",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And Invalid YYYY",
			self:        "ddate",
			args:        []string{"+Some format string", "10", "11", "19_99"},
			date:        "",
			want:        "ddate: strconv.Atoi: parsing \"19_99\": invalid syntax",
			ptrn:        defaultFormat,
			time:        time.Now(),
			exit:        1,
			callBackend: false,
		},
		{
			name:        "Format And DD MM YYYY Backend Failure",
			self:        "ddate",
			args:        []string{"+Some fancy format string", "10", "11", "1999"},
			date:        "",
			want:        "ddate: expected backend error",
			ptrn:        "Some fancy format string",
			time:        time.Date(1999, 11, 10, 0, 0, 0, 0, time.Local),
			exit:        1,
			callBackend: true,
		},
	}

	for _, test := range tests {
		// shadow loop var to prevent nasty bugs
		test := test

		// trim whitespace from name of test case
		name := strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}

			return r
		}, test.name)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var errBuf, outBuf bytes.Buffer // fake streams
			var backendCalls, exitCalls int // call counters
			var exit int                    // record exit code

			// mock io
			defer os.MockAndLockStderr(&errBuf).Unlock()
			defer os.MockAndLockStdout(&outBuf).Unlock()

			// mock argv
			defer os.MockAndLockArgs(test.self, test.args).Unlock()

			// mock backend
			defer mockAndLockBackend(func(format string, date time.Time) (string, error) {
				backendCalls++

				// check that the format is as expected
				if have, want := format, test.ptrn; have != want {
					t.Errorf("wrong format: have %q, want %q", have, want)
				}

				// check that the date is within an hour of the expected date
				if have, want := date, test.time; math.Abs(have.Sub(want).Hours()) > 1 {
					t.Errorf("wrong date: have %s, want %s", have, want)
				}

				// if the expected date is empty, then an error is expected
				if test.date == "" {
					return "", errors.New("expected backend error")
				}

				return test.date, nil
			}).Unlock()

			// mock exit
			defer os.MockAndLockExit(func(code int) {
				exitCalls++

				exit = code

				// simulate exit
				panic("exit")
			}).Unlock()

			// Act
			func() {
				// catch "fake" panic
				defer func() {
					if err := recover(); err != nil {
						if err != "exit" {
							panic(err)
						}
					}
				}()

				main()
			}()

			// Assert
			if test.callBackend {
				if backendCalls != 1 {
					t.Errorf("backend called %d times, want once", backendCalls)
				}
			} else {
				if backendCalls != 0 {
					t.Errorf("backend called was called %d times, want 0", backendCalls)
				}
			}

			if test.exit > 0 && exitCalls != 1 {
				t.Errorf("exit called %d times, want once", exitCalls)
			} else if test.exit == 0 && exitCalls != 0 {
				t.Errorf("exit called was called %d times, want 0", exitCalls)
			}

			if test.exit == 0 {
				// test expects success
				if have, want := outBuf.String(), fmt.Sprintln(test.want); have != want {
					t.Errorf("stdout: have %q, want %q", have, want)
				}

				if have, want := errBuf.String(), ""; have != want {
					t.Errorf("stderr: have %q, want %q", have, want)
				}
			} else {
				// test expects failure
				if have, want := outBuf.String(), ""; have != want {
					t.Errorf("stdout: have %q, want %q", have, want)
				}

				if have, want := errBuf.String(), fmt.Sprintln(test.want); have != want {
					t.Errorf("stderr: have %q, want %q", have, want)
				}
			}

			if have, want := exit, test.exit; have != want {
				t.Errorf("exit code: have %d, want %d", have, want)
			}
		})
	}
}
