package minimumswaps

// MinimumSwaps returns the minimum number of swaps required to bring all the
// elements less than or equal to B together.
//
// Time complexity: O(n)
// Space complexity: O(1)
func MinimumSwaps(A []int, B int) int {
	// Initialize variables
	var count int
	for _, a := range A {
		if a <= B {
			count++
		}
	}

	// Return 0 for edge cases
	if count == 0 || count == 1 || count == len(A) {
		return 0
	}

	// Initialize number of bad elements
	var bad, ans int = 0, 1<<31 - 1

	// Find number of bad elements for the first window
	for i := 0; i < count; i++ {
		if A[i] > B {
			bad++
		}
	}

	// Update the minimum number of bad elements
	if bad < ans {
		ans = bad
	}

	// Initialize the start and end index of the window
	var start, end int = 1, count

	// Loop through the remaining windows
	for end < len(A) {
		// Remove the first element of the previous window
		if A[start-1] > B {
			bad--
		}
		// Add the last element of the current window
		if A[end] > B {
			bad++
		}

		// Update the minimum number of bad elements
		if bad < ans {
			ans = bad
		}

		// Update the start and end index of the window
		start++
		end++
	}
	return ans
}
