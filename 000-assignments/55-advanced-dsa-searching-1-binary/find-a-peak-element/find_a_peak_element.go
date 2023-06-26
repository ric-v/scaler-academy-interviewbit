package findapeakelement

// FindAPeakElement finds a peak element in an array
// A peak element is an element that is greater than its neighbors.
// 
// Time complexity: O(log n)
// Space complexity: O(1)
func FindAPeakElement(A []int) int {

	// If there is only one element, it is a peak element
	if len(A) == 1 {
		return A[0]
	}

	// If there are two elements, return the greater one
	if A[0] > A[1] {
		return A[0]
	}

	// If there are two elements, return the greater one
	if A[len(A)-1] > A[len(A)-2] {
		return A[len(A)-1]
	}

	var high int = len(A) - 1
	var low int = 0

	// Binary search for the peak element
	for low < high {
		mid := (high + low) / 2

		// If the middle element is greater than its neighbors, it is a peak element
		// If the middle element is greater than its right neighbor but less than its left neighbor, the peak element is in the left half
		// If the middle element is less than its right neighbor but greater than its left neighbor, the peak element is in the right half
		if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
			return A[mid]
		} else if A[mid] > A[mid+1] && A[mid] < A[mid-1] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return A[high]
}
