package upload

import (
	"path/filepath"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFile(c *gin.Context) {
	fileHead, err := c.FormFile("file")
	if err != nil {
		global.LOGGER.Error("formfile", zap.Error(err))
		response.JSONResponse(c, "上传文件失败", nil)
	}
	fileOwner := c.PostForm("owner")
	if fileOwner == "" {
		global.LOGGER.Error("need key 'owner' in form data", zap.Error(err))
		response.JSONResponse(c, "上传文件失败", nil)
	}
	filePath, err := SaveUploadedFile(fileHead)
	if err != nil {
		global.LOGGER.Error("save uploaded file error", zap.Error(err))
		response.JSONResponse(c, "上传文件失败", nil)
	}
	uploadFile := FileInfo{
		FileName: fileHead.Filename,
		FilePath: filePath,
		Tag:      filepath.Ext(fileHead.Filename),
		Owner:    fileOwner,
	}
	if err := CreateFileInfo(&uploadFile); err != nil {
		global.LOGGER.Error("create fileInfo failed", zap.Error(err))
		response.JSONResponse(c, "上传文件失败", nil)
	}

	response.JSONResponse(c, "上传文件成功", uploadFile)

}
