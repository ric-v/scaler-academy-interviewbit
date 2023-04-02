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

	// for each row
	for r := 0; r < len(*A); r++ {

		// for each column in the row
		for c := r + 1; c < len(*A); c++ {

			if r != c {

				// swap the values at the current row and column
				// with the values at the current column and row
				// to rotate the matrix 90 degrees clockwise
				//
				(*A)[r][c], (*A)[c][r] = (*A)[c][r], (*A)[r][c]
			}
		}
	}
}
