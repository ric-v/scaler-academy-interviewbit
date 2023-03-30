package secondlargest

// SecondLargest returns the second largest element in the array
//
// Time complexity: O(n)
// Space complexity: O(1)
func SecondLargest(A []int) int {

	// If the array is empty or has only one element, return -1
	var max, secondMax int = A[0], -1

	for _, num := range A {

		// If the current element is greater than the max, update the max and secondMax
		if num > max {
			secondMax = max
			max = num
		}

		// If the current element is greater than the secondMax but less than max, update the secondMax
		if num < max && num > secondMax {
			secondMax = num
		}
	}
	return secondMax
}
