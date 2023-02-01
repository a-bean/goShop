package main

import "fmt"

type Person struct {
	name string
	age int
}

func do() error {
	var err error
	if err == nil {
		fmt.Println("nil interface")
	}
	return err
}

func main() {
	/*
	不同类型的数据零值不一样
	bool       false
	numbers    0
	string     “”
	pointer    nil
	slice 	   nil
	map		   nil
	channel、interface、 function nil

	struct 默认值不是nil、 默认值是具体字段的默认值
	 */

	p1 := Person{
		name: "bobby",
		age:18,
	}

	p2 := Person{
		name: "bobby",
		age:18,
	}
	if p1 == p2 {
		fmt.Println("yes")
	}

	////slice的默认值
	//var ps []Person //nil slice
	//var ps2 = make([]Person, 0) //empty slice
	//if ps2 == nil {
	//	fmt.Println("nil slice")
	//}

	//var m map[string]string //nil map 和 empty map
	var m map[string]string
	var m2 = make(map[string]string, 0)

	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println(m["name"])
	//m["name"] = "bobby"
	m2["name"] = "bobby"

	for key, value := range m2 {
		fmt.Println(key, value)
	}
	//if m2 == nil {
	//	fmt.Println("nil map")
	//}

	err := do()
	fmt.Println(err==nil)
}
