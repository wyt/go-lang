package main

import "fmt"

func f11() (result int) {
	result = 0 //先给返回值赋值
	func() { //再执行defer 函数
		result++
	}()
	return //最后返回
}

func f22() (r int) {
	t := 5
	r = t //赋值指令
	func() { //defer 函数被插入到赋值与返回之间执行，这个例子中返回值r没有被修改
		t = t + 5
	}()
	return //返回
}

func f33() (t int) {
	t = 5 //赋值指令
	func() {
		t = t + 5 //然后执行defer函数,t值被修改
	}()
	return
}
func f44() (r int) {
	r = 1 //给返回值赋值
	func(r int) { //这里的r传值进去的，是原来r的copy，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return
}

func main() {
	fmt.Println(f11())
	fmt.Println(f22())
	fmt.Println(f33())
	fmt.Println(f44())
}
