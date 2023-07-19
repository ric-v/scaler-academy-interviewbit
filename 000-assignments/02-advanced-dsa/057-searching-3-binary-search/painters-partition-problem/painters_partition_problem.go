package painterspartitionproblem

func PaintersPartitionProblem(A, B int, C []int) int {

	l := 0
	h := 0
	ans := 0

	// Find the lower and upper bounds of the answer space
	for i := 0; i < len(C); i++ {
		l = max(l, C[i]) // Update the lower bound with the maximum board length
		h += C[i]        // Calculate the upper bound by summing all board lengths
	}

	for l <= h {
		mid := l + (h-l)/2 // Calculate the mid value

		if check(A, C, mid) {
			ans = mid   // Update the answer with the current mid value
			h = mid - 1 // Adjust the upper bound for further search in the lower half
		} else {
			l = mid + 1 // Adjust the lower bound for further search in the upper half
		}
	}

	return (ans * B % 10000003) // Return the result based on the answer and the given multiplier
}

func check(A int, C []int, mid int) bool {
	sum := 0    // Track the sum of board lengths being painted
	worker := 1 // Track the number of workers being used

	for i := 0; i < len(C); i++ {
		if sum+C[i] > mid {
			worker++   // Increment the worker count since the current board cannot be painted within the given time
			sum = C[i] // Reset the sum to the current board's length
			if worker > A {
				return false // If the number of workers exceeds the available workers, return false
			}
		} else {
			sum += C[i] // Add the current board's length to the sum
		}
	}

	return true // All boards can be painted within the given time and using the available workers
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
