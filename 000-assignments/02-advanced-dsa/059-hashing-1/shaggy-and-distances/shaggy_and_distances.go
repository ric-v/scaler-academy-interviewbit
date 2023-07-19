package shaggyanddistances

import "math"

func ShaggyAndDistances(A []int) int {
	distance := math.MaxInt64
	m := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if val, ok := m[A[i]]; ok {
			distance = min(distance, i-val)
			m[A[i]] = i
		} else {
			m[A[i]] = i
		}
	}
	if distance == math.MaxInt64 {
		return -1
	} else {
		return distance
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
