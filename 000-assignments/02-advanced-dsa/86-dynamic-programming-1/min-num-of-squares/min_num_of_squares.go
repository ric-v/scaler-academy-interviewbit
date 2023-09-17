package minnumofsquares

func MinNumOfSquares(A int) int {

	var dp = make([]int, A+1)
	dp[0] = 0
	for i := 1; i <= A; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], 1+dp[i-j*j])
		}
	}
	return dp[A]
}
