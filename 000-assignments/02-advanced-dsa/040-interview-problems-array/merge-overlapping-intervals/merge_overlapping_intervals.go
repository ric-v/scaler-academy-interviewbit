package mergeoverlappingintervals

import "sort"

type Interval struct {
	start, end int
}

func MergeOverlappingIntervals(A []Interval) []Interval {
	sort.Slice(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})

	res := []Interval{}
	inter := Interval{A[0].start, A[0].end}

	for i := 1; i < len(A); i++ {
		if A[i].start <= inter.end {
			// overlapping
			inter.end = max(inter.end, A[i].end)
		} else {
			// non overlapping
			res = append(res, Interval{inter.start, inter.end})
			inter.start = A[i].start
			inter.end = A[i].end
		}
	}

	res = append(res, Interval{inter.start, inter.end})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
