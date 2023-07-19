package main

import (
	"sort"
)

func NobleInteger(A []int) int {
	sort.Ints(A)
	count := len(A) - 1
	for i := 0; i < len(A)-1; i++ {
		if (count-i) == A[i] && A[i] != A[i+1] {
			return 1
		}
	}
	if A[count] == 0 {
		return 1
	}
	return -1
}
