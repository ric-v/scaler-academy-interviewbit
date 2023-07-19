package sortbycolor

// SortByColor sorts the array of integers A by the color
// For example, given A = [0, 1, 2, 0, 1, 2], the function returns [0, 0, 1, 1, 2, 2]
//
// Time complexity: O(n)
// Space complexity: O(1)
func SortByColor(A []int) []int {

	// sort the array by the color
	low := 0
	mid := 0
	high := len(A) - 1

	for mid <= high {
		if A[mid] == 0 {
			A[low], A[mid] = A[mid], A[low]
			low++
			mid++
		} else if A[mid] == 1 {
			mid++
		} else if A[mid] == 2 {
			A[mid], A[high] = A[high], A[mid]
			high--
		}
	}
	return A
}
