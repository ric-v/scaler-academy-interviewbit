package deleteone

import "math"

func DeleteOne(A []int) int {
	n := len(A)

	pf := make([]int, n)
	pf[0] = A[0]
	for i := 1; i < n; i++ {
		pf[i] = GCD(pf[i-1], A[i])
	}

	sf := make([]int, n)
	sf[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		sf[i] = GCD(sf[i+1], A[i])
	}

	gcd := -1
	for i := 0; i < n-1; i++ {
		if i == 0 {
			gcd = sf[1]
		} else {
			gcd = int(math.Max(float64(gcd), float64(GCD(pf[i-1], sf[i+1]))))
		}
	}
	return gcd
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
