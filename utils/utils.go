package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"

	"github.com/cyansobble/global"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/jordan-wright/email"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var Store = base64Captcha.DefaultMemStore

func GetRandonPicture(num string) []byte {
	client := resty.New()
	// Referer:https://weibo.com/
	//resp, err := client.R().EnableTrace().Get("https://sese.iw233.top/iapi.php?sort=cdnrandom")
	resp, err := client.R().SetQueryParams(map[string]string{
		"sort": "random",
		"type": "json",
		"num":  num,
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

func IsLogin(c *gin.Context) (string, bool) {
	token, _ := GetToken(c)
	if token == "" {
		return "", false
	}
	claims, err := ParseToken(token)
	if err != nil {
		return "", false
	}

	return claims.NickName, true
}

func SendEmail(to, subject, text string) (err error) {
	host := global.CONFIG.Email.Host
	userName := global.CONFIG.Email.UserName
	auth := smtp.PlainAuth("", userName, global.CONFIG.Email.PassWord, host)
	addr := fmt.Sprintf("%s:%s", host, global.CONFIG.Email.Port)
	e := email.Email{
		To:      []string{to},
		From:    userName,
		Subject: subject,
		Text:    []byte(text),
		//HTML: html,
	}
	err = e.Send(addr, auth)
	if err != nil {
		global.LOGGER.Error("send email", zap.Error(err))
	}
	return
}

func GenerateVerificationCode(n int) (s string) {
	// b := make([]byte, n)
	// for i := 0; i < n; i++ {
	// 	b[i] = byte(rand.Intn(10))
	// }
	// return b
	for i := 0; i < n; i++ {
		randomInt := rand.Intn(10)
		s += strconv.Itoa(randomInt)
	}
	return
}

// func VerifyCode() {
// 	db.GetStringInRedis(key)
// }

func GenerateDigitVerificationCode() (id, b64s, answer string, err error) {
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, Store)
	id, b64s, answer, err = captcha.Generate()
	if err != nil {
		global.LOGGER.Error("[GenerateDigitVerificationCode]", zap.Error(err))
	}
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		global.LOGGER.Error("[CheckPasswordHash]", zap.Error(err))
	}
	return err == nil
}
