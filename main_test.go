package main

import (
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// backend = func(string, time.Time) (string, error) {}

func TestParseDDMMYYYY(t *testing.T) {
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
			name: "Invalid Numbers",
			have: [3]string{"_6", "_8", "_1999"},
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
			// Arrange
			day, month, year := test.have[0], test.have[1], test.have[2]

			// Act
			date, err := parseDDMMYYYY(day, month, year)

			// Assert
			if _, want := strconv.Atoi(day); want != nil {
				if want := want.Error(); err == nil {
					t.Fatalf("have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("have %q, want %q", have, want)
				}

				return // don't keep testing, expected failure detected
			} else if have, want := date.Day(), test.want[0]; have != want {
				t.Fatalf("have %d, want %d", have, want)
			}

			if _, want := strconv.Atoi(month); want != nil {
				if want := want.Error(); err == nil {
					t.Fatalf("have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("have %q, want %q", have, want)
				}

				return // don't keep testing, expected failure detected
			} else if have, want := int(date.Month()), test.want[1]; have != want {
				t.Fatalf("have %d, want %d", have, want)
			}

			if _, want := strconv.Atoi(year); want != nil {
				if want := want.Error(); err == nil {
					t.Fatalf("have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("have %q, want %q", have, want)
				}

				return // don't keep testing, expected failure detected
			} else if have, want := date.Year(), test.want[2]; have != want {
				t.Fatalf("have %d, want %d", have, want)
			}
		})
	}
}
