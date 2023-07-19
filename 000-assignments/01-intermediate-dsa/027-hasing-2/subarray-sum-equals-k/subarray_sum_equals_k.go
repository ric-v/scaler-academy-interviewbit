package subarraysumequalsk

// SubArraySumEqualsK : https://www.interviewbit.com/problems/subarray-with-given-xor/
// Given an array of integers A and an integer B.
// Find the total number of subarrays having bitwise XOR of all elements equals to B.
// 
// Time Complexity: O(N)
// Space Complexity: O(N)
func SubArraySumEqualsK(A []int, B int) int {

	n := len(A)
	hm := make(map[int]int)
	hm[0] = 1
	ans := 0
	sum := 0

	for i := 0; i < n; i++ {
		sum += A[i]
		if val, ok := hm[sum-B]; ok {
			ans += val
		}
		hm[sum] = hm[sum] + 1
	}

	return ans
}
