package togglecase

// ToggleCase toggles the case of the string A.
// If the character is lower case, it will be upper case, and vice versa.
//
// Example: "aBc DeF gHi JkL" -> "AbC dEf GhI jKl"
//
// Time complexity: O(n)
// Space complexity: O(n)
func ToggleCase(A string) string {
	// convert string to byte array
	var words []byte = []byte(A)

	// iterate over the byte array and toggle the case
	for i, a := range words {

		// 'a' = 97, 'z' = 122 (lower case) -> 'A' = 65, 'Z' = 90 (upper case)
		// if the character is lower case, convert it to upper case and vice versa
		// 'a' - 'A' = 32 (difference between lower and upper case)
		if a >= 'a' && a <= 'z' {
			words[i] = a - ('a' - 'A')
		} else if a >= 'A' && a <= 'Z' {
			words[i] = a + ('a' - 'A')
		}
	}
	return string(words)
}
