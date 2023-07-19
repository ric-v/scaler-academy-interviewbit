package computencrmodm

func ComputeNcrModM(A, B, C int) int {
	temp := make([][]int, A+1)
	for i := 0; i <= A; i++ {
		temp[i] = make([]int, B+1)
		for j := 0; j <= B; j++ {
			if i < j {
				temp[i][j] = 0
			} else if j == 0 {
				temp[i][j] = 1
			} else {
				temp[i][j] = (temp[i-1][j] + temp[i-1][j-1]) % C
			}
		}
	}
	return temp[A][B]
}

// factorial computes n! % m
// Time Complexity: O(n)
// Space Complexity: O(1)
func factorial(n int, m int) int {
	product := 1
	for i := 2; i <= n; i++ {
		product = (product * i % m) % m
	}
	return product
}

// powFunction computes a^n % m
// Time Complexity: O(log n)
// Space Complexity: O(1)
func powFunction(a int, n int, m int) int {
	if n == 0 {
		return 1
	}
	he := powFunction(a, n/2, m) % m
	ha := (he * he) % m
	if n%2 == 0 {
		return ha
	} else {
		return ((a * ha) % m)
	}
}
