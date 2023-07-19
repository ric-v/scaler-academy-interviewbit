package goodsubarrays

// GoodSubarrays returns the number of good subarrays in A
// A subarray is good if the sum of the elements in the subarray is greater than B
// if the subarray is odd length and less than B if the subarray is even length
//
// Example: A = [1, 2, 3, 4, 5], B = 4
// even length subarrays: [1, 2], [2, 3], [3, 4], [4, 5]
// odd length subarrays: [1, 2, 3], [2, 3, 4], [3, 4, 5]
// Of these, only [1, 2, 3], [1, 2, 3, 4, 5], [2, 3, 4], [3, 4, 5] are good
// hence the function returns 4
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func GoodSubarrays(A []int, B int) int {

	var goodArrCount int
	for i := 0; i < len(A); i++ {

		// sum of elements in the subarray
		var sum int
		for j := i; j < len(A); j++ {
			sum += A[j]

			// if the subarray is even length and the sum is greater than B
			// or if the subarray is odd length and the sum is less than B
			if (isOdd(i, j) && sum > B) || (!isOdd(i, j) && sum < B) {
				goodArrCount++
			}
		}
	}
	return goodArrCount
}

// isOdd returns true if the sum of i and j is odd
// else it returns false
//
// for an array on n^2 iteration,
// (i+j)%2 will be 1 for even length subarrays and 0 for odd length subarrays
// so we invert the result
func isOdd(i, j int) bool {
	return !((i+j)%2 == 1)
}
