package specialindex

// SpecialIndex returns the number of special indices in the array A.
// A special index is an index i such that A[0] + A[1] + ... + A[i-1] == A[i+1] + A[i+2] + ... + A[A.length - 1].
// Sum of odd indices == sum of even indices after we remove the element at index i.
//
// Time complexity: O(n)
// Space complexity: O(n)
func SpecialIndex(A []int) int {
	var n int = len(A)
	// Prefix sums of odd and even indices.
	var oddPrefixSum, evenPrefixSum = make([]int, n), make([]int, n)
	evenPrefixSum[0] = A[0]

	// Calculate prefix sums.
	for i := 1; i < n; i++ {

		// If i is even, then the sum of even indices is A[i] + evenPrefixSum[i-1].
		// The sum of odd indices is oddPrefixSum[i-1].
		//
		// If i is odd, then the sum of odd indices is A[i] + oddPrefixSum[i-1].
		// The sum of even indices is evenPrefixSum[i-1].
		if i%2 == 0 {
			evenPrefixSum[i] = A[i] + evenPrefixSum[i-1]
			oddPrefixSum[i] = oddPrefixSum[i-1]
		} else {
			oddPrefixSum[i] = A[i] + oddPrefixSum[i-1]
			evenPrefixSum[i] = evenPrefixSum[i-1]
		}
	}

	// Count the number of special indices.
	var count int
	for i := 0; i < n; i++ {

		// The sum of odd indices is oddPrefixSum[n-1] - oddPrefixSum[i].
		// The sum of even indices is evenPrefixSum[n-1] - evenPrefixSum[i].
		// e.g. [2,1,6,4] i = 1
		// oddPrefixSum = [0, 1, 1, 5]
		// evenPrefixSum = [2, 2, 8, 8]
		sumEven := oddPrefixSum[n-1] - oddPrefixSum[i]
		sumOdd := evenPrefixSum[n-1] - evenPrefixSum[i]

		// If i is not 0, then we need to add the sum of odd indices before i.
		// If i is not 0, then we need to add the sum of even indices before i.
		if i > 0 {
			sumEven += evenPrefixSum[i-1]
			sumOdd += oddPrefixSum[i-1]
		}

		// If the sum of odd indices == the sum of even indices, then we have a special index.
		if sumEven == sumOdd {
			count++
		}
	}
	return count
}
