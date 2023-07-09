package containerwithmostwater

// For a given array A of N non-negative integers representing vertical lines of unit length and an integer B.
// Find the maximum number of vertical lines which can be contained in this rectangle.
// Note: Lines on edge should not be considered as valid.
//
// Time complexity: O(n)
// Space complexity: O(1)
func ContainerWithMostWater(A []int) int {
	var left, right, ans int = 0, len(A) - 1, 0

	// The idea is to use two pointers, one from the left and one from the right.
	for left < right {

		// The area formed between two lines is always limited by the height of the shorter line.
		ans = max(
			ans,
			min(A[left], A[right])*(right-left),
		)

		// We can't move both pointers at the same time, because it will reduce the area.
		// We need to move the pointer pointing to the shorter line inwards.
		if A[left] < A[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
