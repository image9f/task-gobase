package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("mutex test2")

	var wg sync.WaitGroup
	var count int64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("count=%d\r\n", count)

}
