package interestingarray

// InterestingArray returns "Yes" if the XOR of all the elements in the array is even.
// Otherwise, it returns "No".
//
// Time complexity: O(n)
// Space complexity: O(1)
func InterestingArray(A []int) string {
	// XOR all the elements in the array.
	// If the result is even, then the array is interesting.
	var merge int
	for _, a := range A {
		merge ^= a
	}
	if merge&1 == 0 {
		return "Yes"
	}
	return "No"
}
