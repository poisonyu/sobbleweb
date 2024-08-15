package article

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
		Author:      art.NickName,
		Title:       art.Title,
		Type:        art.Type,
		MdContent:   art.MdContent,
		HtmlContent: art.HtmlContent,
		// IsHTML:  art.IsHTML,
	}
	id, err := CreateArticle(article)
	if err != nil {
		global.LOGGER.Error("add article", zap.Error(err))
		response.JSONResponse(c, "failed", nil)
	}
	location := fmt.Sprintf("/article/%d", id)
	//response.JSONResponse(c, "success", nil)
	c.Redirect(http.StatusFound, location)
}

func DeleteArticle(c *gin.Context) {

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
		"id":      id,
		"article": article,
		// "updateAt": updatedAt,
	})
}

// /article/create
func EditNewArticle(c *gin.Context) {
	response.HTMLResponse(c, "create_article.html", nil)
}

// /edit/:id
func EditArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := GetArticleByID(id)
	// todo 待完善if 条件表达式
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOGGER.Info("recordNotFound", zap.Error(err))
		response.JSONResponse(c, "failed", gin.H{
			"reason": "recordnotfound",
		})
	}
	response.HTMLResponse(c, "edit_article.html", gin.H{
		"article": article,
	})
}

func UpdateArticle(c *gin.Context) {
	var art ReqArticle
	if err := c.ShouldBindJSON(&art); err != nil {
		global.LOGGER.Error("update shouldbindjson", zap.Error(err))
		response.JSONResponse(c, "update failed", nil)
		return
	}
	article, err := GetArticleByID(art.ID)
	// errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		global.LOGGER.Error("get article by id", zap.Error(err))
		response.JSONResponse(c, "update failed", nil)
	}
	article.MdContent = art.MdContent
	article.HtmlContent = art.HtmlContent
	if err := SaveArticle(article); err != nil {
		global.LOGGER.Error("update article", zap.Error(err))
		response.JSONResponse(c, "update failed", nil)
	}
	location := fmt.Sprintf("/article/%d", article.ID)
	//response.JSONResponse(c, "success", nil)
	c.Redirect(http.StatusFound, location)

}
