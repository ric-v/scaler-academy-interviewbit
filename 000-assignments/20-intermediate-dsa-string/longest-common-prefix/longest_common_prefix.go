package longestcommonprefix

import "strings"

// LongestCommonPrefix returns the longest common prefix string amongst an array of strings
//
// Example: abcdefgh, aefghijk, abcefgh -> a
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func LongestCommonPrefix(A []string) string {

	// if A is empty, return empty string
	if len(A) == 0 {
		return ""
	}

	// if A has only one element, return that element
	if len(A) == 1 {
		return A[0]
	}

	// set prefix to first element of A
	prefix := A[0]

	// loop through A starting from second element and check if prefix is a substring of each element
	// if prefix is not a substring of an element, remove the last character of prefix and check again
	// if prefix is a substring of all elements, return prefix
	for i := 1; i < len(A); i++ {
		if !strings.Contains(A[i], prefix) {

			// if prefix is not a substring of an element, remove the last character of prefix and check again
			prefix = prefix[:len(prefix)-1]
			i--
		}
	}
	return prefix
}
