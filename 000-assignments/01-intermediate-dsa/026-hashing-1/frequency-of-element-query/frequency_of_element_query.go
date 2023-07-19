package frequencyofelementquery

// FreqOfElementQuery returns the frequency of elements in B in A
// A and B are arrays of integers of size N and M respectively
//
// Time complexity: O(N+M)
// Space complexity: O(N)
func FreqOfElementQuery(A, B []int) []int {

	// create a frequency map of A
	var freqMap = make(map[int]int)
	for _, a := range A {
		freqMap[a]++
	}

	// create a result array
	var result = make([]int, len(B))
	for i, b := range B {
		result[i] = freqMap[b]
	}
	return result
}
