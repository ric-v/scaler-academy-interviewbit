package checkpalindrome

// CheckPalindrome checks if A is a palindrome
// A palindrome is a word, phrase, number, or other sequence of characters which reads the same backward as forward, such as madam or racecar.
//
// Time complexity: O(n)
// Space complexity: O(1)
func CheckPalindrome(A string) int {
	// base case 1 - empty string
	if len(A) <= 1 {
		return 1
	}

	// base case 2 - first and last characters are not equal
	if A[0] != A[len(A)-1] {
		return 0
	}

	// recursive case - check if the substring is a palindrome or not
	// A[1:len(A)-1] is the substring of A from index 1 to len(A)-1
	return CheckPalindrome(A[1 : len(A)-1])
}
