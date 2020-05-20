package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(nowInMinute())
	fmt.Println(defaultCollectInterval)
	fmt.Println(defaultEvictInterval)
	fmt.Println(myTestInterval)

	connectTimeout := 5
	time.Sleep(time.Duration(connectTimeout) * time.Second)
	fmt.Printf("I have slept 5 sec.")
}

var nowInMinute = func() int64 {
	return time.Now().Unix() / 60
}
var (
	defaultCollectInterval = time.Second * 10
	defaultEvictInterval   = time.Minute
	myTestInterval         = time.Duration(10) * time.Second
)
