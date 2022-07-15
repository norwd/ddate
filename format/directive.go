package format

// Directive specifies a formatting function within a format string.
type Directive string

const (
	// Formats the full name of the day of the week (i.e. Sweetmorn).
	FullWeekdayDirective Directive = "%A"

	// Formats the abbreviated name of the day of the week (i.e. SM).
	AbbrWeekdayDirective Directive = "%a"

	// Formats the full name of the season (i.e. Chaos).
	FullSeasonDirective Directive = "%B"

	// Formats the abbreviated name of the day of the week (i.e. Chs).
	AbbrSeasonDirective Directive = "%b"

	// Formats the ordinal number of the day in the season (i.e. 23).
	OrdinalDayDirective Directive = "%d"

	// Formats the cardinal number of the day in the season (i.e. 23rd).
	CardinalDayDirective Directive = "%e"

	// Formats the ordinal year of our lady of of discord (i.e. 3161).
	OrdinalYearDirective Directive = "%Y"

	// Formats the cardinal year of our lady of discord (i.e. 3161st).
	CardinalYearDirective Directive = "%y"

	// Formats the name of the current Holyday, if any (i.e. Confuflux).
	HolydayDirective Directive = "%H"

	// Is a magic code to prevent the remainder of the format string being
	// printed unless the date is a Holyday.
	NonHolidayDirective Directive = "%N"

	// Formats a newline character.
	NewlineDirective Directive = "%n"

	// Formats a tab character.
	TabDirective Directive = "%t"

	// Formats a literal percent sign character.
	PercentDirective Directive = "%%"

	// Formats the cardinal number of days remaining until X-Day from the given
	// date.
	XDayDirective Directive = "%X"

	// Begins a special block to enclose part of the string.
	StartTibsDayDirective Directive = "%{"

	// Closes a special block opened by %}, the enclosed the part of the string
	// is replaced with the words "St. Tib's Day" if the current day is St.
	// Tib's Day.
	EndTibsDayDirective Directive = "%}"

	// Try it and see...
	MagicDirective Directive = "%."
)
