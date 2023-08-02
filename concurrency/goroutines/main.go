package main

import (
	"fmt"
	"time"
)

func main() {

	for _, i := range []int{1, 2, 3, 4, 5} {
		go print(i)
	}
	time.Sleep(time.Second)
}

func print(i int) {
	fmt.Println("Hello World : ", i)
}
