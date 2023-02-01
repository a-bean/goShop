package main

import "fmt"

func main() {
	//运算符 + - * / % ++ --
	var a, b = 1,2
	var astr, bstr = "hello", "bobby"
	fmt.Println(a+b)
	fmt.Println(astr+bstr)
	fmt.Println(4%2)
	a++
	a += 1
	a = a + 1
	a --
	a -= 1
	a = a-1
	fmt.Println(a)

	//逻辑运算符
	var abool, bbool = true, false
	if abool && bbool {

	}
	if abool || bbool {

	}

	if !abool {

	}

	//位运算符, 性能要求高的时候一般会考虑与运算
	var A = 60
	var B = 13
	//var c *int //指针类型
	//d := &A //取地址 指针
	fmt.Println(A&B)
}