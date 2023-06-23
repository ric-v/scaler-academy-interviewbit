package sortedinsertposition

func SortedInsertPosition(A []int, B int) int {

	high := len(A) - 1
	low := 0

	for low <= high {
		mid := (low + high) / 2
		if A[mid] == B {
			return mid
		} else if A[mid] < B {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high + 1
}
