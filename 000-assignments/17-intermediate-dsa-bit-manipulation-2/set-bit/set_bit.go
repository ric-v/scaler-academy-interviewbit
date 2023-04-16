package setbit

// SetBit returns a number with a 1 in the Ath and Bth bits
// 
// Example: A = 3, B = 5 => 1 << 3 | 1 << 5 => 8 | 32 => 40
//
// Time complexity: O(1)
// Space complexity: O(1)
func SetBit(A, B int) int {
	// 1 << B is a number with a 1 in the Bth bit
	return (1 << A) | (1 << B)
}
