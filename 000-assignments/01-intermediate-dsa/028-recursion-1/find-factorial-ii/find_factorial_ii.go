package findfactorialii

// FindFactorial returns the factorial of A
//
// Time complexity: O(n)
// Space complexity: O(1)
func FindFactorial(A int) int {
	if A == 1 {
		return 1
	}

	// F(n) = n * F(n-1)
	return A * FindFactorial(A-1)
}
