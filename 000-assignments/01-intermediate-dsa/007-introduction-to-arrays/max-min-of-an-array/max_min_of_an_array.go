package maxminofanarray

// MaxMinSumOfAnArray returns the sum of the maximum and minimum elements of an array
//
// Time complexity: O(n)
// Space complexity: O(1)
func MaxMinSumOfAnArray(A []int) int {

	var max, min int = A[0], A[0]
	for _, num := range A {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	return max + min
}

// MinMax returns the maximum and minimum elements of an array
//
// Time complexity: O(n)
// Space complexity: O(1)
func MinMax(A []int) (int, int) {

	var max, min int = A[0], A[0]
	for _, num := range A {
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	return min, max
}
