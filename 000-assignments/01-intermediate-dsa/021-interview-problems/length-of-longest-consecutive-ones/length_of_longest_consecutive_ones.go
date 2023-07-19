package lengthoflongestconsecutiveones

func LenOfLongestConsecutiveOnes(A string) int {
	countOf1s := 0
	for i := 0; i < len(A); i++ {
		if A[i] == '1' {
			countOf1s++
		}
	}

	if countOf1s == len(A) {
		return countOf1s
	}

	maxCount := 0
	for i := 0; i < len(A); i++ {
		left := 0
		right := 0

		if A[i] == '0' {
			for j := i - 1; j >= 0; j-- {
				if A[j] == '1' {
					left++
				} else {
					break
				}
			}
			for k := i + 1; k < len(A); k++ {
				if A[k] == '1' {
					right++
				} else {
					break
				}
			}
		}

		var cnt int
		if left+right == countOf1s {
			cnt = left + right
		} else {
			cnt = left + right + 1
		}

		if cnt > maxCount {
			maxCount = cnt
		}
	}

	return maxCount
}
