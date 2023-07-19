package subarraywith0sum

// SubArrayWith0Sum returns 1 if there is a subarray with 0 sum else returns 0
// A = [1, 2, 3, 4, 5, 6] => 0
// A = [1, 2, 3, 4, -4, 6] => 1
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func SubArrayWith0Sum(A []int) int {

	// prefix sum array
	var pfSum = []int{A[0]}
	for i := 1; i < len(A); i++ {
		pfSum = append(pfSum, pfSum[i-1]+A[i])
	}

	// check if there is a subarray with 0 sum in pfSum
	// if the prefix sum value repeats, then there is a subarray with 0 sum
	var pfMap = make(map[int]int)
	for _, pf := range pfSum {
		if _, ok := pfMap[pf]; ok || pf == 0 {
			return 1
		}
		pfMap[pf]++
	}
	return 0
}
