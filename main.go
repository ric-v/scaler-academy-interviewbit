package main

import "fmt"

func main() {

	fmt.Println(foo(3, 5))
}

func foo(x int, y int) int {
	if y == 0 {
		return 1
	}
	return bar(x, foo(x, y-1))
}

func bar(x int, y int) int {
	if y == 0 {
		return 0
	}
	return (x + bar(x, y-1))
}
