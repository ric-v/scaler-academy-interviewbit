package addonetonumber

func AddOneToNumber(A []int) []int {

	var carry int
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 9 {
			A[i] = 0
			carry = 1
		} else if carry == 1 {
			if A[i] == 9 {
				A[i] = 0
				carry = 1
			}
			A[i] += 1
			carry = 0
			if i > 0 && A[i-1] != 9 {
				break
			}
		} else if i == len(A)-1 {
			A[i] += 1
			break
		}
	}

	if carry != 0 {
		A = append([]int{1}, A...)
	}

	// remove leading zeros
	for i := 0; i < len(A); i++ {
		if A[i] != 0 {
			A = A[i:]
			break
		}
	}
	return A
}
