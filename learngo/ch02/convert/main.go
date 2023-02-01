package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 12
	//var b = uint8(a)

	//var f float32 = 3.14
	//var c = int32(f)

	var f64 = float64(a)
	fmt.Println(f64)

	//type IT int //类型要求很严格
	//var c IT = IT(a)

	//字符串转数字
	var istr = "12"
	myint, err := strconv.Atoi(istr)
	if err != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)


	var myi = 32
	fmt.Println(strconv.Itoa(myi))

	//字符串转换为float32， 转换为bool
	float, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		return
	}
	fmt.Println(float)

	parseInt, err := strconv.ParseInt("12", 8, 64)
	if err != nil {
		return
	}
	fmt.Println(parseInt)

	parseBool, err := strconv.ParseBool("0")
	if err != nil {
		fmt.Println("ParseBool error")
		return
	}
	fmt.Println(parseBool)


	//基本类型转字符串
	boolStr := strconv.FormatBool(true)
	fmt.Println(boolStr)

	floatStr := strconv.FormatFloat(3.1415926, 'E', -1, 64)
	fmt.Println(floatStr)

	fmt.Println(strconv.FormatInt(42, 16))
}
