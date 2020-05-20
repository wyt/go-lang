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

// 发送数据
// chan <- val
// ch <- 0          // 将0放入通道中
// ch <- "hello"    // 将字符串放入通道中

// 接收数据
// [var] <-chan

// 定义只读的channel, 接收数据
// read_only := make (<-chan int)

// 定义只写的channel, 发送数据
// write_only := make (chan<- int)

// 可同时读写
// read_write := make (chan int)
