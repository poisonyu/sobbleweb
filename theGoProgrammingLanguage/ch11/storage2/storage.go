package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

const sender = "purecard77@163.com"
const hostname = "smtp.163.com"

const template = `Warning: you are using %d bytes of storage, 
%d%% of your quota`

// 把发送邮件的逻辑移动到独立的函数中，并且存储到一个不可导出的包级变量notifyUser中
var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, "", hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.Sendmail(%s) failed: %s", username, err)
	}
}

func bytesInUse(username string) int64 {
	return 0
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
