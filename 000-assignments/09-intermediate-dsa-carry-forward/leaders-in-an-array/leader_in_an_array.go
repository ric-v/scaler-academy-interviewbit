package leadersinanarray

// LeadersInAnArray returns the leaders in an array
// A leader is an element which is greater than all the elements to its right side.
// And the rightmost element is always a leader.
// For example in the array {16, 17, 4, 3, 5, 2}, leaders are 17, 5 and 2.
// Note that the elements 17, 5 and 2 are not in sorted order in the output.
// You can return the elements in any order.
//
// Time complexity: O(n)
// Space complexity: O(n)
func LeadersInAnArray(A []int) (result []int) {
	var max int
	// iterate from right to left and keep track of max
	for i := len(A) - 1; i >= 0; i-- {

		// if current element is greater than max, it is a leader
		if A[i] > max {

			// update max and append to result
			max = A[i]
			result = append(result, A[i])
		}
	}
	return
}
