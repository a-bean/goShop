package main

import "fmt"

//class Animal {
//	birthday timestamp
//	legs int
//	mouth int
//
//	walk(){
//
//	}
//}
//
//// duck 必须显示的指定Animal， 隐含的信息： duck也具备了 legs字段， 也具备了birthday字段、也具备了walk方法
//class Duck implents Animal{
//
//}
// 鸭子类型强调的事物的外部行为 而不是内部的结构， 在其他语言

//接口的定义
type Duck interface {
	//方法的申请
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}

func (pd *pskDuck) Walk() {
	fmt.Println("this is pskduck walking")
}

func (pd *pskDuck) Swimming() {
	fmt.Println("this is pskduck swimming")
}


func main() {
	//go语言的接口，鸭子类型， php、python
	//go语言中处处都是interface， 到处都是鸭子类型 duck typing
	/*
	当看到一只鸟走起来像鸭子、游泳起来像鸭子、 叫起来也像鸭子， 那么这只鸟就是鸭子
	动词， 方法， 具备某些方法
	 */

	 var d Duck = &pskDuck{}
	 d.Walk()
}
