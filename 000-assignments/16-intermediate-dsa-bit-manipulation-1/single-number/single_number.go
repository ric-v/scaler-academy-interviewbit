package singlenumber

// SingleNumber returns the number that appears only once in the array.
// The array will always have exactly one number that appears only once.
// The array will always have at least one number.
// 
// Example: SingleNumber([]int{1, 2, 2, 3, 1}) // 3
// 
// Time complexity: O(n)
// Space complexity: O(1)
func SingleNumber(A []int) int {
	var val int

	// XOR all the numbers in the array.
	// The result will be the number that appears only once.
	// why? because XOR is commutative and associative.
	// 1 ^ 2 ^ 2 ^ 3 ^ 1 = 1 ^ 1 ^ 2 ^ 2 ^ 3 = 0 ^ 0 ^ 3 = 3
	for _, v := range A {
		val ^= v
	}
	return val
}
