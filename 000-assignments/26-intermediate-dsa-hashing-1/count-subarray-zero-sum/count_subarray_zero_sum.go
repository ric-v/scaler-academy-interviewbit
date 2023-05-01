package countsubarrayzerosum

// CountSubarrayZeroSum counts the number of subarrays with zero sum
// using prefix sum array and hash map
//
// Time complexity: O(n)
// Space complexity: O(n)
func CountSubarrayZeroSum(A []int) int {

	// prefix sum array
	var pfSum = []int{A[0]}
	for i := 1; i < len(A); i++ {
		pfSum = append(pfSum, pfSum[i-1]+A[i])
	}

	// count subarrays with zero sum using hash map
	var subArrMap = make(map[int]int)
	var count int
	subArrMap[0]++

	// count subarrays with zero sum using hash map
	for _, pf := range pfSum {
		if occ, ok := subArrMap[pf]; ok {

			// if prefix sum already exists in hash map then add its occurence to count
			count += occ
		}
		// add prefix sum to hash map
		subArrMap[pf]++
	}
	return count
}
