package countingsubarrays

// CountingSubarrays returns the number of subarrays with sum less than B
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func CountingSubarrays(A []int, B int) int {

	var count int
	for i := 0; i < len(A); i++ {
		var sum int
		// sum all elements from i to j and check if sum < B
		for j := i; j < len(A); j++ {
			sum += A[j]

			// if sum < B, increment count
			if sum < B {
				count++
			}
		}
	}
	return count
}
