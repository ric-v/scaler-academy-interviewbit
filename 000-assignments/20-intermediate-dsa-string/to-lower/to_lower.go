package tolower

// ToLower converts all uppercase letters to lowercase.
// It modifies the slice in place and returns the modified slice.
// It does not allocate any new memory.
//
// Time complexity: O(n)
// Space complexity: O(1)
func ToLower(A []byte) []byte {

	// for each letter in A
	for i, a := range A {

		// if the letter is uppercase (A-Z) then convert it to lowercase (a-z)
		// by adding the difference between the uppercase and lowercase letters
		// 'a' - 'A' = 32 as per ASCII table
		// https://en.wikipedia.org/wiki/ASCII
		if a >= 'A' && a <= 'Z' {
			A[i] = a + ('a' - 'A')
		}
	}
	return A
}
