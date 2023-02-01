package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai+bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai+bi
	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai+bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as+bs
	default:
		panic("not supported type")
	}
}


func main() {
	re := add("hello ","bobby")
	res, _ := re.(string)
	fmt.Println(strings.Split(res, " "))
	fmt.Println(re)
}
