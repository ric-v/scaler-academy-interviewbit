package dungeonprincess

func DungeonPrincess(A [][]int) int {
	var n, m = len(A), len(A[0])
	var ans = make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				ans[i][j] = max(1, 1-A[i][j])
			} else if i == n-1 { // last row
				x := ans[i][j+1]
				ans[i][j] = max(x-A[i][j], 1)
			} else if j == m-1 {
				x := ans[i+1][j]
				ans[i][j] = max(x-A[i][j], 1)
			} else {
				x := min(ans[i+1][j], ans[i][j+1])
				ans[i][j] = max(x-A[i][j], 1)
			}
		}
	}
	return ans[0][0]
}
