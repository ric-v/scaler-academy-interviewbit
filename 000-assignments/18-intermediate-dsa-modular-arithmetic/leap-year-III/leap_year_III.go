package leapyeariii

// LeapYearIII returns 1 if A is a leap year, else 0.
//
// Example: 2016 is a leap year. 2017 is not a leap year.
// Since 2016 is divisible by 4, it is a leap year.
//
// Time complexity: O(1)
// Space complexity: O(1)
func LeapYearIII(A int) int {

	// A leap year is divisible by 4. Take the last two digits of the year and divide it by 4.
	if (A%400 == 0) || ((A%4 == 0) && (A%100 != 0)) {
		return 1
	}
	return 0
}
