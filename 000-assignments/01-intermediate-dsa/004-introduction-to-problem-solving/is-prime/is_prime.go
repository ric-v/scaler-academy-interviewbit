package isprime

import count "count_factor"

// IsPrime returns 1 if A is prime, else returns 0
// 
// Time Complexity: O(sqrt(n))
func IsPrime(A int64) int {

	if count.CountFactor(int(A)) == 2 {
		return 1
	}
	return 0
}
