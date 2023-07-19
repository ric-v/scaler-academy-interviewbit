package main

import "fmt"

func main() {
	print("Hello, World!")
}

func print(s string) {
	if len(s) == 0 {
		return
	}

	print(s[1:])
	fmt.Print(string(s[0]))
}
