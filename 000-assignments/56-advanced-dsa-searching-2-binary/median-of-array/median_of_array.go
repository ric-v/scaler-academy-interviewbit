package medianofarray

import "math"

func MedianOfArray(A []int, B []int) float64 {
	// DO NOT MODIFY BOTH THE LISTS
	n1 := len(A)
	n2 := len(B)

	if n2 < n1 {
		return MedianOfArray(B, A)
	}
	low := 0
	high := n1

	for low <= high {
		cut1 := (low + high) / 2
		cut2 := (n1+n2+1)/2 - cut1

		var left1 int
		if cut1 == 0 {
			left1 = math.MinInt32
		} else {
			left1 = A[cut1-1]
		}
		var left2 int
		if cut2 == 0 {
			left2 = math.MinInt32
		} else {
			left2 = B[cut2-1]
		}
		var right1 int
		if cut1 == n1 {
			right1 = math.MaxInt32
		} else {
			right1 = A[cut1]
		}
		var right2 int
		if cut2 == n2 {
			right2 = math.MaxInt32
		} else {
			right2 = B[cut2]
		}

		if left1 <= right2 && left2 <= right1 {
			if (n1+n2)%2 == 0 {
				return float64((math.Max(float64(left1), float64(left2)) + math.Min(float64(right1), float64(right2))) / 2.0)
			} else {
				return float64(math.Max(float64(left1), float64(left2)))
			}
		} else if left1 > right2 {
			high = cut1 - 1
		} else {
			low = cut1 + 1
		}
	}
	return 0.0
}
