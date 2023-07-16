package countincreasingtriplets

// CountIncreasingTriplets counts the number of increasing triplets in an array
// of integers.
//
// Time complexity: O(n logn)
// Space complexity: O(n)
func CountIncreasingTriplets(A []int) (count int) {

	// consider each element as the middle element of the triplet
	for i := 0; i < len(A); i++ {

		// count the number of elements to the left of the middle element that
		// are less than the middle element
		var left int
		for j := i - 1; j >= 0; j-- {
			if A[j] < A[i] {
				left++
			}
		}

		// count the number of elements to the right of the middle element that
		// are greater than the middle element
		var right int
		for j := i + 1; j < len(A); j++ {
			if A[j] > A[i] {
				right++
			}
		}

		// the number of increasing triplets with the middle element as the
		// middle element is the product of the number of elements to the left
		// of the middle element that are less than the middle element and the
		// number of elements to the right of the middle element that are
		// greater than the middle element
		count += (left * right)
	}
	return
}
