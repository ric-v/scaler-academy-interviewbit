package fibonaccinumber

// init dp array with -1
var dp []int

// FibonacciNumber returns the nth fibonacci number
// Time complexity: O(n)
// Space complexity: O(n)
func FibonacciNumber(n int) int {

	// init dp array
	dp = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = -1
	}

	return fib(n)
}

// fib returns the nth fibonacci number
// Time complexity: O(n)
// Space complexity: O(n)
func fib(n int) int {
	if n <= 1 {
		return n
	}

	if dp[n] != -1 {
		return dp[n]
	}

	dp[n] = fib(n-1) + fib(n-2)
	return dp[n]
}
