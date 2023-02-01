package main

import "fmt"


func mPrint(datas ...interface{}) {
	for _, value := range datas {
		fmt.Println(value)
	}
}

func mPrint2(data interface{}) {
	fmt.Println(data)
}

type myinfo struct {}
func (mi *myinfo) Error() string {
	return "我不是error"
}

func main() {
	//var data = []interface{}{
	//	"bobby", 18, 1.80,
	//}

	var data = []string{
		"bobby", "bobb2", "bobby3",
	} //不可以的
	//var err error
	//err = &myinfo{}

	var datai []interface{}
	for _, value := range data {
		datai = append(datai, value)
	}

	mPrint(datai...)

	//mPrint2(2)
}
