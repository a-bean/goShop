package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁 - 资源竞争
 */

var total int32
var wg sync.WaitGroup
//var lock sync.Mutex

//锁能复制吗, 复制后就失去了锁的效果
func add() {
	defer wg.Done()
	for i := 0; i<1000000; i++ {
		atomic.AddInt32(&total, 1)
		//lock.Lock()
		//total += 1 //竞争
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i<1000000; i++ {
		atomic.AddInt32(&total, -1)
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
