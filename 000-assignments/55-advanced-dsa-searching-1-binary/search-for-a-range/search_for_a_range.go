package searchforarange

// SearchForARange searches for the first and last occurence of a given number in a sorted array.
//
// Time complexity: O(log n)
// Space complexity: O(1)
func SearchForARange(A []int, B int) []int {

	// If the array is empty, return [-1, -1]
	if len(A) == 1 && A[0] == B {
		return []int{0, 0}
	}

	// If the array is empty, return [-1, -1]
	left, right := -1, -1
	low, high := 0, len(A)-1

	// Find the first occurence of B in A with binary search
	for low <= high {

		// Find the middle index of the array
		// if element at the middle index is equal to B, check if it is the first occurence
		mid := low + (high-low)/2
		if B == A[mid] && (mid == 0 || A[mid] != A[mid-1]) {
			left = mid
			break
		}

		// binary search to the left and right of the middle index
		if A[mid] < B {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	low = 0
	high = len(A) - 1
	// Find the last occurence of B in A with binary search
	for low <= high {

		// Find the middle index of the array
		// if element at the middle index is equal to B, check if it is the last occurence
		mid := low + (high-low)/2
		if A[mid] == B && (mid == len(A)-1 || A[mid] != A[mid+1]) {
			right = mid
			break
		}

		// binary search to the left and right of the middle index
		if A[mid] <= B {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return []int{left, right}
}
