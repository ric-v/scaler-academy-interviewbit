package submatrixsumqueries

// SubMatrixSumQueries solves the following problem:
// Given a matrix A of size N x M and multiple queries Q, for each query find and return the sum of sub-matrix having its upper-left corner at (B, C) and lower-right corner at (D, E).
//
// Time complexity: O(N*M + Q)
// Space complexity: O(N*M)
func SubMatrixSumQueries(A [][]int, B, C, D, E []int) []int {

	// Precompute prefix sum matrix
	var pfSum [][]int = make([][]int, 0, len(A))
	for i := 0; i < len(A); i++ {
		pfSum = append(pfSum, make([]int, len(A[i])))

		// Compute prefix sum for each row of the matrix
		for j := 0; j < len(A[i]); j++ {
			pfSum[i][j] = A[i][j]
			if i > 0 {
				pfSum[i][j] += pfSum[i-1][j]
			}
			if j > 0 {
				pfSum[i][j] += pfSum[i][j-1]
			}
			if i > 0 && j > 0 {
				pfSum[i][j] -= pfSum[i-1][j-1]
			}
		}
	}

	// Compute the sum of sub-matrix for each query using the prefix sum matrix
	var result []int = make([]int, len(B))
	for i := 0; i < len(B); i++ {

		// Convert to zero-based index for the prefix sum matrix
		a1, b1 := B[i]-1, C[i]-1
		a2, b2 := D[i]-1, E[i]-1

		// Compute the sum of sub-matrix using the prefix sum matrix
		result[i] = pfSum[a2][b2]

		// Subtract the sum of the upper sub-matrix and the left sub-matrix
		if a1 > 0 {
			result[i] -= pfSum[a1-1][b2]
		}

		// Subtract the sum of the lower sub-matrix and the right sub-matrix
		if b1 > 0 {
			result[i] -= pfSum[a2][b1-1]
		}

		// Add the sum of the upper-left sub-matrix to compensate for the double subtraction
		if a1 > 0 && b1 > 0 {
			result[i] += pfSum[a1-1][b1-1]
		}
	}
	return result
}
