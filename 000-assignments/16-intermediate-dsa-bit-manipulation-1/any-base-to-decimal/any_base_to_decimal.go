package anybasetodecimal

// AnyBaseToDecimal converts a number from any base to decimal
// A: number to convert
// B: base of the number
// 
// Time complexity: O(log(A))
// Space complexity: O(1)
func AnyBaseToDecimal(A, B int) (result int) {

	var multiplier int = 1
	for A > 0 {

		result = result + (A % 10 * multiplier)
		multiplier *= B
		A /= 10
	}
	return result
}
