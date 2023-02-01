package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type DBError struct {
	msg string
}

func (d *DBError) Error() string {
	return d.msg
}

var ErrNameEmpty = &DBError{"name can't be empty"}

func (s *Student) SetName(name string) error {
	if name == "" || s.Name == "" {
		return ErrNameEmpty
	}
	s.Name = name
	return nil
}

type Student struct {
	Name string
	Age  int
}

/*
		Wrap() 函数在已有错误基础上同时附加堆栈信息和新提示信息
		WithMessage() 函数在已有错误基础上附加新提示信息
		WithStack() 函数在已有错误基础上附加堆栈信息。在实际中可根据情况选择使用
	 	2xx 3xx 4xx 5xx

		错误码， 错误码应该可以转换成http状态码？ 但是可维护性很差，自动完成？

		错误码设计的原则 ，http响应的策略：
			1. 不论是否成功还是失败，http状态码一律200, facebook
	 			{
					"code": 101010,
					"msg": "internal error",
					"status": 1,0,
					"data":{}
				}
			如果我现在想要引入一个第三方监控系统，protmetheus， 无法完成主动监控， 都会去适配主流的http状态
	      2. 内部错误我们可以转换成对应的http状态码
				twitter
				http code
				{
					"code": 101010,
					"message": "internal error",
					"reference":"http://"
				}

		code设计思路：
			1. code不能和http的code，通过这些code可以看出来错误码是来自哪个项目，来自哪个模块
			2. 请求错误了， 我们不要去通过code判断。应该通过http的code来判断
			3. 最好是能有文档对每个错误码描述， 手动维护还是自动维护？
			4. 错误码可以直接返回到前端http响应， 所以msg不能包含敏感信息，
			5. 数据返回应该是标准
		错误码设计的重要性
			net/http 状态码不够
			采用数据 100111
		10： 服务 api服务，mxshop后台管理系统服务， 最多可以有100个服务
		01： 代表服务下的某个模块 商品管理 商品分类管理 每个服务可以有100个模块
		11： 商品查询不到，商品已经下架 每个模块下支持100个错误码

		错误码映射到http状态码
			404 400
			200 执行成功 get
			201 post数据 新增成功
			401： 认证失败
			403： 权限不足
			404： 资源找不到
			400： 参数错误
		500： 服务器错误

		有一些公共的和业务无关的状态码我们就可以先行创建好
		10  00 通用错误
		10  01 01 02 通用-数据库错误
		10  02 通用-认证失败
		10  03 通用- 编解码错误
		11  01 00商品模块 - 商品不存在，商品新建失败

		错误统一采用大写开头
		msg应该尽量简介 不要暴露过多的信息
*/
func NewStudent() (*Student, error) {
	stu := &Student{
		Age: 18,
	}
	err := stu.SetName("")
	if err != nil {
		return stu, errors.Wrap(err, "set name faild")
	}
	return stu, nil
}

func main() {
	_, e := NewStudent()
	//%+v 显示错误信息和堆栈信息，比如文件名 行号等
	var perr *DBError
	if errors.As(e, &perr) {
		fmt.Println("match")
	}
}
