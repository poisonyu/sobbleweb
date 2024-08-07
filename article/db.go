package article

import (
	"github.com/cyansobble/global"
	"gorm.io/gorm"
)

func QueryAllArticle(db *gorm.DB) ([]Article, error) {
	var articles []Article
	result := db.Select("id", "title", "nickname", "updated_at").Find(&articles)
	//global.LOGGER.Infof("query %s records", result.RowsAffected)
	global.LOGGER.Sugar().Infof("[sugar] query %d records", result.RowsAffected)
	return articles, result.Error
}

func CreateArticle(db *gorm.DB, article Article) error {
	err := db.Create(&article).Error
	return err
}

func GetArticleByID(db *gorm.DB, id string) (Article, error) {
	var article Article
	err := db.First(&article, id).Error
	return article, err
}
