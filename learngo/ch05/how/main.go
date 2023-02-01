package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func printSlice(data []string) {
	data[0] = "java"
	for i :=0; i<10; i++ {
		data = append(data, strconv.Itoa(i))
	}
}

type slice struct {
	array unsafe.Pointer   // 用来存储实际数据的数组指针，指向一块连续的内存
	len   int              // 切片中元素的数量
	cap   int              // array数组的长度
}

func main() {
	//go的slice在函数参数传递的时候是值传递还是引用传递： 值传递， 效果又呈现出了引用的效果（不完全是）
	var data []int
	for i:=0;i<2000;i++ {
		data = append(data, i)
		fmt.Printf("len: %d, cap: %d\r\n", len(data), cap(data))
	}
	//data := []int{1,2,3,4,5,6,7,8,9,10}
	//s1 := data[1:6]
	//s2 := data[2:7]
	//
	//fmt.Println(len(s1), cap(s1))
	//fmt.Println(len(s2), cap(s2))
	//
	//fmt.Println(s2)
	//fmt.Println(s1)
}
