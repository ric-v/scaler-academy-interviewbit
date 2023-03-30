package countofprime

import isprime "is_prime"

// CountOfPrime returns the count of prime numbers from 2 to A
// eg: if A=19, then the count of prime numbers from 2 to 19 is 8
// 2, 3, 5, 7, 11, 13, 17, 19
// hence, the function returns 8
// if A=1, then the count of prime numbers from 2 to 1 is 0
// hence, the function returns 0
//
// Time Complexity: O(n)
func CountOfPrime(A int) int {

	var count int
	for i := 2; i <= A; i++ {
		if isprime.IsPrime(int64(i)) == 1 {
			count++
		}
	}
	return count
}
