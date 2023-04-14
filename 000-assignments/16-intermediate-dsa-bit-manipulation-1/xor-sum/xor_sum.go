package xorsum

// XorSum returns the sum of two integers using the XOR operator.
// The XOR operator is a binary operator that returns 1 if the two
// bits are different and 0 if they are the same.
//
// Time complexity: O(1)
// Space complexity: O(1)
func XorSum(A, B int) int {
	return A ^ B
}
