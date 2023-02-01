package main

import (
	log2 "GoStart/fuzz/log"
)

func main() {
	//初始日志
	log2.Init(log2.NewOptions())

	log2.Debug("hello")

	/*
		我们自己封装了一个options，用于隔开zap.Config
		日志初始化，Init(options),
		整个过程中调用法看不到zap的信息，整个开发过程解耦
	*/
}
