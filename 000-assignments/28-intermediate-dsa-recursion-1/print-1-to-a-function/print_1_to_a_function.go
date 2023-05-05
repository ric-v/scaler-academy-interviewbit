package main

import "fmt"

// start is the starting number
var start int

// Print prints 1 to A
func solve(A int) {
	start = 1
	Print(A)
}

// Print prints 1 to A recursively
// 
// Time complexity: O(n)
// Space complexity: O(n)
func Print(A int) {

	// base case - when start reaches A (A is the last number)
	if start == A {
		fmt.Printf("%d\n", A)
		return
	}

	// print the current number and increment start
	fmt.Printf("%d ", start)
	start++
	Print(A)
}

func main() {
	solve(5)
	solve(10)
}
