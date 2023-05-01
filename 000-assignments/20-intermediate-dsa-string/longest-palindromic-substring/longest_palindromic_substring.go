package longestpalindromicsubstring

// LongestPalindromicSubstring returns the longest palindromic substring in A
//
// Example: aaaabaaa -> aaabaaa
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func LongestPalindromicSubstring(A string) string {
	maxPalindrone := 1
	start := 0
	end := 1

	// for each character in A, check if it is the center of a palindrone of odd length
	for i := 1; i < len(A)-1; i++ {
		length := 0
		s := i - 1
		e := i + 1

		// check if the character is the center of a palindrone of odd length
		// and expand the palindrone to the left and right as long as the characters match
		for s >= 0 && e < len(A) && A[s] == A[e] {
			length += 2
			s--
			e++
		}

		// if the length of the palindrone is greater than the current max, update the max
		if maxPalindrone < length {
			maxPalindrone = length
			start = s + 1
			end = e
		}
	}

	// for each character in A, check if it is the center of a palindrone of even length
	for i := 0; i < len(A)-1; i++ {

		// check if the character is the center of a palindrone of even length
		if A[i] == A[i+1] {
			length := 2
			s := i - 1
			e := i + 2

			// expand the palindrone to the left and right as long as the characters match
			for s >= 0 && e < len(A) && A[s] == A[e] {
				length += 2
				s--
				e++
			}

			// if the length of the palindrone is greater than the current max, update the max
			if maxPalindrone < length {
				maxPalindrone = length
				start = s + 1
				end = e
			}
		}
	}
	return A[start:end]
}
