// ddate is a command line tool for converting dates from the Gregorian Calendar
// to the Discordian Calendar.
//
//     go install github.com/norwd/ddate@latest
//
// Name
//
// ddate - converts Gregorian dates to Discordian dates.
//
// Synopsis
//
//     ddate [+format] [date]
//
// Description
//
// ddate prints the date Discordian date format.
//
// If called with no arguments, ddate will get the current system date, convert
// this to the Discordian date format and print this on the standard output.
// Alternatively, a Gregorian date may be specified on the command line, in the
// form of a numerical day, month, year.
//
// If a format string is specified, the Discordian date will be permitted in a
// format specified by the string. This mechanism works similarly to the format
// string mechanism of date(1), only almost completely different. The fields
// are:
//
//     %A
//
//     Full name of the day of the week (i.e. Sweetmorn)
//
//     %a
//
//     Abbreviated name of the day of the week (i.e. SM)
//
//     %B
//
//     Full name of the season (i.e. Chaos)
//
//     %b
//
//     Abbreviated name of the day of the week (i.e. Chs)
//
//     %d
//
//     Ordinal number of the day in the season (i.e. 23)
//
//     %e
//
//     Cardinal number of the day in the season (i.e. 23rd)
//
//     %Y
//
//     Ordinal year of our lady of of discord (i.e. 3161)
//
//     %y
//
//     Cardinal year of our lady of discord (i.e. 3161st)
//
//     %H
//
//     Name of the current Holyday, if any
//
//     %N
//
//     Magic code to prevent the remainder of the format being printed unless
//     the date is a Holyday.
//
//     %n
//
//     Newline character.
//
//     %t
//
//     Tab character.
//
//     %X
//
//     Number of days remaining until X-Day from the given date.
//
//     %{ and %}
//
//     Used to enclose the part of the string which is to be replaced with the
//     words "St. Tib's Day" if the current day is St. Tib's Day.
//
//     %.
//
//     Try it and see...
//
// Examples
//
//     $ ddate
//     > Sweetmorn, Bureaucracy 42, 3161 YOLD
//
//     $ ddate +"Today is %{%A, the %e of %B%}, %Y. %N%nCelebrate %H!"
//     > Today is Sweetmorn, the 42nd of Bureaucracy, 3161.
//
//     $ ddate +"Today is %{%A, the %e of %B%}, %Y. %N%nCelebrate %H!" 26 9 1995
//     > Today is Prickle-Prickle, the 50th of Bureaucracy, 3161.
//     > Celebrate Bureflux
//
//     $ ddate +"Today is %{%A, the %e of %B%}, %Y. %N%nCelebrate %H!" 29 2 1996
//     > Today is St. Tib's Day, 3162.
//
// Bugs
//
// ddate will produce undefined behaviour if asked to produce the date for St.
// Tib's Day with out the %{ and %} delimiters.
//
// Author
//
// The original ddate was written in C by Jeremy Johnson and significantly
// rewritten by Andrew Bulhak. This version is written in Go, completely from
// scratch, maintains backwards compatibility with the original ddate that was
// distributed as part of util-linux-ng.
//
// Distribution
//
// Public Domain. All Rites Reversed.
//
package main // import "github.com/norwd/ddate"

func main() {
}
