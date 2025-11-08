package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var count int
var wg sync.WaitGroup

func add() {
	defer wg.Done()

	mutex.Lock()
	for i := 0; i < 1000; i++ {
		count++
	}
	mutex.Unlock()
}

func main() {
	fmt.Println("mutex test1")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()

	fmt.Printf("count = %d\r\n", count)
}
