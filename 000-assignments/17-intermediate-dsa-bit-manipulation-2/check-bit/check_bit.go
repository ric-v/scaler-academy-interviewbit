package checkbit

// CheckBit returns 1 if the bit at position B in A is 1, 0 otherwise.
// A is a non-negative integer.
//
// Example: CheckBit(5, 2) = 1 because 5 in binary is 101 and the bit at position 2 is 1.
// Example: CheckBit(4, 1) = 0 because 4 in binary is 100 and the bit at position 1 is 0.
//
// Time complexity: O(1)
// Space complexity: O(1)
func CheckBit(A, B int) int {
	if (A & (1 << B)) == 0 {
		return 0
	} else {
		return 1
	}
}
