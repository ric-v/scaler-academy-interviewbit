package rainwatertrapped

func RainWaterTrapped(A []int) int {
	var result int

	// find the max height
	maxHeight := 0
	for _, height := range A {
		if height > maxHeight {
			maxHeight = height
		}
	}

	// leftmax and rightmax
	leftMax := make([]int, len(A))
	rightMax := make([]int, len(A))

	// fill leftmax
	leftMax[0] = A[0]
	for i := 1; i < len(A); i++ {
		if A[i] > leftMax[i-1] {
			leftMax[i] = A[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}

	// fill rightmax
	rightMax[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > rightMax[i+1] {
			rightMax[i] = A[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	// calculate the water
	for i := 0; i < len(A); i++ {
		result += maxHeight - A[i]
	}

	for i := 0; i < len(A); i++ {
		result -= maxHeight - leftMax[i]
	}

	for i := 0; i < len(A); i++ {
		result -= maxHeight - rightMax[i]
	}
	return result
}
