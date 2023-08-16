package maxandmin

func MaxAndMin(A []int) int {
	// A[i] * (# of subarrays in which A[i] is max - # of subarrays in which A[i] is min)
	stack := []int{} // contains maximum elements index
	n := len(A)
	ans := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && A[i] > A[stack[len(stack)-1]] {
			popped_index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prev := -1
			if len(stack) > 0 {
				prev = stack[len(stack)-1]
			}
			ans += A[popped_index] * ((popped_index - prev) * (i - popped_index))
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		popped_index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev := -1
		if len(stack) > 0 {
			prev = stack[len(stack)-1]
		}
		ans += A[popped_index] * ((popped_index - prev) * (n - popped_index))
	}

	stack = []int{} // contains minimum elements index
	for i := 0; i < n; i++ {
		for len(stack) > 0 && A[i] < A[stack[len(stack)-1]] {
			popped_index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prev := -1
			if len(stack) > 0 {
				prev = stack[len(stack)-1]
			}
			ans -= A[popped_index] * ((popped_index - prev) * (i - popped_index))
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		popped_index := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev := -1
		if len(stack) > 0 {
			prev = stack[len(stack)-1]
		}
		ans -= A[popped_index] * ((popped_index - prev) * (n - popped_index))
	}
	return ans % (1000000007)
}
