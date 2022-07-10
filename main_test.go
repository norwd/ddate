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
			want: [3]int{},
		},
		//{
		//	name: "All Zeros",
		//	have: [3]string{"0", "0", "0"},
		//	want: [3]int{0, 0, 0},
		//},
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
			if want, e := strconv.Atoi(day); e != nil {
				if want := e.Error(); err == nil {
					t.Fatalf("have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("have %q, want %q", have, want)
				}
			} else if have := date.Day(); have != want {
				t.Fatalf("have %d, want %d", have, want)
			}

			if want, e := strconv.Atoi(month); e != nil {
				if want := e.Error(); err == nil {
					t.Fatalf("have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("have %q, want %q", have, want)
				}
			} else if have := int(date.Month()); have != want {
				t.Fatalf("have %d, want %d", have, want)
			}

			if want, e := strconv.Atoi(year); e != nil {
				if want := e.Error(); err == nil {
					t.Fatalf("have nil, want %q", want)
				} else if have := err.Error(); have != want {
					t.Fatalf("have %q, want %q", have, want)
				}
			} else if have := date.Year(); have != want {
				t.Fatalf("have %d, want %d", have, want)
			}
		})
	}
}
