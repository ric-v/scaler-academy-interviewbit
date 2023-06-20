package inversioncountinanarray

func InversionCountInAnArray(A []int) int {

	// merge sort the array and count the number of inversions
	_, count := mergeSort(A)
	return count
}

func mergeSort(A []int) ([]int, int) {

	if len(A) == 1 {
		return A, 0
	}

	mid := len(A) / 2

	left, leftCount := mergeSort(A[:mid])
	right, rightCount := mergeSort(A[mid:])

	merged, mergedCount := merge(left, right)

	return merged, leftCount + rightCount + mergedCount
}

func merge(left, right []int) ([]int, int) {

	merged := make([]int, 0)
	count := 0

	for len(left) > 0 && len(right) > 0 {

		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
			count += len(left)
		}
	}

	merged = append(merged, left...)
	merged = append(merged, right...)

	return merged, count
}
