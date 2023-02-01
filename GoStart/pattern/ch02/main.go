package main

import (
	"sync"
	"sync/atomic"
)

type DBPool struct {
	Host     string
	Port     int
	UserName string
}

var dbPoolIns *DBPool
var lock sync.Mutex
var initialed uint32

//有问题的方法，并发
//加锁, 功能上没有问题，但是性能不好
//高并发下，有bug

//goroutine1 进来， 实例化dbPoolIns = &DBPool{} 进行到一半，goroutine2 进来，dbPoolIns读到dbPoolIns != nil，返回dbPoolIns
func GetDBPool() *DBPool {
	if atomic.LoadUint32(&initialed) == 1 {
		return dbPoolIns
	}
	lock.Lock()
	defer lock.Unlock()

	if initialed == 0 {
		dbPoolIns = &DBPool{}
		atomic.StoreUint32(&initialed, 1)
	}

	return dbPoolIns
}

var once sync.Once

func GetDBPool2() *DBPool {
	once.Do(func() {
		dbPoolIns = &DBPool{}
	})
	return dbPoolIns
}
