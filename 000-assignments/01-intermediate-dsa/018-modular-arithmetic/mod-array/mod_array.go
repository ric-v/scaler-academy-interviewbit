package modarray

// ModArray returns the remainder of the array A when divided by B
// A is an array of integers
// B is an integer
// ModArray returns an integer which is the remainder of the array A when divided by B
//
// Example: [1, 4, 3] % 2 => 143 % 2 = 1
// Example: [4, 3, 5, 3, 5, 3, 2, 1] % 47 => 43535321 % 47 = 20
//
// Time complexity: O(n)
// Space complexity: O(1)
func ModArray(A []int, B int) int {
	var t int64 = 1
	var ans int64 = 0

	// 1. Iterate over the array A in reverse order
	for i := len(A) - 1; i >= 0; i-- {

		// 2. Multiply the current element by the current power of 10
		ans = (ans + int64(A[i])*t) % int64(B)

		// 3. Update the current power of 10
		t = (t * 10) % int64(B)
	}
	return int(ans)
}
