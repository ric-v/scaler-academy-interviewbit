package isalphanum

// IsAlNum returns 1 if A is alphanumeric, 0 otherwise (only alphabets and numbers)
//
// Time complexity: O(n)
// Space complexity: O(1)
func IsAlNum(A []byte) int {
	// loop through A and check if each element is an alphabet or a number
	for _, a := range A {

		// if a is not an alphabet or a number, return 0 (false)
		if !((a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9')) {
			return 0
		}
	}
	return 1
}
