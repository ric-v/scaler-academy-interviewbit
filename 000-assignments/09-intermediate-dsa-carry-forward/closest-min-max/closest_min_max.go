package closestminmax

import (
	"math"
	minmax "max-min-of-an-array"
)

// ClosestMinMax returns the length of the smallest subarray that contains the
// minimum and maximum values of the given array.
//
// Example: [1, 3, 2] -> 2
// Example: [2, 6, 1, 6, 9] -> 3
//
// Time complexity: O(n)
// Space complexity: O(1)
func ClosestMinMax(A []int) int {
	// Find the minimum and maximum values of the array
	min, max := minmax.MinMax(A)

	// If the minimum and maximum values are the same, the length of the smallest
	// subarray that contains the minimum and maximum values is 1
	if min == max {
		return 1
	}

	// Initialize the length of the smallest subarray that contains the minimum
	// and maximum values to the length of the array A
	// as the worst case scenario
	var l = len(A)
	var mini, maxi int = -1, -1

	// Find the minimum and maximum indices of the array
	for i := len(A) - 1; i >= 0; i-- {
		// If the current value is the minimum value, update the minimum index
		// If the current value is the maximum value, update the maximum index
		if A[i] == max {
			maxi = i
		} else if A[i] == min {
			mini = i
		}

		if mini > -1 && maxi > -1 {

			// Calculate the length of the smallest subarray that contains the minimum
			currLength := int(math.Abs(float64(maxi-mini)) + 1)

			// If the minimum and maximum values are found, calculate the length of the
			// smallest subarray that contains the minimum and maximum values
			// and update the length if the calculated length is smaller than the
			// current length
			if currLength < l {
				l = currLength
			}
		}
	}
	return l
}
