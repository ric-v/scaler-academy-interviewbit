package powerwithmodules

// PowerWithModules returns A^B % C (A to the power of B modulo C)
// A, B, C are integers
//
// Example: A = 2, B = 3, C = 3 => 2^3 % 3 = 8 % 3 = 2
//
// Time complexity: O(B)
// Space complexity: O(1)
func PowerWithModules(A, B, C int) int {
	var result int = 1

	// do A * A * A * ... * A (B times)
	for i := 0; i < B; i++ {
		result = (result * A) % C
	}

	// return result % C
	return result
}
