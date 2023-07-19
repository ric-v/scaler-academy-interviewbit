package rotatedsortedarraysearch

// RotatedSortedArraySearch searches for an element in a rotated sorted array
// 
// Time complexity: O(log n)
// Space complexity: O(1)
func RotatedSortedArraySearch(A []int, B int) int {

	// high and low are the indexes of the array
	var high, low int = len(A) - 1, 0

	// Binary search
	for low <= high {

		// mid is the index of the middle element
		mid := (low + high) / 2

		// If the element is found, return its index
		if A[mid] == B {
			return mid
		}

		// If the left half is sorted 
		if A[low] <= A[mid] {

			// If the element is in the left half
			if A[low] <= B && B < A[mid] {
				// Search in the left half
				high = mid - 1
			} else {
				// Search in the right half
				low = mid + 1
			}
		} else {

			// If the element is in the right half
			if A[mid] < B && B <= A[high] {
				// Search in the right half
				low = mid + 1
			} else {
				// Search in the left half
				high = mid - 1
			}
		}
	}
	return -1
}
