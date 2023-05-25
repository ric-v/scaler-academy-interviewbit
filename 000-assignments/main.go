package main

import (
	"fmt"
	"strings"
)

/**
 * @input A : String
 * @input B : String
 *
 * @Output string.
 */
//  O(n) + O(n) = O(2n) => O(n)
func ModifyString(str string) string {
	var arr []string
	var arrNew string
	var intArr []string

	s := strings.TrimSpace(str)
	fmt.Println(str)

	for _, j := range s {
		//fmt.Println(s[i], string(j))

		if string(j) >=fmt.Sprint(0) && string(j) <= fmt.Sprint(9) {
			intArr = append(intArr, string(j))

			
		} else {

			arr = append(arr, string(j))
		}

	}


	for j := len(arr)-1; j >= 0; j-- {
		

		arrNew =arrNew+ arr[j]
		
	}

	return arrNew

}

func main() {
	fmt.Println(ModifyString("oll123eH56"))
	// ow ollehrld
}

// o notation order of growth
// O(1) - constant
// O(log n) - logarithmic
// O(n) - linear
// O(n log n) - linearithmic
// O(n^2) - quadratic
// O(n^3) - cubic
// O(2^n) - exponential
// O(n!) - factorial
