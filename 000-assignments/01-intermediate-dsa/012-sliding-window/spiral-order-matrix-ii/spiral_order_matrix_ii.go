package spiralordermatrixii

func SpiralOrderMatrixII(A int) (B [][]int) {

	// create a 2D array
	B = make([][]int, A)

	// create a 1D array for each row
	for i := 0; i < A; i++ {
		B[i] = make([]int, A)
	}

	// fill the array with values in spiral order
	var row, col, c int = 0, 0, 1

	// A is the size of the matrix
	for A > 1 {

		// fill the top row
		for i := 1; i < A; i++ {
			B[row][col] = c
			c++
			col++
		}

		// fill the right column
		for i := 1; i < A; i++ {
			B[row][col] = c
			c++
			row++
		}

		// fill the bottom row
		for i := 1; i < A; i++ {
			B[row][col] = c
			c++
			col--
		}

		// fill the left column
		for i := 1; i < A; i++ {
			B[row][col] = c
			c++
			row--
		}
		A -= 2
		row++
		col++
	}

	// fill the center element
	if A == 1 {
		B[row][col] = c
	}
	return
}
