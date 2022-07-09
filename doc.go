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
// Usage:
//
//     ddate [+format] [date]
//
// Options:
//
//     - %A formats the full name of the day of the week (i.e. Sweetmorn).
//     - %a formats the abbreviated name of the day of the week (i.e. SM).
//     - %B formats the full name of the season (i.e. Chaos).
//     - %b formats the abbreviated name of the day of the week (i.e. Chs).
//     - %d formats the ordinal number of the day in the season (i.e. 23).
//     - %e formats the cardinal number of the day in the season (i.e. 23rd).
//     - %Y formats the ordinal year of our lady of of discord (i.e. 3161).
//     - %y formats the cardinal year of our lady of discord (i.e. 3161st).
//     - %H formats the name of the current Holyday, if any (i.e. Confuflux).
//     - %N is a magic code to prevent the remainder of the format string being
//       printed unless the date is a Holyday.
//     - %n formats a newline character.
//     - %t formats a tab character.
//     - %% formats a literal percent sign character.
//     - %X formats the cardinal number of days remaining until X-Day from the
//       given date.
//     - %{ and %} are used to enclose the part of the string which is to be
//       replaced with the words "St. Tib's Day" if the current day is St. Tib's Day.
//     - %. Try it and see...
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
// string mechanism of date(1), only almost completely different.
//
// Examples
//
// Without any arguments, ddate prints today's Discordian Date according to the
// default format.
//
//     $ ddate
//     > Sweetmorn, Bureaucracy 42, 3161 YOLD
//
// A custom format can be specified with a plus sign and the percent sign escape
// codes listed above.
//
//     $ ddate +"Today is %{%A, the %e of %B%}, %Y. %N%nCelebrate %H!"
//     > Today is Sweetmorn, the 42nd of Bureaucracy, 3161.
//
// A custom date can specified (with or without a custom format) as DD MM YYYY.
//
//     $ ddate +"Today is %{%A, the %e of %B%}, %Y. %N%nCelebrate %H!" 26 9 1995
//     > Today is Prickle-Prickle, the 50th of Bureaucracy, 3161.
//     > Celebrate Bureflux
//
// If the date is February 29th, the Special St. Tib's Day formatters are used
// to display "St. Tib's Day".
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
package main
