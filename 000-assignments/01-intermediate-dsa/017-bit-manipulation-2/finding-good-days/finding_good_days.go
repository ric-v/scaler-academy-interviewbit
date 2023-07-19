package findinggooddays

// FindingGoodDays returns the number of good days in the given month.
//
// A good day is a day where the number of 1s in the binary representation of the day is odd.
//
// Time complexity: O(log(A))
// Space complexity: O(1)
func FindingGoodDays(A int) int {
	count := 0

	// iterate over all bits of A from right to left
	for A > 0 {

		// if the current bit is set, increment the count
		x := A & 1
		if x == 1 {
			count++
		}
		// right shift A by 1, so that the next bit becomes the current bit
		A /= 2
	}
	return count
}
