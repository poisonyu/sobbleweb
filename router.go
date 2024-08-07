package main

import (
	"encoding/json"
	"strings"

	"github.com/cyansobble/article"
	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/cyansobble/user"
	"github.com/cyansobble/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// curl -X POST http://localhost:8888/captcha
// 取变量名真头疼
func Router() {
	router := gin.Default()
	// gin.New()

	router.GET("/", Index)

	router.POST("/user/register", user.Register)
	router.POST("/captcha", user.DigitCaptcha)
	router.POST("/user/login", user.Login)
	//router.POST("/audiocaptcha", AudioCaptcha)

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/pic.html", "templates/index.html")
	router.GET("/random/pic", RandomPicture)

	router.POST("/article/add", article.AddArticle)
	router.GET("/article/list", article.ArticleList)
	router.GET("/article/:id", article.ArticleDetail)
	// todo jwt casbin
	router.Run(":8888")
}

// func test(c *gin.Context) {

//		c.JSON(200, map[string]string{
//			"name": "cyan",
//		})
//	}
func Index(c *gin.Context) {
	response.HTMLResponse(c, "index.html", gin.H{})
}

func RandomPicture(c *gin.Context) {
	b := utils.GetRandonPicture()
	data := map[string]interface{}{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		global.LOGGER.Error("[Unmarshal]:", zap.Error(err))
	}
	picUrl := data["pic"].([]interface{})[0].(string)
	slice := strings.Split(picUrl, "/")
	//https://setu.iw233.top/large/ec43126fgy1hl7qkktp7pj235s1s0qv7.jpg
	newUrl := "https://setu.iw233.top/large/" + slice[len(slice)-1]

	response.HTMLResponse(c, "pic.html", gin.H{
		"title": "RandomPicture",
		"url":   newUrl,
	})
	// Response(c, "success", gin.H{
	// 	"url": newUrl,
	// })
}
