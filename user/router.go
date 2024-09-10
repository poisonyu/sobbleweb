package user

import (
	"errors"
	"fmt"
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

	if !utils.Store.Verify(reg.CaptchaId, reg.Captcha, true) {
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

	if !utils.Store.Verify(l.CaptchaId, l.Captcha, true) {
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

// var store = base64Captcha.DefaultMemStore

func DigitCaptcha(c *gin.Context) {

	// captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	// id, b64s, _, err := captcha.Generate()
	// if err != nil {
	// 	global.LOGGER.Error("captcha generate failed", zap.Error(err))
	// 	response.JSONResponse(c, 0, "captcha generate failed", nil)
	// }
	// global.LOGGER.Info("Captcah generate success", zap.String("id", id))
	// response.JSONResponse(c, 1, "success", gin.H{
	// 	"id":   id,
	// 	"b64s": b64s,
	// })
	id, b64s, _, err := utils.GenerateDigitVerificationCode()
	if err != nil {
		response.JSONResponse(c, 0, "captcha generate failed", nil)
		return
	}
	response.JSONResponse(c, 1, "success", gin.H{
		"id":   id,
		"b64s": b64s,
	})
}

// todo
func AudioCaptcha(c *gin.Context) {
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverAudio, utils.Store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.LOGGER.Error("captcha generate failed", zap.Error(err))
		response.JSONResponse(c, 0, "captcha generate failed", nil)
	}
	global.LOGGER.Info("Captcah generate success", zap.String("id", id))
	response.JSONResponse(c, 1, "success", b64s)
}

// /user/signin
func LoginHtml(c *gin.Context) {
	redirect := c.Request.Header.Get("Referer")
	response.HTMLResponse(c, "login.html", gin.H{
		"redirect":    redirect,
		"loginactive": "active",
	})

}

// /user/signup
func RegisterHtml(c *gin.Context) {
	response.HTMLResponse(c, "register.html", gin.H{
		"registeractive": "active",
	})
}

// /user/info/:id

func UserInfo(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		global.LOGGER.Info("context no claims")
		response.JSONResponse(c, 0, "请登录", nil)
		return
	}
	id := claims.(*utils.CustomClaim).UserID
	user, err := GetUserByID(id)
	if err != nil {
		global.LOGGER.Error("get user by id ", zap.Error(err))
		response.JSONResponse(c, 0, "用户不存在", nil)
		return
	}
	response.HTMLResponse(c, "info.html", gin.H{
		"id":        user.ID,
		"username":  user.UserName,
		"nickname":  user.NickName,
		"email":     user.Email,
		"phone":     user.Phone,
		"headerimg": user.HeaderImg,
		"isLogin":   utils.IsLogin(c),
	})

}

func UserEditInfo(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		global.LOGGER.Error("login shouldbindjson error", zap.Error(err))
		response.JSONResponse(c, 0, "failed", nil)
		return
	}
	fmt.Println("id", u.ID)
	user, err := GetUserByID(u.ID)
	if err != nil {
		global.LOGGER.Error("get user by id", zap.Error(err))
		response.JSONResponse(c, 0, "failed", nil)
		return
	}
	user.NickName = u.NickName
	user.Email = u.Email
	user.Phone = u.Phone
	err = SaveUser(user)
	if err != nil {
		global.LOGGER.Error("save user ", zap.Error(err))
		response.JSONResponse(c, 0, "保存失败", nil)
		return
	}
	response.JSONResponse(c, 1, "保存成功", nil)
}

func ChangePassword(c *gin.Context) {
	var p ReqChangePassword
	err := c.ShouldBindJSON(&p)
	if err != nil {
		global.LOGGER.Error("register shouldbindjson error", zap.Error(err))
		response.JSONResponse(c, 0, "register failed", nil)
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		response.JSONResponse(c, 0, "请登录", nil)
		return
	}

	userid := claims.(*utils.CustomClaim).UserID
	user, ok := isUserExist(userid)
	if !ok {
		response.JSONResponse(c, 0, "user not foud", nil)
		return
	}
	key := fmt.Sprintf("sobbleweb_verification_%d", int(userid))
	val, err := GetStringInRedis(key)
	if err != nil {
		response.JSONResponse(c, 0, "验证码获取失败", nil)
		return
	}
	if p.VerificationCode != val {
		response.JSONResponse(c, 0, "验证码错误", nil)
		return
	}

	// passWord, _ := bcrypt.GenerateFromPassword([]byte(p.PassWord), 7)
	// if user.PassWord != string(passWord) {
	// 	response.JSONResponse(c, 0, "原密码错误", nil)
	// 	return
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(p.PassWord))
	if err != nil {
		response.JSONResponse(c, 0, "原密码错误", nil)
		return
	}

	newPassWord, _ := bcrypt.GenerateFromPassword([]byte(p.NewPassWord), 7)
	user.PassWord = string(newPassWord)
	err = SaveUser(user)
	if err != nil {
		response.JSONResponse(c, 0, "保存失败", nil)
		return
	}
	response.JSONResponse(c, 1, "修改成功", nil)
}

func Verification(c *gin.Context) {
	// var u User
	// err := c.ShouldBindJSON(&u)
	// if err != nil {
	// 	global.LOGGER.Error("verification", zap.Error(err))
	// 	response.JSONResponse(c, 0, "failed", nil)
	// }
	claims, ok := c.Get("claims")
	if !ok {
		response.JSONResponse(c, 0, "请登录", nil)
		return
	}

	userid := claims.(*utils.CustomClaim).UserID
	user, ok := isUserExist(userid)
	if !ok {
		response.JSONResponse(c, 0, "user not foud", nil)
		return
	}
	to := user.Email
	// 生成验证码
	// id, _, verificationCode, err := utils.GenerateDigitVerificationCode()
	// if err != nil {
	// 	response.JSONResponse(c, 0, "生成验证码失败", nil)
	// 	return
	// }
	verificationCode := utils.GenerateVerificationCode(6)

	key := fmt.Sprintf("sobbleweb_verification_%d", int(userid))
	expiration := 30 * time.Minute
	val, _ := SetStringInRedis(key, verificationCode, expiration)
	if val != "OK" {
		response.JSONResponse(c, 0, "发送验证码失败", nil)
		return
	}

	text := fmt.Sprintf("Your email verification code is\n %s \n It will be expired after %s\nPlease protect your verification code.", verificationCode, expiration.String())
	err := utils.SendEmail(to, "Verification Code", text)
	if err != nil {
		response.JSONResponse(c, 0, "发送验证码失败", nil)
		return
	}
	response.JSONResponse(c, 1, "发送验证码成功", nil)

}

func isUserExist(id uint) (User, bool) {
	u, err := GetUserByID(id)
	if err != nil {
		global.LOGGER.Error("user don't exist", zap.Error(err))
		return u, false
	}
	return u, true
}

// func getClaims(c *gin.Context) {
// 	claims, ok := c.Get("claims")
// 	if !ok {
// 		response.JSONResponse(c, 0, "请登录", nil)
// 		return
// 	}
// 	claims.(utils.CustomClaim)
// }
