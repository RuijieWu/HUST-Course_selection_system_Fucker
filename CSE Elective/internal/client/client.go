package client

import (
	"errors"
	"time"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/course"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/user"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/utils"
	"github.com/imroc/req/v3"
)

var (
	ErrGetTimeDiffFailed = errors.New("[*] Get Time difference Failed ")
)

type Fucker struct {
	Client *req.Client
	Token  string
}

func NewFucker() *Fucker {
	return &Fucker{
		Client: req.C().ImpersonateChrome().SetTimeout(10 * time.Second),
	}
}

type Profile struct {
	Content string
}

func (c *Fucker) GetCapchaImage() ([]byte, string, error) {
	return user.GetCapchaImage(c.Client)
}

func (c *Fucker) Login(username string, password string, code string, uuid string) error {
	t, err := user.Login(c.Client, username, password, code, uuid)
	c.SetToken(t)
	return err
}

func (c *Fucker) SetToken(token string) {
	c.Token = token
	c.Client.SetCommonBearerAuthToken(token)
}

func (c *Fucker) GetProfile() (Profile, error) {
	p, err := user.GetProfile(c.Client)
	return parseProfile(p), err
}

func parseProfile(raw []byte) Profile {
	//	err = json.Unmarshal(resp.Bytes(), res)
	return Profile{Content: string(raw)}
}

func (c *Fucker) GetCourses() (*[]course.Course, error) {
	return course.GetCourses(c.Client)
}

func (c *Fucker) SelectCourse(target *course.Course) error {
	return course.SelectCourse(c.Client, target)
}

func (c *Fucker) GetTimeDiff() (time.Duration, error) {
	//* 往返过程中大概率是 Client 发送时间更接近 Server 收到请求的时间
	c_date := time.Now()
	resp, err := c.Client.R().Get("http://222.20.126.201/student/student/course")
	//* fmt.Println(c_date)
	//* fmt.Println(time.Now())
	utils.CheckIfError(err)
	//* fmt.Println(resp.TotalTime())
	date_str, ok := resp.Header["Date"]
	if !ok {
		return -1, ErrGetTimeDiffFailed
	}
	s_date, err := time.Parse(time.RFC1123, date_str[0])
	utils.CheckIfError(err)
	s_date = s_date.Local()
	//* utils.Info("Client Time:%s\nServer Time:%s", c_date, s_date)
	//* 按照之前抢课的经验服务器时间会更快
	return c_date.Sub(s_date), nil
}
