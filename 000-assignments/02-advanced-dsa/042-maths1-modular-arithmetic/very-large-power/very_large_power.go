package verylargepower

func VeryLargePower(A int, B int) int {
	mod := int(1000000007)

	var fact int = 1
	for i := int(2); i <= int(B); i++ {
		fact = (fact * i) % (mod - 1)
	}
	ans := fastPow(A, fact, mod)
	return ans
}

func fastPow(A int, p int, mod int) int {
	if p == 0 {
		return 1
	}

	half := fastPow(A, p/2, mod) % mod

	if p%2 == 0 {
		return (half * half) % mod
	} else {
		return ((((A * half) % mod) * half) % mod)
	}

}
