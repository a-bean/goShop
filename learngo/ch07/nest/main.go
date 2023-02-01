package main

import (
	"encoding/json"
	"fmt"
)

//大小写的重要性

type Person struct {
	//java和c++中的可访问性 public、 private、protected
	//python 没有的， 可访问，不可访问 规范， 可访问性 针对的是struct的字段、函数、变量等
	Name string `json:"name,omitempty"` //一些情况下很有用 omitempty表示忽略空值
	Age int `json:"age,string" bobby:"go"` //很好用， 其他语言中 java， python， php，实际上你想要做到这个很简单
	Height int `json:"-"`
	//那鬼知道这些有哪些写法呢 本质就是一个tag的字符串
}

type empty struct {} //空结构体看起来没用， 但是实际上有些场景下还是很容易用到的
//我需要一个set类型， go语言中没有set类型 set类型值只放值，而且要唯一
//用map来实现一个set


func (p Person) GetName() string{
	return p.Name + "岁"
}

type Student struct {
	////第一种嵌套方式
	//p Person //组合模式, 组合模式和继承有是有很多的重合的功能, 我们更推荐使用组合模式去写代码, has a 和is a

	//第二种嵌套方式， 匿名嵌套
	Person
	score float32
	name string
}


//接收器有两种形态
func (p *Person) print() {
	//有可能该函数中想要改变p的值， 就是person对象很大， 数据较大的时候考虑使用指针方式
	//p.age = 19
	//fmt.Printf("name:%s, age:%d\r\n", p.name, p.age)
}
//
//func (s StructType)funcName(param1 param1Type, ...)(returnVal1 returnType1, ...){ }

//空结构体不占空间
type MySet map[string]struct{}

//go语言没有面向对象一些class的概念，但是代码的封装，面向对象的思想还有的，
//面向对象的3大特性： 封装、继承 和多态(接口)
//把字段和方法封装到一起 这个方法可以操作这个class上的字段

func main() {
	p := Person{
		"bobby", 18, 180,
	}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	//什么问题 占空间
	var PersonSet = make(map[string]bool)
	PersonSet["bobby1"] = false
	PersonSet["bobby2"] = false
	PersonSet["bobby1"] = true

	for key, _ := range PersonSet {
		fmt.Println(key)
	}

	//t := reflect.TypeOf(p)
	//for _, filedName := range []string{"Name", "Age"} {
	//	field, ok := t.FieldByName(filedName)
	//	if !ok {
	//		panic("cant not found")
	//	}
	//
	//	fmt.Println(field.Tag.Get("bobby"))
	//}

	//p.print()
	//fmt.Println(p.age)
	//s := Student{
	//	Person{
	//		"bobby", 18,
	//	},
	//	95.6,
	//	"bobby1",
	//}
	//s.print()

	//s := Student{}
	//s.p.name = "bobby"
	//s.name = "bobby3"
	//fmt.Println(s)
	//fmt.Println(s.name)
}
