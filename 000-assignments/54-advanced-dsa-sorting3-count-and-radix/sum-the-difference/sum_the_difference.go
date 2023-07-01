package sumthedifference

import "sort"

func SumTheDifference(A []int) int {
	N := len(A)
	var min int64 = 0
	var max int64 = 0
	var mod int64 = 1000000007
	var ans int64 = 0

	// by sorting the array, the difference between min & max won't change
	sort.Ints(A)
	// Taking power Array to calculate the power in terms of 2 each and every iteration
	pow := make([]int64, N)
	pow[0] = 1
	for i := 1; i < N; i++ {
		pow[i] = ((int64(pow[i-1])) * 2) % (int64(mod)) // Eg: 2^3 = 2^2 * 2
	}

	// calculating the min of each and every subsequence in the Arrays
	for i := 0; i < N; i++ {
		min = min + (int64(A[i]) * pow[N-1-i])
		min = min % mod
	}
	// calculating the max of each and every subsequence in the Arrays
	for i := N - 1; i >= 0; i-- {
		max = max + (int64(A[i]) * pow[i])
		max = max % mod
	}

	// calculating the min-max
	ans = (max - min + mod) % mod
	return int(ans)
}
