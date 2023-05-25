package checkanagrams

// CheckAnagrams checks if two strings are anagrams of each other
// An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Example: "bat" and "tab" are anagrams of each other
//
// Time complexity: O(n)
// Space complexity: O(n)
func CheckAnagrams(A, B string) int {

	// If the length of the two strings are not equal, they cannot be anagrams
	if len(A) != len(B) {
		return 0
	}

	// Create a map of runes and their count in string A
	var anagramMap = make(map[rune]int)

	// Iterate over string A and increment the count of each rune in the map
	for _, a := range A {
		anagramMap[a]++
	}

	// Iterate over string B and decrement the count of each rune in the map
	for _, b := range B {
		anagramMap[b]--

		// If the count of any rune becomes 0, delete it from the map
		if anagramMap[b] == 0 {
			delete(anagramMap, b)
		}
	}
	
	// If the map is not empty, the strings are not anagrams
	if len(anagramMap) > 0 {
		return 0
	}
	return 1
}
