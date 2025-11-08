package main

import (
	"fmt"
	"sync"
)

func printEven(wg *sync.WaitGroup) {
	fmt.Println("even")
	wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func printOdd(wg *sync.WaitGroup) {
	fmt.Println("odd")
	wg.Done()
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("goroutine test")

	var wg sync.WaitGroup
	wg.Add(2)
	go printEven(&wg)
	go printOdd(&wg)

	wg.Wait()
	fmt.Println("goroutine test end")

	//test2()

}
