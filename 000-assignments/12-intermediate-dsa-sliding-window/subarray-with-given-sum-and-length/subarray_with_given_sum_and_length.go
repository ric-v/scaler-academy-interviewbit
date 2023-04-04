package subarraywithgivensumandlength

// SubarrayWithGivenSumAndLength returns 1 if there is a subarray of size B with sum equal to C
// else returns 0
//
// Time complexity: O(n)
// Space complexity: O(1)
func SubarrayWithGivenSumAndLength(A []int, B int, C int) int {

	// find sum of first B elements
	var sum int
	for i := 0; i < B; i++ {
		sum += A[i]
	}

	// if sum is equal to C, return 1
	if sum == C {
		return 1
	}

	// create start and end variables
	var start, end int = 0, B
	for end < len(A) {
		// subtract first element from sum and add next element to sum
		sum = sum - A[start] + A[end]
		if sum == C {
			return 1
		}
		start++
		end++
	}
	return 0
}
