package toupper

// ToUpper converts the lowercase letters in A to uppercase.
// A consists of multiple words separated by space. The task is to convert lowercase letters to uppercase.
//
// Time complexity: O(n)
// Space complexity: O(1)
func ToUpper(A []byte) []byte {

	// Iterate over the bytes in A.
	for i, a := range A {
		// If the byte is a lowercase letter, convert it to uppercase.
		if a >= 'a' && a <= 'z' {
			// 'a' - 'A' = 32 (the difference between the ASCII codes of 'a' and 'A').
			A[i] = a - 'a' + 'A'
		}
	}
	return A
}
