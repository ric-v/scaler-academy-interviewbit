package checkbit

func CheckBit(A, B int) int {
	return (A & (1 << B)) >> B
}
