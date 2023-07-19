package singleelementinsortedarray

// SingleElementInSortedArray finds the single element in a sorted array
// where all other elements appear twice
//
// Time complexity: O(log n)
// Space complexity: O(1)
func SingleElementInSortedArray(A []int) int {

	// Edge cases
	if len(A) == 1 {
		return A[0]
	}
	if A[0] != A[1] {
		return A[0]
	}
	if A[len(A)-1] != A[len(A)-2] {
		return A[len(A)-1]
	}

	var high, low int = len(A) - 1, 0

	// Binary search
	for low < high {
		mid := low + (high-low)/2

		// If mid is not a duplicate and is not adjacent to any other element
		// then it is the single element
		if A[mid] != A[mid-1] && A[mid] != A[mid+1] {
			return A[mid]
		}

		// If mid is even and is equal to its next element
		if A[mid] == A[mid-1] {
			mid = mid - 1
		}

		// If mid is odd and is equal to its previous element
		if mid%2 == 0 {
			low = mid + 2
		} else {
			high = mid - 1
		}
	}

	return A[low]
}
