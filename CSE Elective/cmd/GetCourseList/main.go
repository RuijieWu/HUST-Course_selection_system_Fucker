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

	c := client.NewFucker()

	t := config.TOKEN
	if t == "" {
		fmt.Println("输入 TOKEN:")
		fmt.Scanln(&t)
	}
	c.SetToken(t)

	targets, err := c.GetCourses()
	if err != nil {
		fmt.Println("获取课程列表失败: ", err)
		return
	}

	fmt.Println("课程列表:")
	for _, target := range *targets {
		fmt.Printf("%+v\n", target)
	}
	data, err := yaml.Marshal(&targets)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(
		config.COURSE_LIST,
		data,
		0666,
	)
}
