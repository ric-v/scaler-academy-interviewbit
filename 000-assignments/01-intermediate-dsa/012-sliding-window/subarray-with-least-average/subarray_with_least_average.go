package subarraywithleastaverage

// SubarrayWithLeastAverage returns the index of the subarray with the least average
// A is the array
// B is the size of the subarray
//
// Time complexity: O(n)
// Space complexity: O(1)
func SubarrayWithLeastAverage(A []int, B int) (result int) {

	// calculate the average of the first B elements
	var sum int
	for i := 0; i < B; i++ {
		sum += A[i]
	}

	// calculate the average of the next B elements
	var start, end, least int = 1, B, sum
	// loop until the end of the array
	for end < len(A) {
		// calculate the average of the next B elements
		sum = (sum - A[start-1]) + A[end]

		// if the average of the next B elements is less than the least average
		if sum < least {
			least = sum
			result = start
		}

		// move the window
		start++
		end++
	}
	return
}
