package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"

	"github.com/cyansobble/global"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
)

func GetRandonPicture() []byte {
	client := resty.New()
	// Referer:https://weibo.com/
	//resp, err := client.R().EnableTrace().Get("https://sese.iw233.top/iapi.php?sort=cdnrandom")
	resp, err := client.R().SetQueryParams(map[string]string{
		"sort": "random",
		"type": "json",
		"num":  "1",
	}).SetHeaders(map[string]string{
		"Referer": "https://weibo.com/",
		"Accept":  "application/json",
	}).Get("https://iw233.cn/api.php")
	if err != nil {
		global.LOGGER.Error("[request]:", zap.Error(err))
		return nil
	}
	global.LOGGER.Info(string(resp.Body()))
	return resp.Body()
	// if resp.StatusCode() == 200 {
	// 	return string(resp.Body())
	// }
	// return ""
}

func IsLogin(c *gin.Context) bool {
	token, _ := GetToken(c)
	return token != ""
}

func SendEmail(to, subject string, html []byte) (err error) {
	host := global.CONFIG.Email.Host
	userName := global.CONFIG.Email.UserName
	auth := smtp.PlainAuth("", userName, global.CONFIG.Email.PassWord, host)
	addr := fmt.Sprintf("%s:%s", host, global.CONFIG.Email.Port)
	e := email.Email{
		To:      []string{to},
		From:    userName,
		Subject: subject,
		//Text:    []byte("sunflower is a flower or a song."),
		HTML: html,
	}
	err = e.Send(addr, auth)
	if err != nil {
		global.LOGGER.Error("send email", zap.Error(err))
	}
	return
}

func GenerateVerificationCode(n int) []byte {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(10))
	}
	return b
}
