package pairsumdivisiblebym

func PairSumDivisibleByM(A []int, B int) int {

	var countArr = make([]int, B)
	for _, a := range A {
		countArr[a%B]++
	}

	var ans int
	ans += countArr[0] * (countArr[0] - 1) / 2
	for i := 1; i <= B/2 && i != B-i; i++ {
		ans += (countArr[i] * countArr[B-i]) % 1000000007
	}

	if B%2 == 0 {
		ans += (countArr[B/2] * (countArr[B/2] - 1) / 2) % 1000000007
	}
	return ans % 1000000007
}
