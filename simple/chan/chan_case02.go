package main

import (
	"fmt"
	"time"
)

// Send operation on ch1 works!
// Receive operation on ch2 works!
// Exit now!

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) {
		<-ch
	}(ch1)

	go func(ch chan int) {
		ch <- 2
	}(ch2)

	time.Sleep(time.Second)

	for {
		select {
		case ch1 <- 1:
			fmt.Println("Send operation on ch1 works!")
		case <-ch2: // 判断能否从ch2通道中接收数据
			fmt.Println("Receive operation on ch2 works!")
		case ch2 <- 1:
			fmt.Println("Send operation on ch2 works!")
		case <-ch1:
			fmt.Println("Receive operation on ch1 works!")
		default:
			fmt.Println("Exit now!")
			return
		}
	}
}
