package equilibriumindexofanarray

import (
	prefixsum "in-place-prefix-sum"
)

// EquilibriumIndex returns the index of the equilibrium point in the array
// An equilibrium point is an index such that the sum of elements at lower
// indexes is equal to the sum of elements at higher indexes.
// For example, in an array A:
// A[0] + A[1] + ... + A[i-1] = A[i+1] + A[i+2] + ... + A[n-1]
// Equilibrium index is at i
// If there is no equilibrium point, then return -1
// If there are more than one equilibrium points, then return the minimum index
// of the equilibrium point
//
// Example:
// A = {-7, 1, 5, 2, -4, 3, 0}
// Output: 3
// Explanation: At index 3, A[0] + A[1] + A[2] = A[4] + A[5] + A[6]
//
// Time complexity: O(n)
// Space complexity: O(1)
func EquilibriumIndex(A []int) int {

	// get the prefix sum of the array
	prefixsum.InPlacePrefixSum(A)

	// iterate over the array and check if the sum of the left side is equal to
	for i := 0; i < len(A); i++ {

		// get the sum of the left side
		// if i is 0, then the sum of the left side is 0
		// otherwise, the sum of the left side is the prefix sum of the previous element
		var left, right int
		if i > 0 {
			left = A[i-1]
		}

		// if i == len(A), then the sum of the right side is 0
		// otherwise, the sum of the right side is the prefix sum of the last element minus the prefix sum of the current element
		if i < len(A)-1 {
			right = A[len(A)-1] - A[i]
		}

		// if the sum of the left side is equal to the sum of the right side, then return the index
		if left == right {
			return i
		}
	}
	return -1
}
