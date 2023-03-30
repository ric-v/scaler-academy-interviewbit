package maximumsubarray

// MaximumSubArray returns the maximum sum of a subarray of C that is less than or equal to B.
// A is the length of C.
// B is the maximum sum of a subarray allowed.
// C is the array of integers.
//
// Example: A = 5, B = 12, C = [2, 1, 3, 4, 5] => 12
//
// Time complexity: O(n^2)
// Space complexity: O(n)
func MaximumSubArray(A int, B int, C []int) int {

	// Create a slice of sums
	sums := make([]int, A+1)
	for i := 1; i <= A; i++ {
		sums[i] = sums[i-1] + C[i-1]
	}

	// Iterate through the sums slice and find the maximum sum of a subarray that is less than or equal to B
	max := 0
	for i := 0; i <= A; i++ {
		for j := i + 1; j <= A; j++ {
			sum := sums[j] - sums[i]
			if sum <= B && sum > max {
				max = sum
			}
		}
	}

	return max
}
