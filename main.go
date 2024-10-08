package main

import (
	"github.com/cyansobble/article"
	"github.com/cyansobble/db"
	"github.com/cyansobble/global"
	"github.com/cyansobble/upload"
	"github.com/cyansobble/user"
	"go.uber.org/zap"
)

func main() {
	// config  viper
	Viper()
	// log zap
	global.LOGGER = InitZapLogger()
	// database gorm
	global.DB = db.ConnectMysqlDb()
	if global.DB == nil {
		global.LOGGER.Info("connect database failed")
		return
	}
	err := global.DB.AutoMigrate(&user.User{}, &article.Article{}, &upload.FileInfo{})
	if err != nil {
		global.LOGGER.Error("auto migrate failed", zap.Error(err))
		return
	}

	global.RedisDb, _ = db.ConnectRedis()
	if global.RedisDb == nil {
		global.LOGGER.Info("connect redis failed")
	}
	//article.GetArticleByDate()
	Router()

}
