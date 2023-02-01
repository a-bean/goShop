package main

import "fmt"

func a() (int, bool){
	return 0, false
}

var name = "bobby" //全局

func main() {
	//匿名变量, 变量的作用域
	var _ int
	_, ok := a()
	if ok {
		//打印
	}

	var mainName = "main"
	fmt.Println(mainName)

	//iota, 特殊常量， 可以认为是一个可以被编译器修改的常量
	const (
		ERR1  = iota+1
		ERR2
		ERR25 = "ha" //iota内部仍然会增加计数器
		ERR3 //iota+1
		ERR31 //iota+1
		ERR32 //iota+1
		ERR33 //iota+1
		ERR34 = 100 //iota+1
		ERR4 = iota
	)
	const (
		ERRNEW1 = iota
	)

	fmt.Println(ERR1, ERR2, ERR25, ERR3, ERR4)
	fmt.Println(ERRNEW1)

	{
		localName := "local"
		fmt.Println(localName)
	}

	//var mname string
	if name == "bobby" {
		mname := "imooc"
		fmt.Println(mname)
	}else{
		//mname := "imooc"
	}

	/*
	如果中断了iota那么必须显式的恢复， 后续会自动递增
	自增类型默认是int类型
	iota能简化const类型的定义
	每次出现const的时候， iota初始化为0
	 */
}
