package rowsum

// RowSum returns the sum of each row in a 2D matrix
// A is a 2D matrix
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func RowSum(A [][]int) (sum []int) {

	// sum is a slice of length equal to the number of rows in A
	sum = make([]int, len(A))

	// iterate over each row in A
	for i := 0; i < len(A); i++ {
		// iterate over each column in A
		for j := 0; j < len(A[0]); j++ {

			// add the value at A[i][j] to the sum of the row
			sum[i] += A[i][j]
		}
	}
	return
}
