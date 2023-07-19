package numsqrt

// Sqrt returns the square root of A
// if A is a perfect square, then return the square root
// else, return -1
//
// Time Complexity: O(n)
func Sqrt(A int) int {

	if A < 0 {
		A = A * -1
	}

	for i := 0; i <= A; i++ {
		if i*i == A {
			return i
		}
	}
	return -1
}
