package sumofevenindices

// SumOfEvenIndices returns the sum of even indices
// 
// Time complexity: O(n)
// Space complexity: O(n)
func SumOfEvenIndices(A []int, B [][]int) []int {
	// create prefix sum array
	var prefixSum = make([]int, len(A))
	prefixSum[0] = A[0]

	// Calculate the prefix sum
	for i := 1; i < len(A); i++ {

		// if the length of the subarray is even
		if i%2 == 0 {
			// add the last element to the prefix sum
			prefixSum[i] = A[i] + prefixSum[i-1]
		} else {
			// if the length of the subarray is odd, the last element is not added
			prefixSum[i] = prefixSum[i-1]
		}
	}

	// create result array
	var result = make([]int, len(B))

	// iterate over the B slice
	for i, b := range B {
		// if the first element of the slice is zero:
		if b[0] == 0 {
			// assign the element at index i to the element at index b[1]
			result[i] = prefixSum[b[1]]
		} else {
			// otherwise, assign the element at index i to the difference between the element at index b[1] and the element at index b[0] - 1
			result[i] = prefixSum[b[1]] - prefixSum[b[0]-1]
		}
	}
	return result
}
