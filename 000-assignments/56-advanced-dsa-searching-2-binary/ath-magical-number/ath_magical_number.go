package athmagicalnumber

func AthMagicalNumber(A, B, C int) int {
	l := Min(B, C)
	h := (A * l)
	lcm := getLCM(B, C)
	mod := 1000000007
	ans := 0
	for l <= h {
		m := ((l + h) / 2)
		if (((m / B) + (m / C)) - (m / lcm)) == A {
			ans = m
			h = (m - 1)
		} else if (((m / B) + (m / C)) - (m / lcm)) > A {
			h = (m - 1)
		} else {
			l = (m + 1)
		}
	}
	ans %= mod
	return int(ans)
}
func getLCM(a int, b int) int {
	x := a
	y := b
	for y != 0 {
		temp := y
		y = (x % y)
		x = temp
	}
	gcd := x
	lcm := ((a * b) / gcd)
	return lcm
}
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
