package computencrmodp

// ComputenCrModP computes nCr % p
// A = n, B = r, C = p
// nCr = n! / (r! * (n-r)!)
// nCr % p = (n! / (r! * (n-r)!)) % p
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func ComputenCrModP(A, B, C int) int {
	nfact := factorial(A, C)
	nrfact := factorial(A-B, C)
	rfact := factorial(B, C)
	nrpow := powFunction(nrfact, C-2, C)
	rpow := powFunction(rfact, C-2, C)
	ans := (((nfact % C) * (nrpow % C)) % C * (rpow % C)) % C
	return int(ans)
}

// factorial computes n! % m
// Time Complexity: O(n)
// Space Complexity: O(1)
func factorial(n int, m int) int {
	product := 1
	for i := 2; i <= n; i++ {
		product = (product * i % m) % m
	}
	return product
}

// powFunction computes a^n % m
// Time Complexity: O(log n)
// Space Complexity: O(1)
func powFunction(a int, n int, m int) int {
	if n == 0 {
		return 1
	}
	he := powFunction(a, n/2, m) % m
	ha := (he * he) % m
	if n%2 == 0 {
		return ha
	} else {
		return ((a * ha) % m)
	}
}
