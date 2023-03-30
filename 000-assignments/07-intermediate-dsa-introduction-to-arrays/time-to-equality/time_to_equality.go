package timetoequality

// TimeToEquality returns the number of seconds required to make all elements equal
//
// Time complexity: O(n)
// Space complexity: O(1)
func TimeToEquality(A []int) int {
	var max, sum int = A[0], 0
	for _, num := range A {
		sum += num
		if num > max {
			max = num
		}
	}
	return max*len(A) - sum
}

func TimeToEquality_Unoptimized(A []int) int {
	var max, secondsReq int = A[0], 0

	// Find the max element in the array
	for _, num := range A {
		if num > max {
			max = num
		}
	}

	// Find the number of seconds required to make all elements equal to max
	for _, num := range A {
		secondsReq += max - num
	}
	return secondsReq
}
