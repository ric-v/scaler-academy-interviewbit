package distinctnumbersinwindow

// DistinctNumbersInWindow : https://www.interviewbit.com/problems/distinct-numbers-in-window/
// You are given an array of N integers, A1, A2 ,…, AN and an integer K.
// Return the of count of distinct numbers in all windows of size K.
// 
// Time Complexity: O(N)
// Space Complexity: O(N)
func DistinctNumbersInWindow(A []int, B int) []int {
	N := len(A)
	ans := make([]int, N-B+1)
	k := 0

	m := map[int]int{}
	for i := 0; i < B; i++ {
		m[A[i]]++
	}
	ans[k] = len(m)
	k++
	for i := 1; i < N-B+1; i++ {
		m[A[i+B-1]]++

		m[A[i-1]]--
		if m[A[i-1]] == 0 {
			delete(m, A[i-1])
		}
		ans[k] = len(m)
		k++
	}
	return ans
}
