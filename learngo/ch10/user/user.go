// user 包， 改包封装了一些结构体
// author: bobby
// datetime: 20221102
package user

//package 用来组织源码， 是多个go源码的集合，代码复用的基础， fmt， os， io
//每个源码文件开始都必须要申明所属的package
//python 不需要去申明package php c# namespace


// Course, 对课程信息的包装
type Course struct {
	Name string
}