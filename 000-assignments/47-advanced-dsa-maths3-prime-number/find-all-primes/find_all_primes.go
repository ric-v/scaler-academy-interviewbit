package findallprimes

import (
	"math"
)

func FindAllPrimes(A int) []int {
	// make []int of size A+1
	primes := make([]int, A+1)

	// initialize all elements to 1
	for i := 0; i <= A; i++ {
		primes[i] = 1
	}

	// 0 and 1 are not prime
	primes[0] = 0
	primes[1] = 0

	// find all primes
	for i := 2; i <= int(math.Sqrt(float64(A))); i++ {
		if primes[i] == 1 {
			for j := 2; i*j <= A; j++ {
				primes[i*j] = 0
			}
		}
	}

	// make a list of primes
	var primesList []int
	for i := 0; i <= A; i++ {
		if primes[i] == 1 {
			primesList = append(primesList, i)
		}
	}
	return primesList
}
