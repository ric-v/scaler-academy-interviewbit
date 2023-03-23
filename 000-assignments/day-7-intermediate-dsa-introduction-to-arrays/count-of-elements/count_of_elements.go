package countofelements

// CountOfElements returns the number of elements in A that are strictly less than the maximum element in A.
// 
// Time complexity: O(n)
func CountOfElements(A []int) int {
	var c, max int
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
			c = i
		} else if A[i] != max {
			c++
		}
	}
	return c
}

func CountOfElements_Unoptimized(A []int) int {
	var c, max int
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	for i := 0; i < len(A); i++ {
		if A[i] < max {
			c++
		}
	}
	return c
}
