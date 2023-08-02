package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var m sync.Mutex
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

	var num int = 0

	for i := 0; i < 100; i++ {
		// wg.Add(1)
		go func(j int) {

			m.Lock()
			num = num + 1
			m.Unlock()

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
