package client

import (
	"os"
	"testing"
)

func TestGetCaptchaImg(t *testing.T) {
	c := NewClient()
	img, _, err := c.GetCapchaImage()
	if err != nil {
		t.Errorf("GetCapchaImage() failed: %v", err)
		return
	}

	t.Logf("GetCapchaImage() success: img size=%d", len(img))

	os.WriteFile("captcha.jpg", img, 0644)

}
