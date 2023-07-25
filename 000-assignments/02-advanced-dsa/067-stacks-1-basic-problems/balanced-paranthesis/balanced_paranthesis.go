package balancedparanthesis

func BalancedParanthesis(A string) int {

	re := 0
	stack := make([]rune, 0)

	for _, ch := range A {
		if ch == '[' || ch == '{' || ch == '(' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return 0
			}
			top := stack[len(stack)-1]

			if isMatch(ch, top) {
				stack = stack[:len(stack)-1]
			} else {
				return 0
			}
		}
	}
	if len(stack) == 0 {
		re = 1
	}
	return re
}

func isMatch(ch rune, top rune) bool {
	if top == '[' && ch == ']' {
		return true
	}
	if top == '{' && ch == '}' {
		return true
	}
	if top == '(' && ch == ')' {
		return true
	}
	return false
}
