package tendigitsorting

import "sort"

// TenDigitSorting sorts the array of integers A by the tens digit
// For example, given A = [15, 11, 7, 19], the function returns [7, 19, 15, 11]
// Given A = [2, 24, 22, 19], the function returns [2, 19, 24, 22]
//
// Time complexity: O(n log n)
// Space complexity: O(1)
func TenDigitSorting(A []int) []int {

	// sort the array by the tens digit
	sort.Slice(A, func(i, j int) bool {

		// get the tens digit of the number at index i and j
		// if A[i] = 15, then tensPlaceI = 15 % 100 / 10 = 1
		// if A[j] = 32, then tensPlaceJ = 32 % 100 / 10 = 3
		tensPlaceI := (A[i] % 100) / 10
		tensPlaceJ := (A[j] % 100) / 10
		if tensPlaceI == tensPlaceJ {
			return A[i] > A[j]
		}
		return tensPlaceI < tensPlaceJ
	})
	return A
}
