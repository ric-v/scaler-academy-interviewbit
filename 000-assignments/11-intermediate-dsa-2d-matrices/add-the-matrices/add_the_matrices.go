package addthematrices

// AddTheMatrices adds two matrices and returns the result
// If the two matrices are not of the same size, the function returns nil
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func AddTheMatrices(A [][]int, B [][]int) (result [][]int) {

	// Create the result matrix.
	// The result matrix will have the same dimensions as the two
	// input matrices, so we can use the length of A to determine
	// the number of rows in the result matrix.
	result = make([][]int, len(A))
	for i, row := range A {

		// The result matrix has the same number of columns as
		// the input matrix, so we can use the length of row to
		// determine the number of columns in the result matrix.
		result[i] = make([]int, len(row))
		for j := range row {

			// Add the corresponding elements of A and B and
			// store the result in the result matrix.
			result[i][j] = A[i][j] + B[i][j]
		}
	}
	return
}
