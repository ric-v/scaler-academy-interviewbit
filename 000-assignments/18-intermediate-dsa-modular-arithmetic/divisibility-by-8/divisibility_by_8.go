package divisibilityby8

import "strconv"

// DivisibilityBy8 checks if the last 3 digits of a number is divisible by 8
//
// Example: 611238745789 is not divisible by 8 because 789 is not divisible by 8
// Example: 611238745888 is divisible by 8 because 888 is divisible by 8
//
// Time complexity: O(1)
// Space complexity: O(1)
func DivisibilityBy8(A string) int {
	// If the number is less than 3 digits, return 0
	if len(A) > 3 {
		A = A[len(A)-3:]
	}

	// Convert the string to an integer
	num, _ := strconv.Atoi(A)

	// Check if the number is divisible by 8
	if num%8 == 0 {
		return 1
	}
	return 0
}
