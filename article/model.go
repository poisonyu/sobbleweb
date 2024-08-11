package article

import (
	"gorm.io/gorm"
)

type ReqArticle struct {
	NickName string `json:"nickname"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

type Article struct {
	gorm.Model
	Author  string `gorm:"column:author"`
	Title   string `json:"title" gorm:"column:title"`
	Type    string `json:"type" gorm:"column:type"`
	Content string `json:"content" gorm:"column:content"`
}
