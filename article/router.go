package article

import (
	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
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
		Author:  art.NickName,
		Title:   art.Title,
		Type:    art.Type,
		Content: art.Content,
	}
	if err = CreateArticle(article); err != nil {
		global.LOGGER.Error("add article", zap.Error(err))
		response.JSONResponse(c, "failed", nil)
	}
	response.JSONResponse(c, "success", nil)

}

// /article/list
func ArticleList(c *gin.Context) {
	articles, err := QueryAllArticle()
	if err != nil {
		global.LOGGER.Error("get article list", zap.Error(err))
		response.JSONResponse(c, "failed", nil)
	}
	response.HTMLResponse(c, "blog_list.html", gin.H{
		"articles": articles,
	})

}

// /article/:id
func ArticleDetail(c *gin.Context) {
	id := c.Param("id")
	// id, err := strconv.Atoi(sid)
	// if err != nil {
	// 	global.LOGGER.Error("atoi", zap.Error(err))
	// 	response.JSONResponse(c, "failed", nil)
	// }
	article, err := GetArticleByID(id)
	if err != nil {
		global.LOGGER.Error("get article by id", zap.Error(err))
		response.JSONResponse(c, "failed", nil)
	}
	// updatedAt := article.UpdatedAt.Format("2024-01-01 15:05")
	response.HTMLResponse(c, "blog_detail.html", gin.H{
		"article": article,
		// "updateAt": updatedAt,
	})
}
