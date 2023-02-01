package main

import (
	"fmt"
	perrors "github.com/pkg/errors"
)

/*
go的error和其他语言的try catch不一样， go语言将错误和异常分开，其他语言 异常
go中认为error是一种值
*/

func divFunc(a, b int) (int, error) {
	if b == 0 {
		return 0, perrors.New("b can't be zero")
	}
	return a / b, nil
}

func main() {
	var a, b = 1, 0
	ret, err := divFunc(a, b)
	if err != nil {
		fmt.Printf("ret is %d, err is %+v\n", ret, err)
	}
}
