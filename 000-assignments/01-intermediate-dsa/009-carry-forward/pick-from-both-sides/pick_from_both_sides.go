package pickfrombothsides

// PickFromBothSides picks B elements from each end of the array A, and returns the maximum sum of the elements picked
//
// Time complexity: O(n)
// Space complexity: O(1)
func PickFromBothSides(A []int, B int) (result int) {
	// Check if the array is empty or the number of elements to be picked is less than 1
	if len(A) <= 0 || B <= 0 {
		return 0
	}

	// Pick the first B elements from the array
	for i := 0; i < B; i++ {
		result += A[i]
	}

	// Store the sum of the first B elements in a variable
	intermediateSum := result

	// Set the start and end pointers
	start, end := B-1, len(A)-1

	// Loop while the start pointer is greater than or equal to 0
	for start >= 0 {
		// Remove the element pointed by the start pointer, and add the element pointed by the end pointer
		intermediateSum = (intermediateSum - A[start]) + A[end]

		// Decrement the start pointer and increment the end pointer
		start--
		end--

		// Check if the sum of the elements picked is greater than the current result
		if intermediateSum > result {
			// If yes, update the result
			result = intermediateSum
		}
	}
	return result
}
