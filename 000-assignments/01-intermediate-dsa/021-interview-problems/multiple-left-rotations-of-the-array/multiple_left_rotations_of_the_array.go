package multipleleftrotationsofthearray

func MultipleLeftRotationsOfTheArray(A, B []int) [][]int {
	n := len(B)
	m := len(A)

	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
		for j := 0; j < m; j++ {
			res[i][j] = A[(j+B[i])%m]
		}
	}

	return res
}
