package singlenumberadvancedii

func SingleNumberAdvancedII(A []int) int {
	var ans int
	for i := 0; i < 32; i++ {

		var c int
		for _, a := range A {
			if a&(1<<uint(i)) != 0 {
				c++
			}
		}

		if c%3 != 0 {
			ans |= 1 << uint(i)
		}
	}
	return ans
}
