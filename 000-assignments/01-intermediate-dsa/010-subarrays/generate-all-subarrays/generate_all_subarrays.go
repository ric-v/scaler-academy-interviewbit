package generateallsubarrays

// GenerateAllSubarrays generates all subarrays of an array
//
// Time complexity: O(n^3)
// Space complexity: O(n^2)
func GenerateAllSubarrays(A []int) [][]int {
	var result [][]int

	// i is the starting index of the subarray
	for i := 0; i < len(A); i++ {

		// j is the ending index of the subarray
		for j := i; j < len(A); j++ {

			var intermediate []int
			// k is the index of the element in the subarray A[i..j]
			for k := i; k <= j; k++ {
				intermediate = append(intermediate, A[k])
			}

			// append the subarray to the result
			result = append(result, intermediate)
		}
	}
	return result
}
