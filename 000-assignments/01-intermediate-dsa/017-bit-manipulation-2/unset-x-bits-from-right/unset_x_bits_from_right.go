package unsetxbitsfromright

// UnsetXBitsFromRight unsets the last B bits of A
// A: 25, B: 3
// A & ~(1 << B - 1) = 25 & ~(1 << 3 - 1) = 25 & ~(8 - 1) = 25 & ~7 = 25 & 24 = 24
//
// Time complexity: O(1)
// Space complexity: O(1)
func UnsetXBitsFromRight(A int64, B int) int64 {
	// 1 << B is a number with only Bth bit set to 1
	// 1 << B - 1 is a number with B bits set to 1
	// ~(1 << B - 1) is a number with B bits set to 0
	// A & ~(1 << B - 1) is a number with B bits set to 0 from right
	return A & ^(1<<B - 1)
}
