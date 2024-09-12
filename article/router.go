package article

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/cyansobble/upload"
	"github.com/cyansobble/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// /article/add
func AddArticle(c *gin.Context) {
	var art ReqArticle
	var author string
	err := c.ShouldBindJSON(&art)
	if err != nil {
		global.LOGGER.Error("add blog shouldbindjson", zap.Error(err))
		response.JSONResponse(c, 0, "add blog failed", nil)
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		author = "default"
	}
	author = claims.(*utils.CustomClaim).NickName
	article := Article{
		Author:      author,
		Title:       art.Title,
		Type:        art.Type,
		MdContent:   art.MdContent,
		HtmlContent: art.HtmlContent,
		// IsHTML:  art.IsHTML,
	}
	id, err := CreateArticle(article)
	if err != nil {
		global.LOGGER.Error("add article", zap.Error(err))
		response.JSONResponse(c, 0, "failed", nil)
		return
	}
	location := fmt.Sprintf("/article/%d", id)
	response.JSONResponse(c, 1, "success", gin.H{
		"redirect": location,
	})
	//c.Redirect(http.StatusFound, location)
}

// /article/delete

func DeleteArticle(c *gin.Context) {
	var art ReqArticle
	if err := c.ShouldBindJSON(&art); err != nil {
		global.LOGGER.Error("delete shouldbindjson", zap.Error(err))
		response.JSONResponse(c, 0, "delete failed", nil)
		return
	}
	_, err := GetArticleByID(art.ID)
	if err != nil {
		global.LOGGER.Error("get article by id", zap.Error(err))
		response.JSONResponse(c, 0, "delete failed", nil)
		return
	}
	if err := DeleteArticleByID(art.ID); err != nil {
		global.LOGGER.Error("delete article by id", zap.Error(err))
		response.JSONResponse(c, 0, "delete failed", nil)
		return
	}
	response.JSONResponse(c, 1, "success", gin.H{
		"redirect": "/article/list",
	})
	//c.Redirect(http.StatusFound, "/article/list")
}

// /article/list
func ArticleList(c *gin.Context) {
	articles, err := QueryAllArticle()
	if err != nil {
		global.LOGGER.Error("get article list", zap.Error(err))
		response.JSONResponse(c, 0, "failed", nil)
		return
	}
	dates, err := GetArticleDescDate()
	if err != nil {
		global.LOGGER.Error("get article desc date", zap.Error(err))

	}
	//archiveDate := Archives(dates)

	response.HTMLResponse(c, "blog_list.html", gin.H{
		"articles":    articles,
		"isLogin":     utils.IsLogin(c),
		"archiveDate": Archives(dates),
		"listactive":  "active",
	})
}

// func ArticleDateList(c *gin.Context) {

// }

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
		response.JSONResponse(c, 0, "failed", nil)
		return
	}
	var previous, next string
	preArticle, err := Previous(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOGGER.Sugar().Infof("ID:%s 没有上一篇了", id)
	} else if err != nil {
		global.LOGGER.Error("previous", zap.Error(err))
	} else {
		previous = strconv.Itoa(int(preArticle.ID))
	}

	nextArticle, err := Next(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOGGER.Sugar().Infof("ID:%s 没有下一篇了", id)
	} else if err != nil {
		global.LOGGER.Error("previous", zap.Error(err))
	} else {
		next = strconv.Itoa(int(nextArticle.ID))
	}
	// updatedAt := article.UpdatedAt.Format("2024-01-01 15:05")
	response.HTMLResponse(c, "blog_detail.html", gin.H{
		"id":       id,
		"article":  article,
		"previous": previous,
		"next":     next,
		// "updateAt": updatedAt,
		"isLogin": utils.IsLogin(c),
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
		response.JSONResponse(c, 0, "failed", gin.H{
			"reason": "recordnotfound",
		})
		return
	}
	response.HTMLResponse(c, "edit_article.html", gin.H{
		"article": article,
	})
}

// /article/update
func UpdateArticle(c *gin.Context) {
	fileHead, err := c.FormFile("cover")
	if err != nil {
		global.LOGGER.Error("formfile", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
		return
	}
	articleID := c.PostForm("id")
	title := c.PostForm("title")
	mdContent := c.PostForm("mdcontent")
	htmlContent := c.PostForm("htmlcontent")
	if articleID == "" || mdContent == "" || htmlContent == "" {
		global.LOGGER.Info("缺少formdata")
		response.JSONResponse(c, 0, "上传文件失败", nil)
		return
	}
	a, ok := isArticleExist(articleID)
	if !ok {
		response.JSONResponse(c, 0, "没有找到对应id的文件", nil)
		return
	}

	filePath, err := upload.SaveUploadedFile(fileHead)
	if err != nil {
		global.LOGGER.Error("save uploaded file error", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
		return
	}
	uploadFile := upload.FileInfo{
		FileName: fileHead.Filename,
		FilePath: filePath,
		Tag:      filepath.Ext(fileHead.Filename),
		Owner:    a.Author,
	}
	if err := upload.CreateFileInfo(&uploadFile); err != nil {
		global.LOGGER.Error("create fileInfo failed", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
		return
	}

	if title != "" {
		a.Title = title
	} else {
		title = a.Title
	}
	a.MdContent = mdContent
	a.HtmlContent = htmlContent
	a.Cover = filePath
	if err := SaveArticle(a); err != nil {
		global.LOGGER.Error("update article", zap.Error(err))
		response.JSONResponse(c, 0, "update failed", nil)
		return
	}
	location := fmt.Sprintf("/article/%d", a.ID)
	response.JSONResponse(c, 1, "success", gin.H{
		"title":    title,
		"redirect": location,
	})

}

// var art ReqArticle
// var title string
// if err := c.ShouldBindJSON(&art); err != nil {
// 	global.LOGGER.Error("update shouldbindjson", zap.Error(err))
// 	response.JSONResponse(c, 0, "update failed", nil)
// 	return
// }

// article, err := GetArticleByID(art.ID)
// errors.Is(err, gorm.ErrRecordNotFound)
// if err != nil {
// 	global.LOGGER.Error("get article by id", zap.Error(err))
// 	response.JSONResponse(c, 0, "update failed", nil)
// 	return
// }
// if art.Title != "" {
// 	title = art.Title
// 	article.Title = art.Title
// } else {
// 	title = article.Title
// }
// article.MdContent = art.MdContent
// article.HtmlContent = art.HtmlContent
// if err := SaveArticle(article); err != nil {
// 	global.LOGGER.Error("update article", zap.Error(err))
// 	response.JSONResponse(c, 0, "update failed", nil)
// 	return
// }
// location := fmt.Sprintf("/article/%d", article.ID)
// response.JSONResponse(c, 1, "success", gin.H{
// 	"title":    title,
// 	"redirect": location,
// })
// c.Redirect(http.StatusFound, location)
// response.RedirectResponse(c, location)
