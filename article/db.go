package article

import (
	"github.com/cyansobble/global"
	"gorm.io/gorm"
)

func QueryAllArticle(db *gorm.DB) ([]Article, error) {
	var articles []Article
	result := db.Select("title", "nickname", "updated_at").Find(&articles)
	//global.LOGGER.Infof("query %s records", result.RowsAffected)
	global.LOGGER.Sugar().Infof("query %s records\n", result.RowsAffected)
	return articles, result.Error
}

func CreateArticle(db *gorm.DB, article Article) error {
	err := db.Create(&article).Error
	return err
}
