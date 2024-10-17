package client

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

type Client struct {
	client *req.Client
	token  string
}

func NewClient() *Client {
	return &Client{
		client: req.C().ImpersonateChrome().SetTimeout(10 * time.Second),
	}
}

func (c *Client) GetCapchaImage() ([]byte, string, error) {
	url := "http://222.20.126.201/dev-api/captchaImage"
	resp, err := c.client.R().Get(url)
	if err != nil || resp.GetStatusCode() != 200 {
		logrus.Errorf("[GetCapchaImage] failed: code=%d, msg=%s, err=%v", resp.GetStatusCode(), resp.String(), err)
		return nil, "", err
	}
	type Response struct {
		Code           int    `json:"code"`
		Msg            string `json:"msg"`
		Img            string `json:"img"`
		Uuid           string `json:"uuid"`
		CaptchaEnabled bool   `json:"captchaEnabled"`
	}

	res := &Response{}
	json.Unmarshal(resp.Bytes(), res)
	if res.Code != 200 {
		logrus.Errorf("get capcha image failed: %s", res.Msg)
		return nil, "", err
	}

	var img []byte
	img, err = base64.StdEncoding.DecodeString(res.Img)
	if err != nil {
		logrus.Errorf("decode capcha image failed: %s", err)
		return nil, "", err
	}
	return img, res.Uuid, nil
}

func (c *Client) Login(username string, password string, code string, uuid string) error {
	url := "http://222.20.126.201/dev-api/login"
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Code     string `json:"code"`
		Uuid     string `json:"uuid"`
	}

	type Response struct {
		Code  int    `json:"code"`
		Token string `json:"token"`
		Msg   string `json:"msg"`
	}

	resp := &Response{}
	res, err := c.client.R().SetBody(Request{
		Username: username,
		Password: password,
		Code:     code,
		Uuid:     uuid,
	}).Post(url)
	if err != nil || res.GetStatusCode() != 200 {
		logrus.Errorf("[Login] failed: code=%d, msg=%s, err=%v", res.GetStatusCode(), res.String(), err)
		return err
	}

	json.Unmarshal(res.Bytes(), resp)
	if resp.Code != 200 {
		logrus.Errorf("login failed: %s", resp.Msg)
		return err
	}
	c.token = resp.Token
	c.client.SetCommonBearerAuthToken(c.token)
	return nil
}

func (c *Client) SetToken(token string) {
	c.token = token
	c.client.SetCommonBearerAuthToken(token)
}
