package countsort

// CountSort sorts an array of integers using the Counting Sort algorithm.
//
// Time complexity: O(n+k), where n is the number of elements in the array and k is the range of input.
// Space complexity: O(n+k).
func CountSort(A []int) []int {

	// Find the maximum element in the array. This is used to determine the size of the count array.
	max := A[0]
	for _, val := range A {
		if val > max {
			max = val
		}
	}

	// Create a count array of size max+1 and initialize all elements to 0.
	count := make([]int, max+1)
	for _, val := range A {
		// Increment the count of each element in the array.
		count[val]++
	}

	// Update the count array such that each element at each index stores the sum of previous counts.
	// This is done to determine the position of each element in the output array.
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// Create an output array of the same size as the input array. Iterate over the input array in reverse order.
	// For each element in the input array, place the element at the index determined by the count array.
	// Decrement the count of each element in the count array.
	sorted := make([]int, len(A))
	for i := len(A) - 1; i >= 0; i-- {
		sorted[count[A[i]]-1] = A[i]
		count[A[i]]--
	}

	return sorted
}
