package reversebits

// ReverseBits reverses the bits of a 32-bit unsigned integer.
// For example, 3 => 3221225472
//
// Time complexity: O(1)
// Space complexity: O(1)
func ReverseBits(num uint32) uint32 {
	var result uint32

	// 32 bits in a uint32 (4 bytes)
	for i := 0; i < 32; i++ {

		// Shift result to the left by 1 bit
		result <<= 1
		// Add the last bit of num to result
		result |= num & 1
		// Shift num to the right by 1 bit
		num >>= 1
	}
	return result
}
