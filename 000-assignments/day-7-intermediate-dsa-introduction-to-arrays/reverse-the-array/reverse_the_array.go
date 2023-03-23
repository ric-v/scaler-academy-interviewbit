package reversethearray

import reverse "reverse-in-a-range"

// ReverseTheArray reverses the whole array in O(1) space and returns it
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func ReverseTheArray(A []int) []int {
	return reverse.ReverseInARange(A, 0, len(A)-1)
}

/*
// full logic
// since reverse in a range is a super set of this package, reusing the existing code.
func solve(A []int )  ([]int) {

	var start, end int = 0, len(A) - 1
	for start < end {

		// swap
		A[start], A[end] = A[end], A[start]
		start++
		end--
	}
	return A
}

*/
