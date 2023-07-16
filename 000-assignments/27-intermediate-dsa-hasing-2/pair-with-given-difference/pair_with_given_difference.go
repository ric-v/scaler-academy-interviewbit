package pairwithgivendifference

func PairWithGivenDifference(A []int, B int) int {
	E := make(map[int]int)
	n := len(A)
	for i := 0; i < n; i++ {
		a := A[i]
		b := A[i] - B
		c := A[i] + B
		if _, ok := E[b]; ok {
			return 1
		}
		if _, ok := E[c]; ok {
			return 1
		}
		E[a] = 1
	}
	return 0
}
