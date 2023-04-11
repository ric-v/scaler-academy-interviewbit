package decimaltoanybase

// DecimalToAnyBase converts a decimal number to any base.
// The function takes two parameters: A and B.
// A is the decimal number to convert.
// B is the base to convert the number to.
//
// Time complexity: O(logB(A))
// Space complexity: O(1)
func DecimalToAnyBase(A int, B int) (result int) {

	// Initialize the tens digit to 1.
	var tens int = 1

	// Loop through each digit of A.
	for A > 0 {

		// Multiply the current digit by the current tens digit.
		result += tens * (A % B)

		// Advance the tens digit.
		tens *= 10

		// Divide A by 10.
		A /= B
	}

	return
}
