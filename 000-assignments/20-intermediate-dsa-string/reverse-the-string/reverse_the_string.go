package reversethestring

import "strings"

// ReverseTheString reverses the string A.
// A consists of multiple words separated by space. The task is to reverse the string.
//
// Time complexity: O(n)
// Space complexity: O(n)
func ReverseTheString(A string) string {
	// Split the string into words.
	words := strings.Split(A, " ")

	// Reverse the words.
	var reversedWords []string
	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}
	// Join the words.
	return strings.TrimSpace(strings.Join(reversedWords, " "))
}
