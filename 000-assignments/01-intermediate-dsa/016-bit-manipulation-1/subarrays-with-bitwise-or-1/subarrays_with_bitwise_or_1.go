package subarrayswithbitwiseor1

// SubArraysWithBitwiseOr returns the number of subarrays of A with bitwise OR equal to B
// A is the length of the array
// B is an array of 0s and 1s
//
// Time complexity: O(n)
// Space complexity: O(1)
func SubArraysWithBitwiseOr(A int, B []int) int64 {

	// The number of subarrays with bitwise OR equal to 1 is the number of 1s in B
	var result, possibleSubarr int64

	// Iterate over B and count the number of 1s
	for i, b := range B {

		// If the current bit is 1, then the number of subarrays with bitwise OR equal to 1
		if b == 1 {
			// plus the number of subarrays that can be formed with the current bit
			possibleSubarr = int64(i + 1)
		}
		// If the current bit is 0, then the number of subarrays with bitwise OR equal to 1
		// is the number of subarrays with bitwise OR equal to 1 in the previous iteration
		result += possibleSubarr
	}
	return result
}
