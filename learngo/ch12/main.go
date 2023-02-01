package main

import (
	"fmt"
	"time"
)

//python,java, php 多线程编程、 多进程编程, 多线程和多进程存在的问题主要是耗费内存
//内存、线程切换， web2.0， 用户级线程， 绿程、轻量级线程、协程， asyncio-python php-swoole java - netty
//内存占用小(2k)、 切换快, go语言的协程， go语言诞生之后就只有协程可用 -goroutine， 非常方便

//主协程
func main() {
	//主死随从

	//匿名函数启动goroutine

	//1. 闭包 2. for循环的问题 for循环的时候 每个变量会重用
	//每次for循环的时候，i变量会被重用， 当我进行到第二轮的for循环的时候 这个i就变了，
	for i := 0; i<100; i++ {
		// 第一种解决办法 tmp := i
		go func(i int) {
			fmt.Println(i)
		}(i) //值传递
	}

	fmt.Println("main goroutine")
	time.Sleep(10*time.Second)
}