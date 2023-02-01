package main

import "fmt"

func main() {
	//go折中， slice 切片 - 数组
	var courses []string
	fmt.Printf("%T\r\n", courses)

	//这个方法很特别， 第一次用很头疼， 原理
	//courses = append(courses, "go")
	//courses = append(courses, "grpc")
	//courses = append(courses, "gin")
	//
	//fmt.Println(courses[1])
	//
	////切片的初始化 3种：1：从数组直接创建 2：使用string{} 3: make
	//allCourses := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	//courseSlice := allCourses[0:len(allCourses)] //左闭右开 [) python
	//fmt.Println(courseSlice)

	//courseSlice := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}

	////make 函数
	//allCourses2 := make([]string, 3)
	//allCourses2[0] = "c"
	//allCourses2[1] = "c"
	//allCourses2[2] = "c"
	//fmt.Println(allCourses2)
	//
	//var allCourses3 []string
	//allCourses3 = append(allCourses3, "c")
	//fmt.Println(allCourses3)

	//访问切片的元素， 访问单个， 访问多个
	//fmt.Println(courseSlice[1])
	//
	////[start:end]
	///*
	//1. 如果只有start 没有end 就表示从start开始到结尾的所有数据
	//2. 如果没有start 有end 就表示从0到end之前的所有数据
	//3. 如果有start 没有有end
	//4. 有start 有end
	// */
	//fmt.Println(courseSlice[1:4])
	//fmt.Println(courseSlice[1:])
	//fmt.Println(courseSlice[:])

	//courseSlice := []string{"go", "grpc"}
	//courseSlice2 := []string{"mysql", "es", "gin"}
	//courseSlice = append(courseSlice, courseSlice2[1:]...)
	//
	////如何删除slice中的元素： 比较麻烦
	//fmt.Println(courseSlice)

	courseSlice := []string{"go", "grpc", "mysql", "es", "gin"}
	myslice := append(courseSlice[:2], courseSlice[3:]...)
	fmt.Println(myslice)

	courseSlice = courseSlice[:3]
	fmt.Println(courseSlice)

	//复制 slice
	//courseSliceCopy := courseSlice
	courseSliceCopy2 := courseSlice[:]
	fmt.Println(courseSliceCopy2)

	var courseSliceCopy = make([]string, len(courseSlice))
	copy(courseSliceCopy, courseSlice)
	fmt.Println(courseSliceCopy)

	fmt.Println("------------------------")
	courseSlice[0] = "java"
	fmt.Println(courseSliceCopy2)
	fmt.Println(courseSliceCopy)
}

