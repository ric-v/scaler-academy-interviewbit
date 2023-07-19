package rangesumquery

import prefixsum "in-place-prefix-sum"

// RangeSumQuery returns the sum of the elements in the given range for each query in B from A
//
// Time complexity: O(n + m)
// Space complexity: O(n)
func RangeSumQuery(A []int, B [][]int) []int64 {

	// // prefixSum[i] = sum of elements from A[0] to A[i]
	// var prefixSum []int = []int{A[0]}
	// for i := 1; i < len(A); i++ {
	// 	prefixSum = append(prefixSum, prefixSum[i-1]+A[i])
	// }
	prefixsum.InPlacePrefixSum(A)

	// rangeSum[i] = sum of elements from A[B[i][0]] to A[B[i][1]]
	var rangeSum []int64
	for _, query := range B {

		// sum of elements from A[0] to A[query[1]]
		var sum int64

		// if query[0] == 0, then sum of elements from A[0] to A[query[1]] is prefixSum[query[1]]
		// else sum of elements from A[0] to A[query[1]] is prefixSum[query[1]] - prefixSum[query[0]-1]
		if query[0] == 0 {
			sum = int64(A[query[1]])
		} else {
			sum = int64(A[query[1]]) - int64(A[query[0]-1])
		}

		// append sum to rangeSum
		rangeSum = append(rangeSum, sum)
	}
	return rangeSum
}
