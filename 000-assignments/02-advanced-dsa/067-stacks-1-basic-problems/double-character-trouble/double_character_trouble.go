package doublecharactertrouble

func DoubleCharacterTrouble(s string) string {
	stack := make([]rune, 0)
	for _, ch := range s {
		if len(stack) == 0 {
			stack = append(stack, ch)
		} else if stack[len(stack)-1] != ch {
			stack = append(stack, ch)
		} else if stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
