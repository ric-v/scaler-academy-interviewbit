package bulbs

// Bulbs returns the minimum number of switches required to turn all the bulbs on
// toggling a switch will toggle all the bulbs to the right of the switch
// A[i] = 0 means the bulb is off
// A[i] = 1 means the bulb is on
// A switch can be flipped only once
// A switch can be flipped only if the bulb is off
//
// Example: A = [0, 1, 0, 1]
// 1. Switch is flipped at index 0 (A[0] = 0) => A = [1, 0, 1, 0]
// 2. Switch is flipped at index 2 (A[2] = 0) => A = [1, 1, 0, 1]
// 3. Switch is flipped at index 3 (A[3] = 1) => A = [1, 1, 1, 0]
// 4. Switch is flipped at index 3 (A[3] = 0) => A = [1, 1, 1, 1]
// So, the minimum number of switches required to turn all the bulbs on is 4
//
// Time complexity: O(n)
// Space complexity: O(1)
func Bulbs(A []int) int {
	n := len(A)

	// if length of A is 0, return 0
	if n == 0 {
		return 0
	}

	// count is the number of switches required to turn all the bulbs on
	// state is the state of the bulb after the last switch
	count := 0
	state := 0

	// iterate over the array A
	for i := 0; i < n; i++ {

		// if the bulb is off and the state of the bulb is off
		// toggle the switch
		// increment the count
		// toggle the state of the bulb to on
		if A[i] == state {
			count++

			// toggle the state of the bulb to on
			// this is done by subtracting the state from 1
			// to get the opposite of the state
			state = 1 - state
		}
	}
	return count
}
