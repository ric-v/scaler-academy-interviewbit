package countsort2

func CountSort(A []int) []int {

	// find max element in A
	max := A[0]
	for _, v := range A {
		if v > max {
			max = v
		}
	}

	// create count array
	count := make([]int, max+1)

	// count elements in A
	for _, v := range A {
		count[v]++
	}

	// create sorted array
	sorted := make([]int, len(A))

	// fill sorted array
	i := 0
	for k, v := range count {
		for j := 0; j < v; j++ {
			sorted[i] = k
			i++
		}
	}
	return sorted
}
