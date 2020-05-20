// +build ignore

package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	i := 2
	fmt.Println("p[i+1:] ==", p[i+1:])
	fmt.Println("p[i:] ==", p[i:])
	copy(p[i+1:], p[i:])

	fmt.Println("p ==", p)
	fmt.Println("p[i+1:] ==", p[i+1:])
	fmt.Println("p[i:] ==", p[i:])

	test01()
}

func test01() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	fmt.Println(slice1)
	fmt.Println(slice2)

	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	//copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中

	fmt.Println(slice1)
	fmt.Println(slice2)
}
