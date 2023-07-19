package kthsymbolhard

func KthSymbolHard(A int, B int64) int {
	if A == 1 {
		return 0
	}
	res := KthSymbolHard(A-1, B/2)
	if B%2 == 0 {
		return res
	} else {
		return 1 - res
	}
}
