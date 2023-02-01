package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//我们新的需求，我可以主动退出监控程序
//共享变量
func cpuIInfo(ctx context.Context)  {
	// 这里能拿到一个请求的id
	fmt.Printf("tracid: %s\r\n", ctx.Value("traceid"))
	//记录一些日志，这次请求是哪个traceid打印的

	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2*time.Second)
			fmt.Println("cpu的信息")
		}
	}
}

func main() {
	//渐进式的方式
	// 有一个goroutine监控cpu的信息
	wg.Add(1)

	//context包提供了三种函数， withCancel， WithTimeout, WithValue
	//如果你的goroutine， 函数中，如果希望被控制， 超时、传值，但是我不希望影响我原来的接口信息的时候，函数参数中第一个参数就尽量的要加上一个ctx
	//1. ctx1, cancel1 := context.WithCancel(context.Background())
	//ctx2, _ := context.WithCancel(ctx1)

	//2. timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	//3. WithDeadline 在时间点cancel

	//4. withValue
	valueCtx := context.WithValue(ctx, "traceid", "gjw12j")
	go cpuIInfo(valueCtx)
	wg.Wait()
	fmt.Println("监控完成")
}
