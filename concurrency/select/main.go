package main

import (
	"fmt"
)

func server1(ch chan string) {
	ch <- "from server1"
}

func server2(ch chan string) {
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	// Here we use select to await both of these values simultaneously, printing each one as it arrives.
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
