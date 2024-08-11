package upload

import "gorm.io/gorm"

type FileInfo struct {
	gorm.Model
	FileName string `gorm:"column:filename"`
	FilePath string `gorm:"column:filepath"`
	Tag      string `gorm:"column:tag"`
	Owner    string `gorm:"column:owner"`
}
