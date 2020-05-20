package main

import (
	"fmt"
	"time"
)

func main() {
	connectTimeout := 5
	time.Sleep(time.Duration(connectTimeout) * time.Second)
	fmt.Printf("I have slept 5 sec.")
}
