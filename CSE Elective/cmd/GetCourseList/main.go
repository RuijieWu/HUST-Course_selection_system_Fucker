package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/config"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client"
	"gopkg.in/yaml.v2"
)

func main() {
	c := client.NewClient()

	t := config.TOKEN
	if t == "" {
		fmt.Println("输入TOKEN:")
		fmt.Scanln(&t)
	}
	c.SetToken(t)

	courses, err := c.GetCourses()
	if err != nil {
		fmt.Println("获取课程列表失败: ", err)
		return
	}

	fmt.Println("课程列表:")
	for _, course := range courses {
		fmt.Printf("%+v\n", course)
	}
	data, err := yaml.Marshal(&courses)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(
		config.CACHE_PATH+config.COURSE_LIST,
		data,
		0666,
	)
}
