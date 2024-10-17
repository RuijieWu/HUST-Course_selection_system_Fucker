package client

import (
	"time"

	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/course"
	"github.com/RuijieWu/HUST-OCSS-Fucker/CSE-Elective/internal/client/user"
	"github.com/imroc/req/v3"
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
