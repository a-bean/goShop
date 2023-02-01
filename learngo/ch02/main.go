package main

import "fmt"

func main()  {
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	//var e int //动态类型，用的时候就会知道 用起来挺麻烦的
	//
	//a = int8(b)
	//
	//var f1 float32 //大约是 3.4e38
	//var f2 float64 //1.8e308
	//
	//f1 = 3
	//f2 = 3.14

	//var c byte //主要适用于存放字符, 存放字符串
	var c2 rune //也是字符
	c2 = 'a'
	//c = 'a' + 1
	fmt.Printf("c=%c", c2)

	var name string
	name = "i am bobby"
	fmt.Println(name)

}

