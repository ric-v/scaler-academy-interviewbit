package main

import "fmt"

// PrintATo1 prints A to 1
//
// Time complexity: O(n)
// Space complexity: O(n)
func PrintATo1(A int) {
	fmt.Printf("%d ", A)

	// base case - when A reaches 1 (1 is the last number)
	if A == 1 {
		fmt.Println()
		return
	}

	// decrement A and call PrintATo1 again
	PrintATo1(A - 1)
}

func main() {
	PrintATo1(5)
	PrintATo1(10)
}
