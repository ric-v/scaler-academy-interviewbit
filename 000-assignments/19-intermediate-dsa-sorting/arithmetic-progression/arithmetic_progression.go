package arithmeticprogression

import "sort"

// ArithmeticProgression checks if the array of integers A is an arithmetic progression
// An arithmetic progression is a sequence of numbers such that the difference between the consecutive terms is constant
// For example, given A = [3, 5, 1], the function returns 1
// Given A = [2, 4, 1], the function returns 0
//
// Time complexity: O(n log n)
// Space complexity: O(1)
func ArithmeticProgression(A []int) int {
	sort.Ints(A)

	// check if the difference between the consecutive terms is constant
	for i := 1; i < len(A); i++ {
		if A[i]-A[i-1] != A[1]-A[0] {
			return 0
		}
	}
	return 1
}
