package generateallparanthesis

func generateParenthesis(A int) []string {
	var cnt_open int = 0
	var cnt_close int = 0
	var str string = ""
	var validParanthesis []string

	generateParenthesisHelper(A, cnt_open, cnt_close, str, &validParanthesis)

	return validParanthesis
}

func generateParenthesisHelper(N int, cnt_open int, cnt_close int, str string, vp *[]string) {
	if cnt_open+cnt_close == 2*N { // Base case
		*vp = append(*vp, str)
		return
	}

	if cnt_open < N { // add opening bracket when cnt_open < N
		generateParenthesisHelper(N, cnt_open+1, cnt_close, str+"(", vp)
	}

	if cnt_close < cnt_open { // add closing bracket
		generateParenthesisHelper(N, cnt_open, cnt_close+1, str+")", vp)
	}
}
