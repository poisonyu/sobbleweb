package user

import (
	"errors"
	"time"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/cyansobble/utils"
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
		response.JSONResponse(c, 0, "register failed", nil)
		return
	}

	// TODO 判断必要的字段满足特定的要求
	//var user User
	if reg.PassWord != reg.PassWordRepeat {
		response.JSONResponse(c, 0, "密码不一致", nil)
		return
	}
	err = global.DB.Model(&User{}).Where("username = ?", reg.UserName).First(&User{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		response.JSONResponse(c, 0, "用户名已经存在", nil)
		return
	}

	if !store.Verify(reg.CaptchaId, reg.Captcha, true) {
		response.JSONResponse(c, 0, "验证码错误", nil)
		return
	}
	passWord, _ := bcrypt.GenerateFromPassword([]byte(reg.PassWord), 7)
	user := User{
		UUID:     uuid.New(),
		UserName: reg.UserName,
		NickName: reg.NickName,
		//NickName:    reg.NickName,
		PassWord:    string(passWord),
		Email:       reg.Email,
		AuthorityId: 6,
		Enable:      1,
	}

	if err = global.DB.Create(&user).Error; err != nil {
		global.LOGGER.Error("create user failed", zap.Error(err))
		response.JSONResponse(c, 0, "服务器错误", nil)
		return
	}
	response.JSONResponse(c, 1, "注册成功", nil)

}

func Login(c *gin.Context) {
	var l ReqLogin
	if err := c.ShouldBindJSON(&l); err != nil {
		global.LOGGER.Error("login shouldbindjson error", zap.Error(err))
		response.JSONResponse(c, 0, "login failed", nil)
		return
	}
	// TODO 判断必要的字段满足特定的要求

	if !store.Verify(l.CaptchaId, l.Captcha, true) {
		response.JSONResponse(c, 0, "验证码错误", nil)
		return
	}

	var user User
	err := global.DB.Model(&User{}).Where("username = ?", l.UserName).First(&user).Error
	if err != nil {
		global.LOGGER.Error("用户不存在", zap.Error(err))
		response.JSONResponse(c, 0, "用户名密码错误", nil)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(l.PassWord))
	if err != nil {
		global.LOGGER.Error("用户密码错误", zap.Error(err))
		response.JSONResponse(c, 0, "用户名密码错误", nil)
		return
	}
	// 设置jwt token
	userClaim := utils.UserClaim{
		Uuid:        user.UUID,
		UserID:      user.ID,
		UserName:    user.UserName,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	}
	claims := utils.CreateCustomClaim(userClaim)
	token, err := utils.CreateToken(claims)
	if err != nil {
		global.LOGGER.Error("create token failed", zap.Error(err))
		response.JSONResponse(c, 0, "登陆失败", nil)
		return
	}
	utils.SetToken(c, token, int(claims.ExpiresAt.Unix()-time.Now().Unix()))
	response.JSONResponse(c, 1, "登录成功", nil)
}

var store = base64Captcha.DefaultMemStore

func DigitCaptcha(c *gin.Context) {

	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.LOGGER.Error("captcha generate failed", zap.Error(err))
		response.JSONResponse(c, 0, "captcha generate failed", nil)
	}
	global.LOGGER.Info("Captcah generate success", zap.String("id", id))
	response.JSONResponse(c, 1, "success", gin.H{
		"id":   id,
		"b64s": b64s,
	})

}

func AudioCaptcha(c *gin.Context) {

	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverAudio, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.LOGGER.Error("captcha generate failed", zap.Error(err))
		response.JSONResponse(c, 0, "captcha generate failed", nil)
	}
	global.LOGGER.Info("Captcah generate success", zap.String("id", id))
	response.JSONResponse(c, 1, "success", b64s)

}

func LoginHtml(c *gin.Context) {
	response.HTMLResponse(c, "login.html", nil)

}

func RegisterHtml(c *gin.Context) {
	response.HTMLResponse(c, "register.html", nil)
}
