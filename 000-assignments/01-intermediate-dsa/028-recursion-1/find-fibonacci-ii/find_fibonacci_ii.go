package findfibonacciii

// FindFibonacci returns the Ath fibonacci number
//
// Time complexity: O(2^n)
// Space complexity: O(1)
func FindFibonacci(A int) int {
	if A == 0 {
		return 0
	} else if A == 1 {
		return 1
	} else {
		// F(n) = F(n-1) + F(n-2)
		return FindFibonacci(A-1) + FindFibonacci(A-2)
	}
}
