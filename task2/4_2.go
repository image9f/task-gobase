package main

import (
	"fmt"
	"time"
)

func Producer(ch chan<- int, b chan bool) {
	for i := 0; i < 100; i++ {
		ch <- i

		time.Sleep(10 * time.Millisecond)
	}
	close(ch)
	b <- true
}

func Consumer(ch <-chan int, b chan bool) {
	for v := range ch {
		fmt.Printf("consumer recv:%d \r\n", v)
	}
	b <- true
}

func main() {
	fmt.Println("channel test2...")

	ch := make(chan int)
	p := make(chan bool)
	c := make(chan bool)

	go Producer(ch, p)
	go Consumer(ch, c)

	<-p
	<-c
	fmt.Println("channel test2 ok...")
}
