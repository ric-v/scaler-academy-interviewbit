package findperfectnumbers

import (
	"math"
)

// iterate until the sqrt of A since, the multiple will keep repeating.
/*
	n = 24 => from i = 1 to sqrt(24)

	i   n/i
	_______
	1	24
	2	12
	3	8
	4	6
	-----
	6	4
	8	3
	12	2
	24	1
*/
// by applying this knowledge, find the other factor of A by doing A/i
// hence, on each iteration do, result = result + i + (A/i), if A/i is not same as A.
func FindPerfectNumbers(A int) int {

	// if A is not a +ve num, return 0
	if A <= 1 {
		return 0
	}

	// set result to 1 since 1 is a factor for any number
	var result int = 1
	for i := 2; i <= int(math.Sqrt(float64(A))); i++ {

		// if the number is divisible by i, add i and A/i to factor sum.
		// eg: if A=24
		// 2+12 & 3+8 & 4+6
		if A%i == 0 {
			result += i + (A / i)
		}
	}

	// check if result is same as A, if not reset result to 0
	if result != A {
		return 0
	} else {
		return 1
	}
}
