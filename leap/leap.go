// Package leap provides a way to see if a given year is a leap year
package leap

const testVersion = 2

// IsLeapYear calculates if a given year is a leap year or not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
