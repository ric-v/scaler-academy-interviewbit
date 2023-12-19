package primemoduloinverse

func PrimeModuloInverse(A int, B int) int {

	// solve with Fermat's little theorem
	// https://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/

	// if A and B are coprime
	// then A^(B-1) % B = 1
	// so A^(B-2) = A^(-1) % B
	return power(A, B-2, B)

}

func power(x int, y int, m int) int {

	if y == 0 {
		return 1
	}

	p := power(x, y/2, m) % m
	p = (p * p) % m

	if y%2 == 0 {
		return p
	} else {
		return (x * p) % m
	}

}

/*
dry run:
A = 3
B = 11

A^(-1) % B = 4

A^(B-1) % B = 1
3^(11-1) % 11 = 1
3^10 % 11 = 1

A^(B-2) % B = A^(-1) % B
3^(11-2) % 11 = 4
3^9 % 11 = 4

3^9 % 11 = 3^3 * 3^3 * 3^3 % 11
3^9 % 11 = 27 * 27 * 27 % 11
3^9 % 11 = 5 * 5 * 5 % 11
3^9 % 11 = 125 % 11
3^9 % 11 = 4

*/
