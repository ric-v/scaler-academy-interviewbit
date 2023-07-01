package squarerootofinteger

// SquareRootOfInteger returns the square root of an integer
//
// Time complexity: O(log n)
// Space complexity: O(1)
func SquareRootOfInteger(A int) int {
	var high, low = A, 0

	// Binary search for the square root
	for low <= high {

		// Find the middle point
		mid := (low + high) / 2

		// If the middle point is the square root, return it
		if mid*mid == A {
			return mid
		}

		// If the middle point is greater than the square root, search the left half
		// Else search the right half
		if mid*mid < A {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// If the square root is not found, return the floor of the square root
	return high
}
