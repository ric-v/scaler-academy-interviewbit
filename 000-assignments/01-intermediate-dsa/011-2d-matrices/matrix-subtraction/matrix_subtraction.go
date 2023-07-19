package matrixsubtraction

// MatrixSubtraction subtracts two matrices and returns the result
// If the two matrices are not of the same size, the function returns nil
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func MatrixSubtraction(A [][]int, B [][]int) (result [][]int) {

	// make a slice of slices with the same length as the input
	result = make([][]int, len(A))

	// iterate over each row
	for i, row := range A {
		// make a new slice with the same length as the row
		result[i] = make([]int, len(row))

		// iterate over each column
		for j := range row {
			// subtract the corresponding value from the other matrix
			// and store the result in the result matrix
			result[i][j] = A[i][j] - B[i][j]
		}
	}
	return
}
