package simplereverse

// SimpleReverse reverses a string. This is a simple implementation of the reverse function.
// 
// Time complexity: O(n), where n is the length of the string.
// Space complexity: O(n).
func SimpleReverse(A string) string {
	var result string
	// Iterate over the string in reverse order and append each character to the result string.
	for i := len(A) - 1; i >= 0; i-- {
		result += string(A[i])
	}
	return result
}
