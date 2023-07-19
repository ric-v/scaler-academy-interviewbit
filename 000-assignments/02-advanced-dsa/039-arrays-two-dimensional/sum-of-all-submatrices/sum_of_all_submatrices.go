package sumofallsubmatrices

// SumOfAllSubmatrices solves the following problem:
// Given a matrix of integers A of size N x M.
// Return sum of all submatrices of A.
// NOTE: Since the answer can be very large, return the answer modulo (10^9 + 7).
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func SumOfAllSubmatrices(A [][]int) int {
	var sum int = 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			sum += A[i][j] * (i + 1) * (j + 1) * (len(A) - i) * (len(A[i]) - j)
		}
	}
	return sum
}
