package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//长度计算
	//如果你想知道一个字符串(中文)长度， 如果你只有英文，这个时候直接len
	name := "imooc体系课学习"
	bytes := []rune(name)
	fmt.Println(len(bytes))

	//转义符
	courseName := "hello\r\ngo体系课"
	fmt.Println(courseName)
	fmt.Print("hello\r\n")
	fmt.Print("world\r\n")

	//格式化输出
	username := "bobby"
	age := 18
	address := "北京"
	mobile := "18989898989"

	var ages []int = []int{1,2,3}

	fmt.Println("用户名: " + username, ",年龄: "+strconv.Itoa(age) + ",地址: "+address, ", 电话: "+mobile)//极难维护
	fmt.Printf("用户名:%s, 年龄: %d, ,地址:%s, 电话: %s\r\n", username, age, address, mobile) //这个非常的常用， 但是性能没有上面的好
	userMsg := fmt.Sprintf("用户名:%T, 年龄: %T, ,地址:%s, 电话: %s\r\n", ages, age, address, mobile)
	fmt.Println(userMsg)

	//通过string的builder进行字符串拼接， 高性能
	var builder strings.Builder
	builder.WriteString("用户名:")
	builder.WriteString(username)
	builder.WriteString(",年龄:")
	builder.WriteString(strconv.Itoa(age))
	builder.WriteString(",地址:")
	builder.WriteString(address)
	builder.WriteString(", 电话: ")
	builder.WriteString(mobile)

	re := builder.String()
	fmt.Println(re)
}
