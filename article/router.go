package article

import (
	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/cyansobble/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// /article/add
func AddArticle(c *gin.Context) {
	var art ReqArticle
	err := c.ShouldBindJSON(&art)
	if err != nil {
		global.LOGGER.Error("add blog shouldbindjson", zap.Error(err))
		response.JSONResponse(c, "add blog failed", nil)
		return
	}
	article := Article{
		Author:  user.Author{NickName: art.NickName},
		Title:   art.Title,
		Type:    art.Type,
		Content: art.Content,
	}
	if err = CreateArticle(global.DB, article); err != nil {
		global.LOGGER.Error("add article", zap.Error(err))
		response.JSONResponse(c, "failed", nil)
	}
	response.JSONResponse(c, "success", nil)

}

// /article/list
func GetArticleList(c *gin.Context) {
	articles, err := QueryAllArticle(global.DB)
	if err != nil {
		global.LOGGER.Error("get article list", zap.Error(err))
		response.JSONResponse(c, "failed", nil)
	}
	response.HTMLResponse(c, "", gin.H{
		"articles": articles,
	})

}
