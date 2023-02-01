package main

import (
	"fmt"
	"sync"
	"time"
)

//锁本质上是将并行的代码串行化了， 使用lock肯定会影响性能
//即使是设计锁，那么也应该尽量的保证并行
// 我们有两组协程， 其中一组负责写数据，另一个组负责读数据，web系统中绝大部分场景都是读多写少，
// 虽然有多个goroutine，但是仔细分析我们会发现， 读协程之间应该并发， 读和写之间应该串行， 读和读之间也不应该并行
// 读写锁


func main() {
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)
	//写的goroutine
	go func() {
		time.Sleep(time.Second*3)
		defer wg.Done()
		rwlock.Lock() //加写锁， 写锁会防止别的写锁获取，和读锁获取
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		time.Sleep(time.Second*5)
	}()


	// 读的goroutine
	for i:=0; i<5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwlock.RLock() //加读锁， 读锁不会阻止别人的读
				time.Sleep(500*time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}
		}()
	}

	wg.Wait()
}
