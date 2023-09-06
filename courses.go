// courses.go

package main

var Courses = []struct {
	CourseName      string
	CourseNumber    string
	ClassroomNumber string
	People          string
	Contain         string
}{
	{
		CourseName:      "钢笔绘画",
		CourseNumber:    "1437248",
		ClassroomNumber: "202311437248001",
		People:          "100",
		Contain:         "99",
	},
	{
		CourseName:      "逻辑与幽默",
		CourseNumber:    "1434177",
		ClassroomNumber: "202311434177001",
		People:          "220",
		Contain:         "220",
	},
	// 添加更多课程信息
}
