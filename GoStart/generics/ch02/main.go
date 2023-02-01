package main

type Mymap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

type Man struct {
}
type Woman struct {
}

type Company[T Man | Woman] struct {
	Name string
	CEO  T
}

type MyChannel[T int | string] chan T

//类型嵌套

type WowStruct[T string | int, S []T] struct {
	A T
	B S
}

//错误用法1, 类型参数不能单独使用
//type CommonType[T int | string] T

// 错误用法2
type CommonType[T interface{ *int } | string] []T

//匿名结构体不支持泛型
//泛型不支持switch判断

//匿名函数不支持泛型

func main() {
	//company := Company[Man]{
	//	Name: "bobby",
	//	CEO: Man{},
	//}

	//company := Company[Woman]{
	//	Name: "bobby",
	//	CEO:  Woman{},
	//}

	//var c MyChannel[string]

	//几种常见的错误

}
