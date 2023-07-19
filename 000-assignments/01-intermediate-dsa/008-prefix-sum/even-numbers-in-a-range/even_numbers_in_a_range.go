package evennumbersinarange

import prefixsum "in-place-prefix-sum"

// EvenNumbers returns the number of even numbers in a range.
//
// Time complexity: O(n)
// Space complexity: O(1)
func EvenNumbers(A []int, B [][]int) []int {

	// calculate the prefix sum of the array
	// the prefix sum of an array is the sum of count of all even numbers in the array
	// up to and including the current element.
	// eg: A = [1, 2, 3, 4, 5] => A = [0, 1, 1, 2, 2]
	prefixsum.FilterPrefixSum(A, func(x int) int {
		if x%2 == 0 {
			return 1
		}
		return 0
	})

	// calculate the number of even numbers in each range
	var result []int
	for _, query := range B {
		var sum int

		// if the range starts at the beginning of the array then the sum is the value at the end of the range
		// otherwise the sum is the value at the end of the range minus the value at the start of the range
		if query[0] == 0 {
			sum = A[query[1]]
		} else {
			sum = A[query[1]] - A[query[0]-1]
		}

		// append the sum to the result
		result = append(result, sum)
	}
	return result
}
