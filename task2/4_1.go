package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func recv(c <-chan int) {
	for v := range c {
		fmt.Printf("recv :%d\n", v)
	}

}

func main() {
	fmt.Println("channel test1")
	ch := make(chan int, 3)

	go send(ch)

	go recv(ch)

	timeout := time.After(1 * time.Second)

	select {
	case <-timeout:
		fmt.Println("timeout")
	case i, ok := <-ch:
		if ok == false {
			fmt.Println("channel close")
			return
		}
		fmt.Println("received:", i)
	default:
		fmt.Println("waitting...")
		time.Sleep(100 * time.Millisecond)
	}

}
