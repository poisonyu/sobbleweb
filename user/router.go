package user

import (
	"errors"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var reg ReqRegister
	err := c.ShouldBindJSON(&reg)
	if err != nil {
		global.LOGGER.Error("register shouldbindjson error", zap.Error(err))
		response.JSONResponse(c, "register failed", nil)
		return
	}

	// TODO 判断必要的字段满足特定的要求
	//var user User
	err = global.DB.Model(&User{}).Where("username = ?", reg.UserName).First(&User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.JSONResponse(c, "用户名已经存在", nil)
		return
	}

	if !store.Verify(reg.CaptchaId, reg.Captcha, true) {
		response.JSONResponse(c, "验证码错误", nil)
		return
	}
	passWord, _ := bcrypt.GenerateFromPassword([]byte(reg.PassWord), 7)
	user := User{
		UUID:     uuid.New(),
		UserName: reg.UserName,
		Author:   Author{NickName: reg.NickName},
		//NickName:    reg.NickName,
		PassWord:    string(passWord),
		Email:       reg.Email,
		AuthorityId: 6,
		Enable:      1,
	}

	if err = global.DB.Create(&user).Error; err != nil {
		global.LOGGER.Error("create user failed", zap.Error(err))
		response.JSONResponse(c, "服务器错误，注册失败", nil)
		return
	}
	response.JSONResponse(c, "注册成功", nil)

}

func Login(c *gin.Context) {
	var l ReqLogin
	if err := c.ShouldBindJSON(&l); err != nil {
		global.LOGGER.Error("login shouldbindjson error", zap.Error(err))
		response.JSONResponse(c, "login failed", nil)
		return
	}
	// TODO 判断必要的字段满足特定的要求

	if !store.Verify(l.CaptchaId, l.Captcha, true) {
		response.JSONResponse(c, "验证码错误", nil)
		return
	}

	var user User
	err := global.DB.Model(&User{}).Where("username = ?", l.UserName).First(&user).Error
	if err != nil {
		global.LOGGER.Error("用户不存在", zap.Error(err))
		response.JSONResponse(c, "用户名密码错误", nil)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(l.PassWord))
	if err != nil {
		global.LOGGER.Error("用户密码错误", zap.Error(err))
		response.JSONResponse(c, "用户名密码错误", nil)
		return
	}
	// 设置jwt token

	response.JSONResponse(c, "登录成功", nil)
}

var store = base64Captcha.DefaultMemStore

func DigitCaptcha(c *gin.Context) {

	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.LOGGER.Error("captcha generate failed", zap.Error(err))
		response.JSONResponse(c, "captcha generate failed", nil)
	}
	global.LOGGER.Info("Captcah generate success", zap.String("id", id))
	response.JSONResponse(c, "success", b64s)

}

func AudioCaptcha(c *gin.Context) {

	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverAudio, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.LOGGER.Error("captcha generate failed", zap.Error(err))
		response.JSONResponse(c, "captcha generate failed", nil)
	}
	global.LOGGER.Info("Captcah generate success", zap.String("id", id))
	response.JSONResponse(c, "success", b64s)

}
