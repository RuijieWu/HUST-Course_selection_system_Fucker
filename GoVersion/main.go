// main.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// SelectCourse 封装课程选课功能
func SelectCourse(sessionID string, courseName string, courseNumber string, classroomNumber string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		startTime := time.Now() // 记录选课开始时间

		data := url.Values{
			"kcbh":   {courseNumber},
			"kczxf":  {"2"},
			"ktbh":   {classroomNumber},
			"ktrl":   {"100"}, // 请根据实际情况调整课堂容量和人数
			"ktrs":   {"100"},
			"markZB": {"classroomNumber"},
			"kcmc":   {courseName},
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://wsxk.hust.edu.cn/zxqstudentcourse/zxqcoursesresult.action", strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Printf("Error creating request for %s: %v\n", courseName, err)
			return
		}

		req.Header.Add("Connection", "keep-alive")
		req.Header.Add("Cache-Control", "max-age=1")
		req.Header.Add("Origin", "http://wsxk.hust.edu.cn")
		req.Header.Add("Upgrade-Insecure-Requests", "1")
		req.Header.Add("DNT", "1")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		req.Header.Add("Referer", "http://wsxk.hust.edu.cn/zxqstudentcourse/coursesandclassroom.action?markZB=6&ggkdl=&GGKDLBH=0&skZC=&skJC=&kcmc=%E6%98%93%E7%BB%8F") // 请替换为你的课程URL
		req.Header.Add("Accept-Encoding", "gzip, deflate")
		req.Header.Add("Accept-Language", "en-US,en;q=0.9")

		// 添加Cookie头部
		req.Header.Add("Cookie", sessionID)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for %s: %v\n", courseName, err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response for %s: %v\n", courseName, err)
			return
		}

		// 输出选课结果和选课耗时
		fmt.Printf("Time: %s, Course: %s\n", startTime.Format("2006-01-02 15:04:05"), courseName)

		if strings.Contains(string(body), "失败") {
			fmt.Printf("Failed to select %s. Retrying...\n", courseName)
			// 等待一段时间再次尝试选课，可以自行调整等待时间
			time.Sleep(time.Second * 10)
		} else {
			fmt.Printf("Successfully selected %s.\n", courseName)
			break // 成功选课后退出循环
		}
	}
}

func main() {
	// 创建一个等待组，用于等待所有goroutines完成
	var wg sync.WaitGroup
	// 启动并发选课
	for true {
		wg.Add(1)
		go SelectCourse(sessionID, course.CourseName, course.CourseNumber, course.ClassroomNumber, &wg)
	}
	// 等待所有goroutines完成
	wg.Wait()
}
