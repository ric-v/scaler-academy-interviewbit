package mergetwosortedarrays

// MergeTwoSortedArrays takes two sorted arrays and returns a single sorted array
// 
// Time Complexity: O(n)
// Space Complexity: O(n)
func MergeTwoSortedArrays(A []int, B []int) []int {
	var p1, p2 int = 0, 0
	var result []int

	// Loop through both arrays and append the smaller value to the result array
	for p1 < len(A) && p2 < len(B) {
		if A[p1] < B[p2] {
			result = append(result, A[p1])
			p1++
		} else {
			result = append(result, B[p2])
			p2++
		}
	}

	// Once one of the arrays is exhausted, append the rest of the other array to the result array
	for p1 < len(A) {
		result = append(result, A[p1])
		p1++
	}
	for p2 < len(B) {
		result = append(result, B[p2])
		p2++
	}
	return result
}
