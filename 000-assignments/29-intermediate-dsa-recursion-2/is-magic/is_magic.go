package ismagic

// IsMagic checks if a number is magic or not
// A number is magic if the sum of its digits is 1
// If the sum of the digits is greater than 9, the sum is calculated again
// and the process is repeated until the sum is 1
//
// Example: 83557 -> 8 + 3 + 5 + 5 + 7 = 28 -> 2 + 8 = 10 -> 1 + 0 = 1 -> Magic
// Example: 12345 -> 1 + 2 + 3 + 4 + 5 = 15 -> 1 + 5 = 6 -> Not Magic
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsMagic(A int) int {
	if A < 10 {
		if A == 1 {
			return 1
		}
		return 0
	}
	return IsMagic(sum(A))
}

// sum calculates the sum of the digits of a number recursively
//
// Example: 83557 -> 8 + 3 + 5 + 5 + 7 = 28
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func sum(A int) int {
	if A == 0 {
		return 0
	}
	digit := A % 10
	A /= 10
	return digit + sum(A)
}
