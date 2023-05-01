package elementsremoval

import (
	"sort"
)

func ElementsRemoval(A []int) int {
	// reverse sort A
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	// sum all elements in A
	var sum int
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	// calculate min cost of removing elements from A
	var minCost int
	for i := 0; i < len(A); i++ {
		minCost += sum
		sum -= A[i]
	}
	return minCost
}
