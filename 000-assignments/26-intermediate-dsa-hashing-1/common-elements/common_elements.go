package commonelements

// CommonElements returns the common elements between two arrays.
//
// Time complexity: O(n)
// Space complexity: O(n)
func CommonElements(A, B []int) []int {

	// Count the occurrences of each element in A.
	var countMap = make(map[int]int)
	for _, a := range A {
		countMap[a]++
	}

	// Iterate over B and check if the element is present in the countMap.
	var commonElements []int
	for _, b := range B {

		// If the element is present in the countMap, append it to the commonElements slice. 
		// Decrement the occurrence of the element in the countMap.
		if occ, ok := countMap[b]; ok && occ > 0 {
			commonElements = append(commonElements, b)
			countMap[b]--
		}
	}
	return commonElements
}
