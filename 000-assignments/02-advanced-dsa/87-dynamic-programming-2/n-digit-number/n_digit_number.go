package ndigitnumber

const mod = 1000000007

var dp [][]int

func NDigitNumber(A, B int) int {
	dp = make([][]int, A+1)
	for i := 0; i <= A; i++ {
		dp[i] = make([]int, B+1)
		for j := 0; j <= B; j++ {
			dp[i][j] = -1
		}
	}

	ans := 0
	for i := 1; i < 10; i++ {
		ans += count(A-1, B-i)
		ans %= mod
	}

	return ans
}

func count(digits, sum int) int {
	if digits == 0 && sum == 0 {
		return 1
	}
	if sum < 0 {
		return 0
	}
	if digits <= 0 {
		return 0
	}
	if dp[digits][sum] != -1 {
		return dp[digits][sum]
	}

	ans := 0
	for i := 0; i < 10; i++ {
		ans += count(digits-1, sum-i)
		ans %= mod
	}
	dp[digits][sum] = ans

	return dp[digits][sum]
}
