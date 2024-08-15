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

func CreateArticle(article Article) (uint, error) {
	result := global.DB.Create(&article)
	return article.ID, result.Error
}

func GetArticleByID(id string) (Article, error) {
	var article Article
	err := global.DB.First(&article, id).Error
	return article, err
}

func SaveArticle(article Article) error {
	result := global.DB.Save(&article)

	return result.Error
}
