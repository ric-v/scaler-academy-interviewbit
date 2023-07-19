package minordiagonalsum

// MinorDiagonalSum returns the sum of the minor diagonal of a 2D matrix
// A minor diagonal is the diagonal that runs from the bottom left to the top right
// of a matrix
//
// Time complexity: O(n)
// Space complexity: O(1)
func MinorDiagonalSum(A [][]int) int {

	// Initialize variables
	var sum int

	// Loop through the array
	for row, col := 0, len(A)-1; col >= 0 && row < len(A); {
		// Add the value at the current position to the sum
		sum += A[row][col]
		// Increment row and decrement col
		row++
		col--
	}
	// Return the sum
	return sum
}
