package helpfromsam

// HelpFromSam returns the minimum number of operations required to convert A to 0
// Operations:
// 1. If A is even, divide A by 2
// 2. If A is odd, subtract 1 from A
//
// A = 3 => 2 (3 -> 2 -> 1 -> 0)
// A = 6 => 2 (6 -> 3 -> 2 -> 1 -> 0)
//
// Time complexity: O(logN)
// Space complexity: O(1)
func HelpFromSam(A int) int {
	ctr := 0

	// while A is not 0
	for A > 0 {

		// if A is odd subtract 1 from A and increment counter by 1
		// else divide A by 2
		if (A & 1) == 1 {
			A -= 1
			ctr++
		} else {
			A /= 2
		}
	}
	return ctr
}
