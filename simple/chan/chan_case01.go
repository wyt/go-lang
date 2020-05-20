package main

import (
	"fmt"
	"time"
)

func main() {
	case01()
	case02()
}

func case01() {

	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")
}

func case02() {
	c := make(chan int)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
}

//只能向chan里写数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

//只能取channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
