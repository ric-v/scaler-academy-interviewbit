package kthsymbol

import "math"

// KthSymbol returns the Kth symbol in the grammar
//
// Time complexity: O(logN)
// Space complexity: O(1)
func KthSymbol(A, B int) int {
	if A == 0 {
		return 0
	}

	// Mid point of the string is 2^(A-1) / 2
	mid := int(math.Pow(2, float64(A))) / 2

	// If B is less than mid, then it is in the first half of the string
	if B < mid {

		// Recursively call the function with A-1 and B
		return KthSymbol(A-1, B)
	} else {

		// Recursively call the function with A-1 and (B-mid)
		return 1 - KthSymbol(A-1, (B-mid))
	}
}
