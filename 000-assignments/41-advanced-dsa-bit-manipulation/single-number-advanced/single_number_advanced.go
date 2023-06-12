package singlenumberadvanced

func SingleNumberAdvanced(A []int) int {
	var ans int
	for _, a := range A {
		ans ^= a
	}
	return ans
}
