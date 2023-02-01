package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
	MyWriter //interface也是一个类型 想放入一个写文件的实现，我想放一个写数据库的实现
}
type fileWriter struct {
	fiePath string
}
type databaseWriter struct {
	host string
	port string
	db string
}

func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil
}

func (dw *databaseWriter) Write(string) error {
	fmt.Println("write string to database")
	return nil
}


func (wc *writerCloser) Close() error {
	fmt.Println("close")
	return nil
}


func main() {
	var mw MyWriter = &writerCloser{
		&databaseWriter{},
	}
	/*
	100行代码
	 */
	mw.Write("bobby")
}

