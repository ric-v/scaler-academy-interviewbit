package rowtocolumnzero

// Make a copy of the array
func RowToColumnZero(A [][]int) (result [][]int) {
	// Create two arrays to store the rows and columns
	// that need to be set to zero.
	var row, col []int

	// Iterate over the matrix, and store the rows
	// and columns that need to be set to zero.
	for i := range A {
		for j := range A[i] {
			if A[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}

	// Iterate over the rows and columns that need to be
	// set to zero, and do it.
	for _, i := range row {
		for j := range A[i] {
			A[i][j] = 0
		}
	}
	for _, j := range col {
		for i := range A {
			A[i][j] = 0
		}
	}
	return A
}
