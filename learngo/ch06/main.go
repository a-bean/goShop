package main

import (
	"fmt"
	"time"
)

//函数参数传递的时候， 值传递，引用传递， go语言中全部都是值传递 指针
func add(items ...int) (sum int, err error) {
	for _, value := range items {
		sum += value
	}
	return sum, err
}
func add2(a, b int){
	fmt.Printf("sum is: %d\r\n", a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func cal(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("不是加法也不是减法")
		}
	}
}

//省略号

func runForever() {
	for {
		time.Sleep(time.Second)
		fmt.Println("doing")
	}
}


func auoIncrement() func() int {
	local := 0 //一个函数中访问另一个函数的局部变量 不行的 闭包
	return func() int{
		local += 1
		return local
	}
}


func main() {
	//go函数支持普通函数、匿名函数、闭包
	/*
	go中函数是“一等公民”
		1. 函数本身可以当做变量
		2. 匿名函数 闭包
		3. 函数可以满足接口
	 */

	funcVar := add

	a := 1
	b := 2
	sum, _ := funcVar(a,b, 3,4)
	fmt.Println(sum)
	fmt.Println(a)

	cal("+")()

	//匿名函数
	localFunc := func(a,b int){
		fmt.Printf("total is: %d\r\n", a+b)
	}
	callback(1, localFunc)

	//闭包
	//有个需求 ，我希望有个函数每次调用一次返回的结果值都是增加一次之后的值
	nextFunc := auoIncrement()
	for i := 0; i<5; i++ {
		fmt.Println(nextFunc())
	}

	nextFunc2 := auoIncrement()
	for i := 0; i<3; i++ {
		fmt.Println(nextFunc2())
	}
}
