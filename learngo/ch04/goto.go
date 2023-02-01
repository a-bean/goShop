package main

import "fmt"

/*
goto 语句可以让我们的代码跳到指定的代码块中运行
所以说很少用，

go语言的goto 语句可以实现程序的跳转， goto语句使用场景最多的是程序的错误处理，也就是说当程序出现错误的时候统一跳转到相应的标签处，统一处理
 */

func main() {
	for i := 0; i<5; i++ {
		for j :=0; j<4; j++ {
			if j == 2 {
				goto over
			}
			fmt.Println(i, j)
		}
	}
	over:
		fmt.Println("over")
}

