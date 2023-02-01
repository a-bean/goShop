package main

import (
	"fmt"
	"strings"
)


func main()  {
	////字符串的比较
	//a := "hello"
	//b := "bello"
	//fmt.Println(a!=b)
	//
	////字符串的大小比较
	//fmt.Println(a>b)

	//是否包含
	name := "imooc体系课go工程师go,go"
	fmt.Println(strings.Contains(name, "go"))

	//出现次数
	fmt.Println(strings.Count(name, "o"))

	//分割字符串
	fmt.Println(strings.Split(name, "-"))

	//字符串是否包含前缀，是否包含后缀
	fmt.Println(strings.HasPrefix(name, "imooc"))
	fmt.Println(strings.HasSuffix(name, "工程师"))

	//查找子串出现的位置
	fmt.Println(strings.Index(name, "go"))

	//子串替换
	fmt.Println(strings.Replace(name, "go", "java", 2))

	//大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("java"))

	//去掉特殊字符
	fmt.Println(strings.Trim("#$hello #go#", "#$"))
}
