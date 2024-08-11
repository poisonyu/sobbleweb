package upload

import (
	"github.com/cyansobble/global"
)

func CreateFileInfo(f *FileInfo) error {
	return global.DB.Create(f).Error
}

func QueryFileList(pageInfo PageInfo) ([]FileInfo, int64, error) {
	var files []FileInfo
	limit := pageInfo.PageSize
	offset := limit * (pageInfo.Page - 1)
	keyWord := pageInfo.KeyWord
	db := global.DB.Model(&PageInfo{})
	if len(keyWord) > 0 {
		db = db.Where("filename like ?", "%"+keyWord+"%")
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return files, total, err
	}

	err := db.Limit(limit).Offset(offset).Find(&files).Error
	return files, total, err
}

func GetFileByID(id uint) (FileInfo, error) {
	var fileInfo FileInfo
	err := global.DB.First(&fileInfo, id).Error
	return fileInfo, err

}

func DeleteFileByID(id uint) error {
	return global.DB.Delete(&FileInfo{}, id).Error
}
