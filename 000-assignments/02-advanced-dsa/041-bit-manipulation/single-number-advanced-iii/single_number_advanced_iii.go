package singlenumberadvancediii

// SingleNumberAdvancedIII returns the two numbers that appear only once in the array A.
// The other numbers appear exactly twice.
//
// Time complexity: O(n)
// Space complexity: O(1)
func SingleNumberAdvancedIII(A []int) []int {

	// 1. XOR all the numbers in the array. The result is the XOR of the two numbers that appear only once.
	var xorA int
	for i := 0; i < len(A); i++ {
		xorA ^= A[i]
	}

	// 2. Find the last set bit in the XOR result. This is the bit that is different between the two numbers.
	var lastSetBit int
	for i := 0; i < 32; i++ {

		// Check if the bit is set. If it is, then we have found the last set bit.
		if xorA&(1<<i) != 0 {
			lastSetBit = i
			break
		}
	}

	// 3. XOR all the numbers in the array that have the last set bit set. The result is one of the two numbers.
	// XOR all the numbers in the array that have the last set bit not set. The result is the other number.
	var xorB, xorC int
	for i := 0; i < len(A); i++ {

		// Check if the last set bit is set in the current number. If it is, then XOR it with xorB.
		// Otherwise, XOR it with xorC.
		if A[i]&(1<<lastSetBit) != 0 {
			xorB ^= A[i]
		} else {
			xorC ^= A[i]
		}
	}

	// 4. Return the two numbers in ascending order.
	if xorB > xorC {
		xorB, xorC = xorC, xorB
	}

	// xorB and xorC are the two numbers that appear only once in the array.
	return []int{xorB, xorC}
}
