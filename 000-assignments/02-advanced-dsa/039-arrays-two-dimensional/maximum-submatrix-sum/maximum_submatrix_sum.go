package maximumsubmatrixsum

import (
	"math"
)

func MaximumSubmatrixSum(A [][]int) int64 {
	//take suffix sum of the whole array
	n := len(A)
	m := len(A[0])
	pfSum := make([][]int, n)
	for i := 0; i < n; i++ {
		pfSum[i] = make([]int, m)
	}
	pfSum[0][0] = A[0][0]
	//make prefix sum
	//initialize first row
	for i := 1; i < m; i++ {
		pfSum[0][i] = pfSum[0][i-1] + A[0][i]
	}
	//initialize first column
	for i := 1; i < n; i++ {
		pfSum[i][0] = pfSum[i-1][0] + A[i][0]
	}

	//initialize rest of prefix
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			pfSum[i][j] = A[i][j] + pfSum[i-1][j] + pfSum[i][j-1] - pfSum[i-1][j-1]
		}
	}

	sum := int64(0)
	bottomRight := int64(pfSum[n-1][m-1])
	var largest int64 = math.MinInt64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//edge cases checking
			if i == 0 && j == 0 {
				sum = bottomRight
				if sum > largest {
					largest = sum
				}
			} else if i == 0 { //last row subtrced
				sum = bottomRight - int64(pfSum[n-1][j-1])
				if sum > largest {
					largest = sum
				}
			} else if j == 0 { //last column subtrced
				sum = bottomRight - int64(pfSum[i-1][m-1])
				if sum > largest {
					largest = sum
				}
			} else {
				sum = bottomRight - int64(pfSum[i-1][m-1]) - int64(pfSum[n-1][j-1]) + int64(pfSum[i-1][j-1])
				if sum > largest {
					largest = sum
				}
			}
		}
	}
	return largest
}
