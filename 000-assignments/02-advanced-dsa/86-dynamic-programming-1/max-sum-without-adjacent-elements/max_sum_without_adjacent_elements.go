package maxsumwithoutadjacentelements

func MaxSumWithoutAdjacentElements(A [][]int) int {
	m := len(A[0])
	B := make([]int, m)
	for j := 0; j < m; j++ {
		B[j] = max(A[0][j], A[1][j])
	}

	maxsum := make([]int, m)
	for i := 0; i < m; i++ {
		if i == 0 {
			maxsum[i] = B[0]
		} else if i == 1 {
			maxsum[i] = max(B[i], maxsum[i-1])
		} else {
			maxsum[i] = max(B[i]+maxsum[i-2], maxsum[i-1])
		}
	}
	return maxsum[m-1]
}
