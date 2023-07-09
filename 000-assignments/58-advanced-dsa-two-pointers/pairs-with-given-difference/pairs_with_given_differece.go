package pairswithgivendifference

import "sort"

func PairsWithGivenDifference(A []int, B int) int {
	n := len(A)
	sort.Ints(A)
	p1 := 0
	p2 := 1
	count := 0
	for p2 < n {
		diff := A[p2] - A[p1]
		if diff == B {
			count++
			x := A[p1]
			y := A[p2]
			for p1 < n && A[p1] == x {
				p1++
			}
			for p2 < n && A[p2] == y {
				p2++
			}
		} else if diff > B {
			p1++
			if p1 == p2 {
				p2++
			}
		} else {
			p2++
		}
	}
	return count
}
