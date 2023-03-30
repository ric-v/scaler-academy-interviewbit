package linearsearchpmultipleoccurences

// LinearSearchMultipleOccurences returns the number of occurences of a number in an array
//
// Time complexity: O(n)
// Space complexity: O(1)
func LinearSearchMultipleOccurences(A []int, B int) int {
	var occ int
	for _, num := range A {
		if num == B {
			occ++
		}
	}
	return occ
}
