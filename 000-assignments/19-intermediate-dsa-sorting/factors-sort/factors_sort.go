package factorssort

import (
	count "count_factor"
	"sort"
)

// FactorSort sorts the array of integers A in ascending order of the number of factors
// If two numbers have the same number of factors, then the larger number comes first
// For example, given A = [6, 8, 9], the function returns [9, 6, 8]
// Given A = [2, 4, 7], the function returns [2, 7, 4]
//
// Time complexity: O(n log n)
// Space complexity: O(n)
func FactorSort(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		// sort by number of factors, if the count is same, then sort by value
		if count.CountFactor(A[i]) == count.CountFactor(A[j]) {
			return A[i] < A[j]
		}
		return count.CountFactor(A[i]) < count.CountFactor(A[j])
	})
	return A
}
