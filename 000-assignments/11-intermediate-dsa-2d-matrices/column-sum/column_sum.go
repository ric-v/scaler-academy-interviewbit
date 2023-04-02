package columnsum

// ColumnSum returns the sum of each column in a 2D matrix
// A is a 2D matrix
// sum is the sum of each column in A
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func ColumnSum(A [][]int) (sum []int) {

	// sum is a slice of length equal to the number of columns in A
	sum = make([]int, len(A[0]))

	// iterate over each row in A
	for i := 0; i < len(A); i++ {

		// iterate over each column in A
		for j := 0; j < len(A[0]); j++ {

			// add the value at A[i][j] to the sum of the column
			sum[j] += A[i][j]
		}
	}
	return
}
