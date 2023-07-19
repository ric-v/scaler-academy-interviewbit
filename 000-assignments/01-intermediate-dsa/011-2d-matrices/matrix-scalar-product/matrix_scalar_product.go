package martrixscalarproduct

// MatrixScalarProduct multiplies a matrix by a scalar
// and returns the result
//
// Example: [[1, 2, 3], [4, 5, 6], [7, 8, 9]] * 2 = [[2, 4, 6], [8, 10, 12], [14, 16, 18]]
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func MatrixScalarProduct(A [][]int, B int) (result [][]int) {

	result = make([][]int, len(A))
	// Iterate over the rows of A
	for i := 0; i < len(A); i++ {

		result[i] = make([]int, len(A[i]))
		// Iterate over the columns of A
		for j := 0; j < len(A[i]); j++ {
			// Multiply the element of A by B and store it in result
			result[i][j] = A[i][j] * B
		}
	}
	return
}
