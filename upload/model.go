package upload

import "gorm.io/gorm"

type FileInfo struct {
	gorm.Model
	FileName string `gorm:"column:filename" json:"filename"`
	FilePath string `gorm:"column:filepath" json:"filepath"`
	Tag      string `gorm:"column:tag" json:"tag"`
	Owner    string `gorm:"column:owner" json:"owner"`
}

type PageInfo struct {
	PageSize int    `json:"pagesize"`
	Page     int    `json:"page"`
	KeyWord  string `json:"keyword"`
}
