package maximumsubarray

import "fmt"

func MaximumSubArray(A int, B int, C []int) int {
	var maxSum int

	for i := 0; i < A; i++ {
		sum := C[i] * ((i - 1) * (A - i))
		fmt.Println(sum)
		if sum > maxSum && sum <= B {
			maxSum = sum
		}
	}
	return maxSum
}
