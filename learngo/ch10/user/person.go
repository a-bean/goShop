package user

// GetCourse, 获取课程的信息
// 参数：
// 	 c: Course对象
// 返回值：
// 课程的名称
func GetCourse(c Course) string {
	return c.Name
}
