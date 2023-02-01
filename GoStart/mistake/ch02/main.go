package main

import "fmt"

type Course struct {
	Name string
	Desc string
}

func (c *Course) String() string {
	return c.Name + c.Desc
}

func main() {
	//结构体空指针, 指针类型一定要初始化 nil
	c := new(Course)
	fmt.Println(c.String())
}
