package main

import (
	"fmt"
	"time"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/course"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/utils"
)

func main() {

	c := client.NewFucker()

	t := config.TOKEN
	if t == "" {
		utils.Info("[*] 输入 TOKEN:")
		fmt.Scanln(&t)
	}
	c.SetToken(t)

	targets, err := c.GetCourses()
	if err != nil {
		utils.Info("[!] 获取课程列表失败: ", err)
		return
	}

	utils.Info("[*] 课程列表:")
	for index, target := range *targets {
		utils.Info(
			"[%d] %s(%s) %s/%s",
			index,
			target.CourseName,
			target.SemesterName,
			target.Credit,
			target.CreditHour,
		)
	}

	utils.Info("%s", "[*] 输入选课目标编号")
	var i int
	fmt.Scanln(&i)

	id := (*targets)[i].CourseId
	date, err := time.Parse("2006-01-02 15:04:05", (*targets)[i].CStartDate)
	utils.CheckIfError(err)
	toki := date.Add(time.Second * config.AHEAD_OF_TIME)

	utils.Info("[*] 即将在 %s 开始抢课\n请等待 %s", toki, toki.Sub(time.Now()))
	time.Sleep(toki.Sub(time.Now()))

	for !time.Now().Before(toki) {
		startTime := time.Now()
		err := c.SelectCourse(&course.Course{
			CourseId: id,
		})
		end := time.Now()
		if err != nil {
			fmt.Printf("[!] 选课失败: %s, 耗时: %s, now: %s\n", err, end.Sub(startTime), end)
			time.Sleep(config.TIME_INTERVAL * time.Millisecond)
		} else {
			utils.Info("[*] 选课成功")
			break
		}
	}

}
