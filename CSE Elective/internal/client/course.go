package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Course struct {
	CourseId          int    `json:"courseId"`
	CourseClassNumber string `json:"-"`
	CourseCode        string `json:"courseCode"`
	CourseName        string `json:"courseName"`
	SemesterName      string `json:"semesterName"`
	Major             string `json:"major"`
	Optional          int    `json:"optional"`
	Selected          int    `json:"selected"`
	CStartDate        string `json:"cStartDate"`
	CEndDate          string `json:"cEndDate"`
	Status            int    `json:"status"`
	Credit            string `json:"credit"`
	CreditHour        string `json:"creditHour"`
	Chosen            int    `json:"chosen"`
	Choosable         int    `json:"choosable"`
}

func (c *Client) GetCourses() ([]Course, error) {
	// activeSemester=true: 当前学期
	// chosen=false: 未选
	// choosable=true: 可选
	url := "http://222.20.126.201/dev-api/xuanke/course/student/?activeSemester=true&chosen=false&choosable=true&pageNum=1&pageSize=100"
	type Response struct {
		Code    int      `json:"code"`
		Msg     string   `json:"msg"`
		Total   int      `json:"total"`
		Courses []Course `json:"rows"`
	}

	resp := &Response{}
	res, err := c.client.R().Get(url)
	if err != nil || res.GetStatusCode() != 200 {
		logrus.Errorf("[GetCourses] failed: code=%d, msg=%s, err=%v", res.GetStatusCode(), res.String(), err)
		return nil, err
	}

	json.Unmarshal(res.Bytes(), resp)
	if resp.Code != 200 {
		logrus.Errorf("get courses failed: %s", resp.Msg)
		return nil, err
	}

	return resp.Courses, nil
}

func (c *Client) getCourseClassNumber(course *Course) error {
	// http://222.20.126.201/dev-api/xuanke/class/{course_id}/student

	url := fmt.Sprintf("http://222.20.126.201/dev-api/xuanke/class/%d/student", course.CourseId)

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Rows []struct {
			ClassId   string `json:"classNumber"`
			ClassName string `json:"className"`
		} `json:"rows"`
	}

	resp := &Response{}
	res, err := c.client.R().Get(url)
	if err != nil || res.GetStatusCode() != 200 {
		logrus.Errorf("[GetCourses] failed: code=%d, msg=%s, err=%v", res.GetStatusCode(), res.String(), err)
		return errors.New("get class number failed")
	}

	json.Unmarshal(res.Bytes(), resp)
	if resp.Code != 200 {
		logrus.Errorf("get class number failed: %s", resp.Msg)
		return errors.New("get class number failed")
	}

	if len(resp.Rows) == 0 {
		logrus.Error("没有可选择的课堂")
		return errors.New("class number not found")
	}
	course.CourseClassNumber = resp.Rows[0].ClassId
	return nil
}

var (
	ErrCourseLimit         = errors.New("选课人数已达上限！")
	ErrCourseTimeNotProper = errors.New("不在选课时段范围内！")
)

func (c *Client) SelectCourse(course Course) error {
	if course.CourseClassNumber == "" {
		// 没有classNumber，先获取
		err := c.getCourseClassNumber(&course)
		if err != nil {
			return err
		}
	}
	url := fmt.Sprintf("http://222.20.126.201/dev-api/xuanke/course/%d/select?classNumber=%s", course.CourseId, course.CourseClassNumber)

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}

	resp := &Response{}
	res, err := c.client.R().Put(url)
	if err != nil || res.GetStatusCode() != 200 {
		logrus.Errorf("[SelectCourse] failed: code=%d, msg=%s, err=%v", res.GetStatusCode(), res.String(), err)
		return errors.New("select course failed")
	}

	json.Unmarshal(res.Bytes(), resp)
	if resp.Code != 200 {
		switch resp.Msg {
		case "选课人数已达上限！":
			return ErrCourseLimit
		case "不在选课时段范围内！":
			return ErrCourseTimeNotProper
		default:
			return errors.New(resp.Msg)
		}
	}

	return nil
}
