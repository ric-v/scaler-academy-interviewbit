package maxsumcontguoussubarray

// MaxSumContiguousSubarray returns the maximum sum of a contiguous subarray
// within a one-dimensional array of numbers which has the largest sum.
// 
// Time complexity: O(n)
// Space complexity: O(1)
func MaxSumContiguousSubarray(A []int) int {

	// Kadane's algorithm (https://en.wikipedia.org/wiki/Maximum_subarray_problem#Kadane's_algorithm)
	// sum is the sum of the current subarray and maxSum is the maximum sum of the subarrays
	var sum, maxSum int = 0, A[0]

	// Iterate over the array and update sum and maxSum accordingly
	for _, a := range A {

		// If sum is greater than maxSum, update maxSum
		sum += a
		if sum > maxSum {
			maxSum = sum
		}

		// If sum is less than zero, reset sum to zero
		// This is because the sum of the subarray will be greater if we start from the next element
		// rather than continuing with the current subarray
		// For example, if the array is [-2, 1, -3, 4, -1, 2, 1, -5, 4]
		// then the sum of the subarray [-2, 1, -3, 4] is 0 and the sum of the subarray [4, -1, 2, 1] is 6
		// So, we reset sum to zero when sum is less than zero
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
