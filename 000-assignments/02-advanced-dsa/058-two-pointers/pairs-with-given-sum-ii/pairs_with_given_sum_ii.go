package pairswithgivensumii

func PairsWithGivenSumII(A []int, B int) int {
	n := len(A)
	p1 := 0
	p2 := n - 1
	count := 0
	mod := 1000000007
	for p1 < p2 {
		if A[p1]+A[p2] > B {
			p2--
		} else if A[p1]+A[p2] < B {
			p1++
		} else {
			if A[p1] == A[p2] {
				dupCount := p2 - p1 + 1
				count += int((dupCount * (dupCount - 1)) / 2)
				break
			} else {
				leftCount := 0
				rightCount := 0
				x := A[p1]
				y := A[p2]
				for A[p1] == x {
					leftCount++
					p1++
				}
				for A[p2] == y {
					rightCount++
					p2--
				}
				count += int(leftCount * rightCount)
			}
		}
	}
	count %= mod
	return count
}
