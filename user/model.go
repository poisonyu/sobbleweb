package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type Author struct {
// 	NickName string `json:"nickname" gorm:"column:nickname"`
// }

type User struct {
	gorm.Model
	UUID        uuid.UUID
	UserName    string `json:"username" gorm:"column:username"`
	PassWord    string `json:"-" gorm:"column:password"`
	NickName    string `json:"nickname" gorm:"column:nickname"`
	Email       string `json:"email" gorm:"column:email"`
	AuthorityId int    `json:"authorityid" gorm:"default:6;column:authorityid"`
	Phone       string `json:"phone"`
	Enable      int    `json:"enable" gorm:"default:1"`
	HeaderImg   string `json:"headerimg" gorm:"default:;column:headerimg"`
}

type ReqRegister struct {
	UserName       string `json:"username"`
	PassWord       string `json:"password"`
	PassWordRepeat string `json:"passwordrepeat"`
	NickName       string `json:"nickname"`
	Email          string `json:"email"`
	CaptchaId      string `json:"captchaid"`
	Captcha        string `json:"captcha"`
}

type ReqLogin struct {
	UserName  string `json:"username"`
	PassWord  string `json:"password"`
	Email     string `json:"email"`
	CaptchaId string `json:"captchaid"`
	Captcha   string `json:"captcha"`
}

type ReqChangePassword struct {
	PassWord    string `json:"password"`
	NewPassWord string `json:"newpassword"`
	// CaptchaId   string `json:"captchaid"`
	// Captcha     string `json:"captcha"`
	VerificationCode string `json:"verificationcode"`
}
