package uniquepathsinagrid

func UniquePathsInAGrid(A [][]int) int {
	var ways = make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		ways[i] = make([]int, len(A[0]))
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				ways[i][j] = 0
			} else if i == 0 && j == 0 {
				ways[i][j] = 1
			} else if i == 0 {
				ways[i][j] = ways[i][j-1]
			} else if j == 0 {
				ways[i][j] = ways[i-1][j]
			} else {
				ways[i][j] = ways[i][j-1] + ways[i-1][j]
			}
		}
	}
	return ways[len(A)-1][len(A[0])-1]
}
