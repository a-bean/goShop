package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	//panic("this is an panic") //panic会导致程序的退出， 平时开发中不要随便用，一般我们在哪里用到：我们一个服务启动的过程中
	//比如我的服务想要启动，必须有些依赖服务准备好，日志文件存在、mysql能链接通、比如配置文件没有问题，这个时候服务才能启动的时候
	//如果我们的服务启动检查中发现了这些任何一个不满足那就调用panic 主动调用
	//但是你的服务一旦启动了，这个时候你的某行代码中不小心panic那么不好意思你的程序挂了， 这是重大事故
	//但是架不住有些地方我们的代码写的不小心会导致被动触发panic
	//recover 这个函数能捕获到panic
	defer func() {
		if r := recover(); r!= nil {
			fmt.Println("recoverd if A: ", r)
		}
	}()

	panic("this is an panic")
	var names map[string]string
	names["go"] = "go工程师"


	//fmt.Println("this is a func")
	return 0, errors.New("this is an error")

	// 1. defer需要放在panic之前定义， 另外 recover 只有在defer调用的函数中才会生效
	// 2. recover 处理异常后， 逻辑并不会恢复到panic的那个点去
	// 3. 多个defer会形成栈， 后定义的defer会先执行
}


func main() {
	//panic 这个函数会导致你的程序退出， 不推荐随便使用panic
	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}