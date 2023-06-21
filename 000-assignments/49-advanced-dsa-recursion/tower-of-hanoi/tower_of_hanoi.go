package towerofhanoi

import "math"

// towerOfHanoi is a function that returns the steps to solve the tower of hanoi problem
// 
// Time Complexity: O(2^n)
// Space Complexity: O(2^n)
func towerOfHanoi(A int) [][]int {
	i := 0
	ans := make([][]int, int(math.Pow(2, float64(A))-1))
	toh(A, 1, 3, 2, &i, ans)
	return ans
}

// toh is a helper function that returns the steps to solve the tower of hanoi problem
func toh(N int, S int, D int, E int, i *int, ans [][]int) {
	if N == 0 {
		return
	}
	toh(N-1, S, E, D, i, ans)
	ans[*i] = []int{N, S, D}
	*i++
	toh(N-1, E, D, S, i, ans)
}
