package stairs

func Stairs(A int) int {
	dp := make([]int, A+1)
	mod := 1000000007
	return climb(A, dp, mod)
}

func climb(A int, dp []int, mod int) int {
	if A == 1 || A == 2 {
		return A
	}

	ans := 0
	if dp[A] != 0 {
		ans = dp[A]
	} else {
		ans = ((climb(A-1, dp, mod) % mod) + (climb(A-2, dp, mod) % mod)) % mod
	}

	dp[A] = ans

	return ans
}
