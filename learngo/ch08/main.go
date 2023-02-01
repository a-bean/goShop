package main

import "fmt"

type Person struct {
	name string
}

//接收者
func (pn *Person) SayHello() {

}

//通过指针交换两个值
func swap(a, b *int) {
	//a篮子中存放b篮子中的值， b篮子中存放a篮子的值 好像没有问题
	//临时值
	t := *a

	//将a地址的值改为b指向的值
	*a = *b

	//将b地址的值改为a指向的值
	*b = t
}

func changeName(p *Person) {
	p.name = "imooc"
}

func main() {
	//指针， 提一个需求，我希望结构体传值的时候 我在函数中修改的值可以反应到变量中
	//p := Person{
	//	name:"bobby",
	//}

	//var pi *Person = &p
	//changeName()
	//fmt.Printf("%p", pi)

	//指针的定义
	//var po *Person
	po := &Person{
		name:"bobby2",
	}

	//(*po).name = "bobby3"
	po.name = "bobby4"
	//po = &p
	//第一个不同的点就出来 第二个点 go语言限制了指针的运算，在c语言中你拿到一个指针之后可以进行+1 在go语言中不行 不能参加运算
	//go的指针是一个阉割版，unsafe包里面，所以说一般我们不会使用unsafe包，但是当你要用到的时候是可以使用
	fmt.Println(po)

	//var a int = 10
	//b := &a

	//指针需要初始化
	//var b *int
	//var p *Person //不行 指针要初始化
	//var ps Person
	//ps := &Person{} //第一种初始化方式
	//var emptyPerson Person
	//pi := &emptyPerson //第二种初始化方式

	//初始化两个关键字， map、 channel、 slice初始化推荐使用make方法
	//指针初始化推荐使用new函数， 指针要初始化否则会出现nil pointer
	// map必须初始化

	var pp = new(Person) //第三种初始化方式
	fmt.Println(pp.name)

	a, b := 1,2
	swap(&a,&b)
	fmt.Println(a, b)
}
