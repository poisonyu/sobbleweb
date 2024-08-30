package upload

import (
	"os"
	"path/filepath"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 上传文件
func UploadFile(c *gin.Context) {
	fileHead, err := c.FormFile("file")
	if err != nil {
		global.LOGGER.Error("formfile", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
	}
	fileOwner := c.PostForm("owner")
	if fileOwner == "" {
		global.LOGGER.Error("need key 'owner' in form data", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
	}
	filePath, err := SaveUploadedFile(fileHead)
	if err != nil {
		global.LOGGER.Error("save uploaded file error", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
	}
	uploadFile := FileInfo{
		FileName: fileHead.Filename,
		FilePath: filePath,
		Tag:      filepath.Ext(fileHead.Filename),
		Owner:    fileOwner,
	}
	if err := CreateFileInfo(&uploadFile); err != nil {
		global.LOGGER.Error("create fileInfo failed", zap.Error(err))
		response.JSONResponse(c, 0, "上传文件失败", nil)
	}

	response.JSONResponse(c, 1, "上传文件成功", uploadFile)

}

// 获取文件列表
func GetFileList(c *gin.Context) {
	var pageInfo PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		global.LOGGER.Error("should bind json pagainfo", zap.Error(err))
		response.JSONResponse(c, 0, "获取文件列表失败", nil)
	}
	files, total, err := QueryFileList(pageInfo)
	if err != nil {
		global.LOGGER.Error("query file list", zap.Error(err))
		response.JSONResponse(c, 0, "获取文件列表失败", nil)
	}
	response.JSONResponse(c, 1, "成功获取文件列表", gin.H{
		"file":     files,
		"total":    total,
		"page":     pageInfo.Page,
		"pagesize": pageInfo.PageSize,
	})

}

// 删除文件
func DeleteFile(c *gin.Context) {
	var fileInfo FileInfo
	if err := c.ShouldBindJSON(&fileInfo); err != nil {
		global.LOGGER.Error("should bind json FileInfo", zap.Error(err))
		response.JSONResponse(c, 0, "删除文件失败", nil)
	}
	newFileInfo, err := GetFileByID(fileInfo.ID)
	if err != nil {
		global.LOGGER.Error("get file by id", zap.Error(err))
		response.JSONResponse(c, 0, "删除文件失败", nil)
	}
	if err := os.Remove(newFileInfo.FileName); err != nil {
		global.LOGGER.Error("remove file", zap.Error(err))
		response.JSONResponse(c, 0, "删除文件失败", nil)
	}
	if err := DeleteFileByID(newFileInfo.ID); err != nil {
		global.LOGGER.Error("delete file in db", zap.Error(err))
		response.JSONResponse(c, 0, "删除文件失败", nil)
	}
	response.JSONResponse(c, 1, "删除文件成功", nil)

}

// 批量删除文件
