package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var once sync.Once
	// var pool sync.Pool
	// var wg sync.WaitGroup

	// once.Do(func() {
	// 	fmt.Println("Hello World : ", 1)
	// })

	// pool.Put("Hello World : 2")
	// pool.Put("Hello World : 3")
	// pool.New = func() interface{} {
	// 	return "Hello World : 4"
	// }
	// fmt.Println(pool.Get())
	// fmt.Println(pool.Get())
	// fmt.Println(pool.Get())

	var num atomic.Int32

	for i := 0; i < 10; i++ {
		// wg.Add(1)
		go func(j int) {

			num.Add(1)
			fmt.Println(num.Load())
			num.Store(int32(j))

			once.Do(func() {
				fmt.Println("Hello World : ", j)
			})
		}(i)
	}
	time.Sleep(time.Second)
	// wg.Wait()
}

func print(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World : ", i)
}
