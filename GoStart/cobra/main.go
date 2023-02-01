package main

import (
	"flag"
	"fmt"
)

func main() {
	//1. web服务 2. 消费者 3. 工具类 (kratos, goctl)
	// 1. 命令行的方式启动 2. 读取配置文件 3. 启动带参数
	//docker run -p 3306:3306 --name mymysql -v $PWD/conf:/etc/mysql/conf.d -v $PWD/logs:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
	//kubernetes、go-zero、kratos、istio、docker
	// 命令 - 子命令、 参数 flag 配置文件viper
	//启动的时候可以自动生成提示 --help 一些参数可以映射到我们的代码的config中 yaml中的配置是否可以自动映射到config
	//启动的是否可以支持配置文件的检验

	//cobra, pflag, viper 同一个作者， 这三者的集成很简单

	//支持help export
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Printf("Usage: admin-cli [command]")
		return
	}

	switch args[0] {
	case "help":
		fmt.Println("this is help")
	case "export":
		if len(args) == 2 { //导出到文件
			fmt.Println("export to file")
		} else if len(args) == 1 {
			fmt.Println("export to default file")
		}
	default:
		fmt.Println("default")
	}
}
