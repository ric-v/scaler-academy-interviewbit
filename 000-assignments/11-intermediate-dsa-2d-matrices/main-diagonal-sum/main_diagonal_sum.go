package maindiagonalsum

// MainDiagonalSum returns the sum of the main diagonal of a 2D matrix.
// The main diagonal is the diagonal from the top left to the bottom right.
// The matrix is guaranteed to be square.
// Main diagonal will have an index of [i][i] for all i. e.g.: [0][0], [1][1], [2][2], etc.
//
// Time complexity: O(n) where n is the number of rows in the matrix.
// Space complexity: O(1) since we are not creating any new data structures.
func MainDiagonalSum(A [][]int) (sum int) {

	// Iterate through the matrix and add the value at the main diagonal to the sum.
	for i := 0; i < len(A); i++ {

		// Add the value at the main diagonal to the sum.
		// The main diagonal is the diagonal from the top left to the bottom right.
		sum += A[i][i]
	}
	return
}
