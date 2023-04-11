package transposematrix

// TransposeMatrix returns the transpose of a matrix
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func TransposeMatrix(A [][]int) (result [][]int) {
	// Create a new 2D slice with the dimensions swapped
	var n, m int = len(A), len(A[0])

	// Create a new slice with the dimensions swapped
	result = make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	// Copy the values from the original slice to the new slice
	// Loop over the rows of the original slice
	for i := 0; i < n; i++ {

		// Loop over the columns of the original slice
		for j := 0; j < m; j++ {

			// transpose the values
			result[j][i] = A[i][j]
		}
	}
	return result
}
