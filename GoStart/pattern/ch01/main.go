package main

import "fmt"

type DbOptions struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
}

type Option func(*DbOptions)

// 这个函数主要用来设置Host
func WithHost(host string) Option {
	return func(o *DbOptions) {
		o.Host = host
	}
}

func NewOpts(options ...Option) DbOptions {
	//先实例化好dbOptions，填充上默认值
	dbopts := &DbOptions{
		Host: "127.0.0.1",
		Port: 3306,
	}
	for _, option := range options {
		option(dbopts)
	}
	return *dbopts
}

func main() {
	//opts := NewOpts(WithHost("192.168.0.1"))
	opts := NewOpts()
	fmt.Println(opts)

	//函数选项模式大量引用
}
