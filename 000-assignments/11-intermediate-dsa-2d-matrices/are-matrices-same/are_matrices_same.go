package arematricessame

// AreMatricesSame checks if two matrices are the same
// and returns 1 if they are and 0 if they are not
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func AreMatricesSame(A [][]int, B [][]int) int {

	// Check if the two matrices have the same number of rows
	if len(A) != len(B) {
		return 0
	}

	// Iterate over the rows
	for i := 0; i < len(A); i++ {

		// Check if the two matrices have the same number of columns
		if len(A[i]) != len(B[i]) {
			return 0
		}

		// Iterate over the columns
		for j := 0; j < len(A[i]); j++ {

			// Check if the values in the two matrices are not the same
			if A[i][j] != B[i][j] {
				return 0
			}
		}
	}
	return 1
}
