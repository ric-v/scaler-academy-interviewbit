package count

import "math"

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
// 
// Time Complexity: O(sqrt(n))
func CountFactor(n int) (c int) {

	// since factors can be duplicated with i, n/i formulae,
	// finding the 1st sqrt(n) numbers will suffice in finding the count of factors
	// i = 1 -> sqrt(n)
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {

		// if n is divisible by i, then
		if n%i == 0 {

			// if i == n/i, then only count factor once, since both numbers are same
			// else, both numbers are unique, hence, count as 2.
			// eg: 1, 24 or 3, 8
			if i == n/i {
				c++
			} else {
				c += 2
			}
		}
	}
	return
}

func BruteForce(n int) (c int) {

	// for each i from 1 -> n, if n%i is 0, divisible, c++
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			c++
		}
	}
	return
}
