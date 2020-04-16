package leap

// IsLeapYear returns if the year is a leap year on gregorian calendar
func IsLeapYear(yr int) bool {
	return (yr % 400 == 0) || (yr % 100 != 0 && yr % 4 == 0)
}
