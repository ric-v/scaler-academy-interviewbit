package findnthmagicnumber

// FindNthMagicNumber returns the nth magic number
// A magic number is defined as a number which can be expressed as a power of 5 or sum of unique powers of 5.
// First few magic numbers are 5, 25, 30(5 + 25), 125, 130(125 + 5), ….
// Example: 3rd magic number is 30 (5 + 25) and 10th magic number is 650 (5 + 5 + 5 + 5 + 5 + 5 + 5 + 5 + 5 + 25)
//
// Time complexity: O(log5(A))
// Space complexity: O(1)
func FindNthMagicNumber(A int) int {
	ans := 0
	x := 1

	// iterate over all bits of A from right to left
	for A > 0 {

		// if the current bit is set, add x to the answer
		// (x is the current power of 5)
		x *= 5
		if (A & 1) == 1 {
			ans += x
		}

		// right shift A by 1, so that the next bit becomes the current bit
		A = A >> 1
	}
	return ans
}
