package article

import (
	"github.com/cyansobble/global"
)

func QueryAllArticle() ([]Article, error) {
	var articles []Article
	result := global.DB.Select("id", "title", "nickname", "updated_at").Find(&articles)
	//global.LOGGER.Infof("query %s records", result.RowsAffected)
	global.LOGGER.Sugar().Infof("[sugar] query %d records", result.RowsAffected)
	return articles, result.Error
}

func CreateArticle(article Article) error {
	err := global.DB.Create(&article).Error
	return err
}

func GetArticleByID(id string) (Article, error) {
	var article Article
	err := global.DB.First(&article, id).Error
	return article, err
}
