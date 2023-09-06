// main.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// SelectCourse 封装课程选课功能
// SelectCourse 封装课程选课功能
func SelectCourse(sessionID string, courseName string, courseNumber string, classroomNumber string, People string, Contain string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		startTime := time.Now() // 记录选课开始时间

		data := url.Values{
			"kcbh":   {courseNumber},
			"kczxf":  {"2.0"},
			"ktbh":   {"con"},
			"ktrl":   {People},
			"ktrs":   {Contain},
			"markZB": {},
			"kcmc":   {courseName},
		}

		client := &http.Client{}

		// 创建 HTTP 请求
		req, err := http.NewRequest("POST", "http://wsxk.hust.edu.cn/zxqstudentcourse/zxqcoursesresult.action", strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Printf("Error creating request for %s: %v\n", courseName, err)
			return
		}

		// 设置请求头
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/117.0")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
		req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Origin", "http://wsxk.hust.edu.cn")
		req.Header.Set("Connection", "close")
		referer := fmt.Sprintf("http://wsxk.hust.edu.cn/zxqstudentcourse/zxqclassroom.action?kcbh=%s&ggkdl=&markZB=", courseNumber, classroomNumber)
		req.Header.Set("Referer", referer)
		req.Header.Set("Cookie", sessionID) // 请确保 SESSIONID 是已定义的变量
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("DNT", "1")
		req.Header.Set("Sec-GPC", "1")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request for %s: %v\n", courseName, err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response for %s: %v\n", courseName, err)
			return
		}

		// 输出选课结果和选课耗时
		fmt.Printf("Time: %s, Course: %s\n", startTime.Format("2006-01-02 15:04:05"), courseName)

		if strings.Contains(string(body), "选课失败") {
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
	sessionID := "JSESSIONID=5udpAN-VQcLmPsYS4i5NEJMm32Q4AO9O_kLQZkQn5ec_p9TXxGU0!-232618489; BIGipServerpool-hub-wsxkxt=2970225162.22811.0000" // 请替换为你的Cookie
	// 创建一个等待组，用于等待所有goroutines完成
	var wg sync.WaitGroup

	// 从courses.go文件导入课程信息
	courses := Courses

	// 启动并发选课
	for _, course := range courses {
		wg.Add(1)
		go SelectCourse(sessionID, course.CourseName, course.CourseNumber, course.ClassroomNumber, course.People, course.Contain, &wg)
	}

	// 等待所有goroutines完成
	wg.Wait()
}
