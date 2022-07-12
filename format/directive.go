package format

// Directive specifies a formatting function within a format string.
type Directive string

const (
	FullWeekdayDirective  Directive = "%A" // Formats the full name of the day of the week (i.e. Sweetmorn).
	AbbrWeekdayDirective  Directive = "%a" // Formats the abbreviated name of the day of the week (i.e. SM).
	FullSeasonDirective   Directive = "%B" // Formats the full name of the season (i.e. Chaos).
	AbbrSeasonDirective   Directive = "%b" // Formats the abbreviated name of the day of the week (i.e. Chs).
	OrdinalDayDirective   Directive = "%d" // Formats the ordinal number of the day in the season (i.e. 23).
	CardinalDayDirective  Directive = "%e" // Formats the cardinal number of the day in the season (i.e. 23rd).
	OrdinalYearDirective  Directive = "%Y" // Formats the ordinal year of our lady of of discord (i.e. 3161).
	CardinalYearDirective Directive = "%y" // Formats the cardinal year of our lady of discord (i.e. 3161st).
	HolydayDirective      Directive = "%H" // Formats the name of the current Holyday, if any (i.e. Confuflux).
	NonHolidayDirective   Directive = "%N" // Is a magic code to prevent the remainder of the format string being printed unless the date is a Holyday.
	NewlineDirective      Directive = "%n" // Formats a newline character.
	TabDirective          Directive = "%t" // Formats a tab character.
	PercentDirective      Directive = "%%" // Formats a literal percent sign character.
	XDayDirective         Directive = "%X" // Formats the cardinal number of days remaining until X-Day from the given date.
	StartTibsDayDirective Directive = "%{" // Begins a special block to enclose part of the string.
	EndTibsDayDirective   Directive = "%}" // Closes a special block opened by %}, the enclosed the part of the string is replaced with the words "St. Tib's Day" if the current day is St. Tib's Day.
	MagicDirective        Directive = "%." // Try it and see...
)
