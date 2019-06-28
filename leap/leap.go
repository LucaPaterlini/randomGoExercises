// Package leap contains functions about the leap yaer
package leap

// IsLeapYear check if a give year is a leap year
func IsLeapYear(x int) bool {
	return x%400==0 || (x%100!=0 && x%4==0)
}
