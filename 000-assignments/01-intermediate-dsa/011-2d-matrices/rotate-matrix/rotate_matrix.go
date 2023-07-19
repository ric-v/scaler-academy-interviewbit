package rotatematrix

// RotateMatrix rotates a 2D matrix 90 degrees clockwise.
// The matrix is represented as a pointer to a slice of slices.
//
// Example:
//
//	1 2 3
//	4 5 6
//	7 8 9
//
//	1 4 7
//	2 5 8
//	3 6 9
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func RotateMatrix(A *[][]int) {

	// Transpose the matrix  A
	// Iterate through each row.
	for i := 0; i < len(*A); i++ {
		// Iterate through each column.
		for j := i; j < len((*A)[i]); j++ {
			// Swap the values at the start and end of the row.
			(*A)[i][j], (*A)[j][i] = (*A)[j][i], (*A)[i][j]
		}
	}

	// Iterate through each row.
	for i := 0; i < len(*A); i++ {
		// Define the start and end of the row.
		var row, col int = 0, len(*A) - 1
		// Iterate through the row.
		for row < col {
			// Swap the values at the start and end of the row.
			(*A)[i][row], (*A)[i][col] = (*A)[i][col], (*A)[i][row]
			// Move the start and end to the next item.
			row++
			col--
		}
	}
}
