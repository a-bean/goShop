package main

import "fmt"

func deferReturn() (ret int) {
	defer func() {
		ret ++
	}()
	return 10
}

//error， panic， recover
//go语言错误处理的理念， 一个函数可能出错， try catch 去包住这个函数，
//开发函数的人需要有一个返回值去告诉调用者是否成功， go设计者要求我们必须要处理这个err ，代码中大量的会出现 if err != nil
//go设计者认为必须要处理这个error ， 防御编程

func main()  {
	////连接数据库、 打开文件、开始锁， 无论如何 最后都要记得去关闭数据库、关闭文件、解锁
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//
	//defer func() {
	//
	//}()
	//
	//
	//fmt.Println("main")
	//return
	///*
	//
	// */
	ret := deferReturn()
	fmt.Printf("ret = %d\r\n", ret)
}
