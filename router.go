package main

import (
	"encoding/json"

	"github.com/cyansobble/article"
	"github.com/cyansobble/global"
	"github.com/cyansobble/middleware"
	"github.com/cyansobble/response"
	"github.com/cyansobble/upload"
	"github.com/cyansobble/user"
	"github.com/cyansobble/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router() {
	router := gin.Default()
	// gin.New()

	router.Static("/static", "./static")
	router.Static("/upload", "./upload")
	// router.Static("/dist", "./dist")
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/pic.html", "templates/index.html")
	router.Use(middleware.Cors())

	router.GET("/user/register", user.Register)
	router.POST("/user/register", user.Register)

	router.GET("/user/login", user.Login)
	router.POST("/user/login", user.Login)

	router.GET("/captcha", user.DigitCaptcha)

	// router.GET("/user/signin", user.LoginHtml)
	// router.GET("/user/signup", user.RegisterHtml)
	//router.POST("/audiocaptcha", AudioCaptcha)

	router.GET("/", Index)

	router.GET("/article/list", article.ArticleList)
	router.GET("/article/:id", article.ArticleDetail)

	router.GET("/random/pic", RandomPicture)
	router.GET("/multiplepic", GetPictures)
	router.GET("/video", ParseM3U8)
	router.GET("/bootstrap", BootStrap)

	router.Use(middleware.JWTAuth())

	router.GET("/user/info", user.UserInfo)
	router.POST("/user/editinfo", user.UserEditInfo)
	router.POST("/user/changepassword", user.ChangePassword)
	router.GET("/user/verification", user.Verification)

	router.Use(middleware.AuthorityAuth())

	router.POST("/article/create", article.CreateNewArticle)
	router.GET("/article/create", article.CreateNewArticle)

	router.POST("/article/update", article.UpdateArticle)
	router.GET("/edit/:id", article.UpdateArticle)

	router.POST("/article/delete", article.DeleteArticle)

	router.POST("/file/upload", upload.UploadFile)
	router.POST("/file/list", upload.GetFileList)
	router.POST("/file/delete", upload.DeleteFile)

	// todo jwt casbin
	router.Run(":8888")
}

func Index(c *gin.Context) {
	nickname, isLogin := utils.IsLogin(c)
	articles, count, err := article.QueryAllArticleDesc()
	if err != nil {
		global.LOGGER.Error("get article list", zap.Error(err))
		response.HTMLResponse(c, "index.html", gin.H{
			"isLogin":     isLogin,
			"indexactive": "active",
			"nickname":    nickname,
		})
		return
	}
	if count > 10 {
		articles = articles[:10]
	}
	// fmt.Println(articles[0], "cover:", articles[0].Cover)
	response.HTMLResponse(c, "index.html", gin.H{
		"articles":    articles,
		"isLogin":     isLogin,
		"indexactive": "active",
		"nickname":    nickname,
	})
}

func BootStrap(c *gin.Context) {
	response.HTMLResponse(c, "bootstrap.html", nil)
}

func RandomPicture(c *gin.Context) {
	b := utils.GetRandonPicture("1")
	data := map[string]interface{}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		global.LOGGER.Error("[Unmarshal]:", zap.Error(err))
	}
	picUrl := data["pic"].([]interface{})[0].(string)
	// slice := strings.Split(picUrl, "/")
	// //https://setu.iw233.top/large/ec43126fgy1hl7qkktp7pj235s1s0qv7.jpg
	// newUrl := "https://setu.iw233.top/large/" + slice[len(slice)-1]
	nickname, isLogin := utils.IsLogin(c)
	response.HTMLResponse(c, "pic.html", gin.H{
		"title": "RandomPicture",
		// "url":   newUrl,
		"url":      picUrl,
		"isLogin":  isLogin,
		"nickname": nickname,
	})

}

func GetPictures(c *gin.Context) {
	b := utils.GetRandonPicture("10")
	data := map[string]interface{}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		global.LOGGER.Error("[Unmarshal]:", zap.Error(err))
	}
	urls := data["pic"].([]interface{})
	nickname, isLogin := utils.IsLogin(c)
	response.HTMLResponse(c, "multiplepic.html", gin.H{
		"title":    "MultiplePicture",
		"urls":     urls,
		"isLogin":  isLogin,
		"nickname": nickname,
	})
}
func ParseM3U8(c *gin.Context) {
	response.HTMLResponse(c, "player.html", nil)
}

// func test(c *gin.Context) {

//		c.JSON(200, map[string]string{
//			"name": "cyan",
//		})
//	}
