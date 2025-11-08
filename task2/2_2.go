package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var tm_start time.Time

func timeUsed(start time.Time) {
	tm := time.Since(start)
	fmt.Println("time use", tm)
}

func task(i int) {
	defer wg.Done()
	fmt.Printf("tasd %d running\r\n", i)
}

func main() {
	fmt.Println("goroutine test2 ")

	for i := 0; i < 3; i++ {

		wg.Add(1)
		defer timeUsed(time.Now())
		
		go task(i)

	}
	wg.Wait()

	fmt.Println("goroutine test2 end ")

}
