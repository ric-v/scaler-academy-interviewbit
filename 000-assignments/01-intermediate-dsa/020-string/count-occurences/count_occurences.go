package countoccurences

// CountOccurences counts the number of occurences of the word "bob" in a string
//
// Example: "bobob" has 2 occurences of the word "bob" in it
//
// Time complexity: O(n)
// Space complexity: O(1)
func CountOccurences(A string) int {

	// If the length of the string is less than 3, it cannot contain the word "bob"
	var count int
	var isBob []rune

	// Iterate over the string and check if the current letter is 'b' or 'o'
	for _, a := range A {
		// if previous word is "bo" and current letter is 'b' then its bob
		if string(isBob) == "bo" && a == 'b' {
			count++
			isBob = []rune{'b'}
			continue
		}

		if string(isBob) == "b" && a == 'o' {
			isBob = append(isBob, a)
		} else {
			isBob = []rune{}
		}
		if a == 'b' {
			isBob = append(isBob, a)
		}
	}
	return count
}
