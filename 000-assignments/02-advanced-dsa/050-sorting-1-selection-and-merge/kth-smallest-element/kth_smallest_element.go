package kthsmallestelement

// KthSmallestElement returns the kth smallest element in the array A
// using selection sort
// 
// Time complexity: O(n^2)
// Space complexity: O(1)
func KthSmallestElement(A []int, B int) int {

	// Selection sort the array
	// Find the smallest element in the array and swap it with the first element
	for i := 0; i < len(A); i++ {
		min := A[i]
		minIndex := i

		// Find the smallest element in the unsorted array
		// and swap it with the first element in the unsorted array
		for j := i; j < len(A); j++ {

			if A[j] < min {
				min = A[j]
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	return A[B-1]
}
