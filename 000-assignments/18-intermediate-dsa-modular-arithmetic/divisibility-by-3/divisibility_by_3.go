package divisibilityby3

// DivisibilityBy3 returns 1 if the sum of the digits of the number is divisible by 3
//
// A = [1, 2, 3] => 1 (123 % 3 == 0)
// A = [1, 0, 0, 1, 2] => 0 (10201 % 3 == 1)
//
// Time complexity: O(N)
// Space complexity: O(1)
func DivisibilityBy3(A []int) int {
	var result int
	// sum all digits
	for i := 0; i < len(A); i++ {
		result += (A[i])
	}

	// check if the sum is divisible by 3
	if result%3 == 0 {
		return 1
	}
	return 0
}
