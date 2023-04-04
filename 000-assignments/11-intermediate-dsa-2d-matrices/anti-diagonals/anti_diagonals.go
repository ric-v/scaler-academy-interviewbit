package antidiagonals

// AntiDiagonals returns the anti-diagonals of a matrix
// in a clockwise direction
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func AntiDiagonals(A [][]int) (result [][]int) {

	// Check if the matrix is empty
	if len(A) == 0 || len(A[0]) == 0 {
		return
	}

	// Outer loop to traverse the first row
	for i := 0; i < len(A[0]); i++ {
		// Initialize the row and column indices
		var row, col int = 0, i
		// Initialize a slice to hold the intermediate results
		var intermediate []int

		// Inner loop to traverse the diagonal
		for row < len(A) && col >= 0 {
			// Append the element to the intermediate slice
			intermediate = append(intermediate, A[row][col])
			// Increment the row index and decrement the column index
			row++
			col--
		}

		// Append zeros to the intermediate slice to make the length of the slice equal to the length of the matrix
		for len(intermediate) < len(A) {
			intermediate = append(intermediate, 0)
		}

		// Append the intermediate slice to the result slice
		result = append(result, intermediate)
	}

	// Outer loop to traverse the last column
	for i := 1; i < len(A); i++ {
		// Initialize the row and column indices
		var row, col int = i, len(A[0]) - 1
		// Initialize a slice to hold the intermediate results
		var intermediate []int

		// Inner loop to traverse the diagonal
		for row < len(A) && col >= 0 {
			// Append the element to the intermediate slice
			intermediate = append(intermediate, A[row][col])
			// Increment the row index and decrement the column index
			row++
			col--
		}

		// Append zeros to the intermediate slice to make the length of the slice equal to the length of the matrix
		for len(intermediate) < len(A) {
			intermediate = append(intermediate, 0)
		}

		// Append the intermediate slice to the result slice
		result = append(result, intermediate)
	}

	return
}
