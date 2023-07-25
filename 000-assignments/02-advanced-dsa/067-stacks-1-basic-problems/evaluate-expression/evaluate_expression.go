package evaluateexpression

import "strconv"

func EvaluateExpression(A []string) int {
	n := len(A)
	if n == 1 {
		return 0
	}
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		s := A[i]
		if isOperand(s) {
			t, _ := strconv.Atoi(s)
			stack = append(stack, t)
		} else {
			if len(stack) == 0 {
				return 0
			}
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			c := perform(a, b, s)
			stack = append(stack, c)

		}
	}
	return stack[len(stack)-1]
}

func isOperand(o string) bool {
	if o == "%" || o == "/" || o == "*" || o == "+" || o == "-" {
		return false
	}
	return true
}

func perform(a int, b int, ch string) int {
	if ch == "+" {
		return a + b
	}
	if ch == "-" {
		return b - a
	}
	if ch == "/" {
		return b / a
	}

	return a * b

}
