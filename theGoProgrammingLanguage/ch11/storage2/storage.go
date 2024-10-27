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
