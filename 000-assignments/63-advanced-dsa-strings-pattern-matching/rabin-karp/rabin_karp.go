package rabinkarp

func RabinKarp(A, B string) int {
	count := 0
	mod := 1000000007

	// calculate the position vector values so as to use while calculating hash values
	base := 29
	positionVector := make([]int, len(B))
	positionVector[0] = 1
	for i := 1; i < len(B); i++ {
		positionVector[i] = (positionVector[i-1] * base) % mod
	}

	// calculate the hash value for string B
	var BhashValue int
	index := 0
	for i := len(B) - 1; i >= 0; i-- {
		BhashValue = (BhashValue + (positionVector[index]*int(B[i]))%mod) % mod
		index++
	}

	// calculate the hash value for the first window in string A
	var AhashValue int
	p1 := 0
	p2 := len(B) - 1
	index = 0
	for i := p2; i >= p1; i-- {
		AhashValue = (AhashValue + (positionVector[index]*int(A[i]))%mod) % mod
		index++
	}
	if AhashValue == BhashValue {
		count++
	}

	// calculate hashvalue for remaining windows and compare whether it equals to BhashValue or not
	p1 = 1
	p2 = len(B)
	index = 0
	for p2 < len(A) {
		AhashValue = ((((AhashValue-(int(A[p1-1])*positionVector[len(B)-1]))%mod+mod)%mod*base)%mod + (int(A[p2])*positionVector[0])%mod) % mod
		p1++
		p2++
		if AhashValue == BhashValue {
			count++
		}
	}

	return count
}
