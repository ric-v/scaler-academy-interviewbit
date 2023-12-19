package pubg

// PUBG  is a function that returns the GCD of all the elements of the array A.
//
// Time Complexity: O(NLogN)
// Space Complexity: O(1)
func PUBG(A []int) int {
	ans := A[0]
	N := len(A)
	for i := 1; i < N; i++ {
		ans = GCD(ans, A[i])
	}
	return ans
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
