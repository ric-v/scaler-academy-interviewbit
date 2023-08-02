package main

import (
	"context"
	"fmt"
	"time"
)

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go func() {
// 		// simulate work
// 		random := rand.Intn(2)
// 		fmt.Println(random)
// 		if random > 0 {
// 			time.Sleep(1 * time.Second)
// 		} else {
// 			time.Sleep(5 * time.Second)
// 		}
// 		cancel() // sends a cancellation signal
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("Received Done: ", ctx.Err())
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timeout")
// 	}

// }

func main() {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	go func() {

		for {
			deadl, ok := ctx.Deadline()
			fmt.Println(time.Now(), deadl, ok, ctx.Value("key"))
			time.Sleep(1 * time.Second)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Received Done: ", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
	time.Sleep(5 * time.Second)

}
