package subarraywithgivensum

// SubArrayWithGivenSum returns the subarray with given sum
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func SubArrayWithGivenSum(A []int, B int) []int {
	var left, right, sum int
	sum = A[0]

	// left and right are the two pointers
	for right < len(A) {

		// if sum is equal to B, return the subarray
		// if sum is greater than B, remove the left element from the sum
		// if sum is less than B, add the right element to the sum
		if sum == B {
			return A[left:right]
		} else if sum > B {
			sum -= A[left]
			left++
		} else {
			right++
			if right < len(A) {
				sum += A[right]
			}
		}
	}
	return []int{-1}
}
