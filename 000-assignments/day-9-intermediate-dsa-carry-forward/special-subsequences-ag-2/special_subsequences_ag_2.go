package specialsubsequencesag2

// SpecialSubsequencesAG2 returns the number of subsequences of the given string
// that contain the character 'A' and the character 'G' in that order.
// The subsequences do not have to be contiguous.
//
// Example: "ABCGAG" -> 3
// Example: "GAB" -> 0
//
// Time complexity: O(n)
// Space complexity: O(1)
func SpecialSubsequencesAG2(A string) (result int64) {
	var countA int

	// Iterate over the string A and count the number of 'A' characters
	// If the current character is 'G', add the number of 'A' characters to the
	// result
	for _, char := range A {
		switch char {
		case 'A':
			countA++
		case 'G':
			result += int64(countA)
		}
	}
	return
}
