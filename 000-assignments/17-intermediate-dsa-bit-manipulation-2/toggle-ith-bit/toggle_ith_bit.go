package toggleithbit

// ToggleIthBit toggles the ith bit of A
// A: 4, B: 1
// A ^ (1 << B) = 4 ^ (1 << 1) = 4 ^ 2 = 6
//
// Time complexity: O(1)
// Space complexity: O(1)
func ToggleIthBit(A int, B int) int {
	// 1 << B is a number with only Bth bit set to 1
	return A ^ (1 << B)
}
