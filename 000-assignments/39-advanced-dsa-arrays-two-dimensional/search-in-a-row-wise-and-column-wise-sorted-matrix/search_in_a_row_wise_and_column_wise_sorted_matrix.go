package searchinarowwiseandcolumnwisesortedmatrix

// SearchInMatrix searches for an element in a row-wise and column-wise sorted matrix.
// The function returns the position of the element if found, otherwise returns -1.
//
// Time complexity: O(n^2)
// Space complexity: O(1)
func SearchInMatrix(A [][]int, B int) int {
	var res int = 0

	// Iterate over the matrix and find the element.
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {

			// If the element is found, calculate the position and compare it with the previous result.
			// If the position is smaller than the previous result, update the result.
			if A[i][j] == B {
				value := (i+1)*1009 + (j + 1)
				if res == 0 || value < res {
					res = value
				}
			}
		}
	}

	// If res is 0, it means that the element was not found.
	if res == 0 {
		return -1
	}
	return res
}
