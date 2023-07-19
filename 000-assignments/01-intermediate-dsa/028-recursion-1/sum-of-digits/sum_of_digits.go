package sumofdigits

// SumOfDigits returns the sum of digits of a number
// For example, if A = 123, the function returns 6
//
// Time complexity: O(n)
// Space complexity: O(1)
func SumOfDigits(A int) int {
	if A == 0 {
		return 0
	}
	digit := A % 10
	A /= 10
	return digit + SumOfDigits(A)
}
