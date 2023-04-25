package concatenatethreenumbers

import "strconv"

// ConcatenateThreeNumbers takes three integers A, B, C and returns the number
// that is formed by concatenating A, B, and C.
//
// For example, if A = 1, B = 2, and C = 3, then the function should return 123.
//
// Time complexity: O(log(A+B))
// Space complexity: O(1)
func ConcatenateThreeNumbers(A, B, C int) int {

	// Find the minimum and maximum of A, B, and C.
	min := A
	if B < min {
		min = B
	}
	if C < min {
		min = C
	}

	// Find the maximum of A, B, and C.
	max := A
	if B > max {
		max = B
	}
	if C > max {
		max = C
	}

	// Find the middle number.
	mid := A + B + C - (min + max)

	// Concatenate the numbers.
	s := strconv.Itoa(min) + strconv.Itoa(mid) + strconv.Itoa(max)

	// Convert the string to an integer.
	ret, _ := strconv.Atoi(s)
	return ret
}
