package addbinarystrings

import (
	"math"
	"strings"
)

func AddBinaryStrings(A, B string) string {
	var sb strings.Builder
	var n strings.Builder

	for i := 0; i < int(math.Abs(float64(len(A)-len(B)))); i++ {
		n.WriteByte('0')
	}

	// adding zeros to smaller string to make both length of bits equal, which will be easy to calculate.
	if len(A) > len(B) {
		B = n.String() + B
	} else {
		A = n.String() + A
	}
	carry := 0
	i := len(A) - 1 // as A and B both lengths are equal now
	for i >= 0 {
		if A[i] == '1' && B[i] == '1' {
			if carry == 1 {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
			carry = 1
		} else if A[i] == '1' || B[i] == '1' {
			if carry == 1 {
				sb.WriteByte('0')
			} else {
				sb.WriteByte('1')
			}
		} else {
			if carry == 1 {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
			carry = 0
		}
		i--
	}
	if carry == 1 {
		sb.WriteByte('1')
	}
	return reverse(sb.String())
}

func reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
