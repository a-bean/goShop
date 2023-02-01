package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		//我怎么去做到， 应该此处， 等待另一个goroutine来通知我
		<- number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		//我怎么去做到， 应该此处， 等待另一个goroutine来通知我
		<- letter
		if i >= len(str) {
			return
		}
		fmt.Print(str[i:i+2])
		i += 2
		number <- true
	}
}

func main() {
	/*
	使⽤两个goroutine交替打印序列，⼀个goroutine打印数字， 另外⼀个goroutine打印字⺟， 最终效果如下：
	12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
	 */
	go printNum()
	go printLetter()
	number <- true

	time.Sleep(time.Second *100)
}
