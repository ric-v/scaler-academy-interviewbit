package numberof1bits

// NumberOf1Bits returns the number of 1 bits in the binary representation of A
// Example: 11 (1011) => 3 (3 1 bits)
//
// Time complexity: O(log n)
// Space complexity: O(1)
func NumberOf1Bits(A int) int {
	var count int
	// A & 1 is 1 if the last bit of A is 1, 0 otherwise
	// A >> 1 shifts A to the right by 1 bit
	for A > 0 {
		count += A & 1
		A = A >> 1
	}
	return count
}
