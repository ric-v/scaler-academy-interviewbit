package isprime

import count "count_factor"

func IsPrime(A int64) int {

	if count.CountFactor(int(A)) == 2 {
		return 1
	}
	return 0
}
