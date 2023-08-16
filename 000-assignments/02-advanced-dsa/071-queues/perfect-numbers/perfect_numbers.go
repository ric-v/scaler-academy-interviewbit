package perfectnumbers

import "strings"

func PerfectNumbers(A int) string {
	queue := []string{"1", "2"}

	for i := 0; i < A-1; i++ {
		s := queue[0]
		queue = queue[1:]
		queue = append(queue, s+"1")
		queue = append(queue, s+"2")
	}

	firstElement := queue[0]
	return firstElement + strings.Join(reverse(firstElement), "")
}

func reverse(s string) []string {
	var result []string
	for _, v := range s {
		result = append([]string{string(v)}, result...)
	}
	return result
}
