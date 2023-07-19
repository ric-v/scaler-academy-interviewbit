package arrayrotation

import (
	reverse "reverse-in-a-range"
)

func ArrayRotation(A []int, B int) []int {
	B = B % len(A)
	reverse.ReverseInARange(A, 0, len(A)-1)
	reverse.ReverseInARange(A, 0, B-1)
	reverse.ReverseInARange(A, B, len(A)-1)
	return A
}
