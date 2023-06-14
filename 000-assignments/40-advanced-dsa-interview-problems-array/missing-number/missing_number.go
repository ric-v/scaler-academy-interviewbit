package missingnumber

import (
	"sort"
)

func MissingNumber(A []int) int {
	sort.Ints(A)
	var i int
	for i = 1; i <= len(A); i++ {
		if A[i-1] != i {
			return i
		}
	}
	return len(A)+1
}
