package nearestsmallerelement

func NearestSmallerElement(A []int) []int {
	ans := make([]int, len(A))
	st := make([]int, 0, len(A))
	for i := 0; i < len(A); i++ {
		for len(st) > 0 && A[st[len(st)-1]] >= A[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			ans[i] = -1
		} else {
			ans[i] = A[st[len(st)-1]]
		}
		st = append(st, i)
	}
	return ans
}
