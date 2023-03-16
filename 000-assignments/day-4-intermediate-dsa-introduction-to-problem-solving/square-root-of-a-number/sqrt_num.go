package numsqrt

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
