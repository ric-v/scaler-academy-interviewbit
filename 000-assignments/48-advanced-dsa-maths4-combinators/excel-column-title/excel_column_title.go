package excelcolumntitle

import (
	"fmt"
)

func ExcelColTitle(A int) string {
	res := ""
	for A > 0 {
		temp := (A - 1) % 26
		res = fmt.Sprintf("%c", 'A'+temp) + res
		A = (A - 1) / 26
	}
	return res
}
