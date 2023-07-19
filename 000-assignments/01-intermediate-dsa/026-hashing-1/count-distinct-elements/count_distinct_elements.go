package countdistinctelements

// CountDistinctElements returns the count of distinct elements in A
// A = [1, 2, 3, 4, 5, 6] => 6 (all elements are distinct)
// A = [1, 2, 3, 4, 4, 6] => 5 (4 is repeated)
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func CountDistinctElements(A []int) int {

	// count the number of distinct elements in A
	var countMap = make(map[int]int)
	for _, a := range A {
		countMap[a]++
	}
	// return the count of distinct elements
	return len(countMap)
}
