package article

import (
	"time"

	"github.com/cyansobble/global"
	"go.uber.org/zap"
)

func QueryAllArticle() ([]Article, error) {
	var articles []Article
	result := global.DB.Select("id", "title", "nickname", "updated_at").Find(&articles)
	//global.LOGGER.Infof("query %s records", result.RowsAffected)
	global.LOGGER.Sugar().Infof("[sugar] query %d records", result.RowsAffected)
	return articles, result.Error
}

func QueryAllArticleDesc() ([]Article, error) {
	var articles []Article
	result := global.DB.Select("id", "title", "nickname", "updated_at").Order("updated_at DESC").Find(&articles)
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

func DeleteArticleByID(id string) error {
	result := global.DB.Delete(&Article{}, id)
	return result.Error
}

func Next(id string) (Article, error) {
	var article Article
	result := global.DB.Where("id > ?", id).First(&article)
	return article, result.Error
}

func Previous(id string) (Article, error) {
	var article Article
	result := global.DB.Where("id < ?", id).Last(&article)
	return article, result.Error
}

func GetArticleDescDate() ([]time.Time, error) {
	var dates []time.Time
	db := global.DB.Model(&Article{})
	err := db.Select("updated_at").Order("updated_at DESC").Find(&dates).Error
	return dates, err

}

func Archives(dates []time.Time) []time.Time {
	if len(dates) == 0 {
		return nil
	}
	var newDates []time.Time
	firstDate := dates[len(dates)-1]
	lastDate := dates[0]
	loc, _ := time.LoadLocation("")
	startDate := time.Date(firstDate.Year(), firstDate.Month(), 1, 0, 0, 0, 0, loc)

	for startDate.Before(lastDate) {
		newDates = append(newDates, startDate)
		startDate = startDate.AddDate(0, 1, 0)
	}
	return newDates
}

func isArticleExist(id string) (Article, bool) {
	a, err := GetArticleByID(id)
	if err != nil {
		global.LOGGER.Error("article don't exist", zap.Error(err))
		return a, false
	}
	return a, true
}
