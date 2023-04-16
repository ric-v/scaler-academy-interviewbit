package unsetithbit

// UnsetIthBit returns A with the ith bit unset.
//
// Example: A = 4, B = 1 => 4 & (4 ^ (1 << 1)) => 4 & (4 ^ 2) => 4 & 2 => 0
//
// Time complexity: O(1)
// Space complexity: O(1)
func UnsetIthBit(A, B int) int {
	return A & (A ^ (1 << B))
}
