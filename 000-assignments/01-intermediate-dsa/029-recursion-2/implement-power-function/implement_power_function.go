package implementpowerfunction

// PowerFunction is a recursive function that takes three integers A, B and C as arguments and returns (A^B) % C.
// 1 <= A, B, C <= 5 * 10^6 and B is non-negative.
//
// Example: A = 2, B = 3, C = 3 => (2^3) % 3 = 2 => 8 % 3 = 2 => 2
//
// Time Complexity: O(log(B))
// Space Complexity: O(log(B))
func PowerFunction(A, B, C int) int {
	// B = B % C;
	if A == 0 && B == 0 {
		return 0
	} else if B == 0 {
		return 1
	} else if A == 0 {
		return 0
	}

	p := PowerFunction(A%C, B/2, C)

	if B%2 == 0 {
		return (p%C*p%C + C) % C
	} else {
		return (((p % C * p % C) * A % C) + C) % C
	}
}
