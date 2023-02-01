package main

func main() {
	//map一定要初始化, slice 可以不用初始化
	var course = make(map[string]string, 2)
	course["name"] = "go体系课"

}
