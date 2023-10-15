package main

import "fmt"

func solve(A []int) int {
	const MOD = 1000000007
	n := len(A)
	result := 0

	for i := 0; i < n; i++ {
		minVal := A[i]
		maxVal := A[i]

		for j := i; j < n; j++ {
			if A[j] < minVal {
				minVal = A[j]
			}
			if A[j] > maxVal {
				maxVal = A[j]
			}

			result = (result + maxVal - minVal) % MOD
		}
	}

	return (result + MOD) % MOD
}

func main() {
	fmt.Println(solve([]int{4, 7, 3, 8}))
}
