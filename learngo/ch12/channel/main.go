package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	不要通过共享内存来通信， 而要通过通信来实现内存共享
	php、python、java、多线程编程的时候，两个goroutine之间通信最常用的方式是 一个全局
	也会提供消息队列的机制， python-queue java、 消费者和生产者之前的关系
	channel 再加上语法糖让使用channel更加简单
	 */
	var msg chan string

	//无缓冲channel适用于 通知， B要第一时间知道A是否已经完成
	//有缓冲channel适用于消费者和生产者之间的通信
	/*
	go中channel的应用场景：
		1. 消息传递、消息过滤
		2. 信号广播
		3. 事件订阅和广播
		4. 任务分发
		5. 结果汇总
		6. 并发控制
		7. 同步和异步
		...
	 */

	//又缓冲和无缓冲的channel
	msg = make(chan string, 0) //channel的初始化值 如果为0的话，你放值进去会阻塞
	//msg = make(chan string, 0) //无缓冲的channel
	//msg = make(chan string, 10) //无缓冲的channel
	go func(msg chan string) { //go有一种happen-before的机制， 可以保障
		data := <- msg
		fmt.Println(data)
	}(msg)

	msg <- "bobby"//放值到channel中
	// waitgroup 如果少了done调用，容易出现deadlock， 无缓冲的channel也容易出现
	time.Sleep(time.Second*10)
}
