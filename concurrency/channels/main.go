package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	ch := make(chan string)
	var wg sync.WaitGroup

	// Start the store function as a goroutine.
	go store(ch)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- fmt.Sprintf("%d\n", i)
		}(i)
	}

	wg.Wait()

	// Close the channel after all goroutines have finished
	close(ch)

}

func store(ch <-chan string) {

	var s string

	for str := range ch {
		s += str
	}

	os.WriteFile("./log", []byte(s), 0644)

}
