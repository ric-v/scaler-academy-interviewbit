package subarrayingivenrange

// SubArrayInGivenRange returns a subarray of A from B to C
// B and C are inclusive
//
// Example: [4, 3, 2, 6], 1, 3 => [3, 2, 6]
//
// Time complexity: O(n)
// Space complexity: O(n)
func SubArrayInGivenRange(A []int, B int, C int) []int {
	var result []int

	// iterate from B to C and append to result
	for i := B; i <= C; i++ {
		result = append(result, A[i])
	}
	return result
}
