package amazingsubarrays

// AmazingSubarrays returns the number of amazing subarrays.
// An amazing subarray is a subarray that starts with a vowel.
// A subarray is a contiguous part of an array.
//
// Example: "ABEC" has 6 amazing subarrays: "A", "AB", "ABE", "ABEC", "B", "BE"
//
// Time complexity: O(n)
// Space complexity: O(1)
func AmazingSubarrays(A string) int {
	// result is the number of amazing subarrays.
	// n is the length of A.
	var result, n = 0, len(A)

	// Iterate over A
	for i, char := range A {

		// If the current character is a vowel, then
		// the number of amazing subarrays that start with
		// the current character is n - i.
		if isVowel(char) {
			result += n - i
		}
	}
	return result % 10003
}

// isVowel returns true if the rune is a vowel.
// Otherwise, it returns false.
func isVowel(r rune) bool {
	switch r {
	case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
		return true
	}
	return false
}
