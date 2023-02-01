package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	//假设我有一批商品的id，我现在想要拿到这批商品id的详情，并发启动多个goroutine去拿这批商品的详情
	goodsID := []uint64{1, 2, 3, 4, 5}

	for _, id := range goodsID {
		//值传递
		go func(id uint64) {
			fmt.Println("正在查询商品：" + strconv.Itoa(int(id)))
		}(id)
	}

	time.Sleep(time.Second * 5)
}
