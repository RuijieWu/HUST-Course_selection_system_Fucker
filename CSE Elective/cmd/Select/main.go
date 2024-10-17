package main

import (
	"fmt"
	"time"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/course"
)

func main() {

	c := client.NewFucker()

	t := config.TOKEN
	if t == "" {
		fmt.Println("输入 TOKEN:")
		fmt.Scanln(&t)
	}
	c.SetToken(t)

	id := config.COURSE_ID
	if id == -1 {
		fmt.Println("输入课程 ID:")
		fmt.Scanln(&id)
	}

	for {
		startTime := time.Now()
		err := c.SelectCourse(&course.Course{
			CourseId: id,
		})
		end := time.Now()
		if err != nil {
			fmt.Printf("选课失败: %s, 耗时: %s, now: %s\n", err, end.Sub(startTime), end)
			time.Sleep(100 * time.Millisecond)
		} else {
			fmt.Println("选课成功")
			break
		}
	}
}
