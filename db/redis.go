package db

import (
	"database/sql"
	"fmt"

	"github.com/cyansobble/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectRedis() (rdb *redis.Client, err error) {
	host := global.CONFIG.Redis.Host
	port := global.CONFIG.Redis.Port
	user := global.CONFIG.Redis.User
	pass := global.CONFIG.Redis.PassWord
	db := global.CONFIG.Redis.Db
	redisURL := fmt.Sprintf("redis://%s:%s@%s:%s/%s", user, pass, host, port, db)
	global.LOGGER.Info("redisURL: " + redisURL)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		global.LOGGER.Error("connect redis", zap.Error(err))
		return
	}
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// 	DB:   0,
	// })
	rdb = redis.NewClient(opt)
	return

	// val, err := rdb.Get(ctx, "foo").Result()
	// d := 5 * time.Minute
	// v, err := rdb.Set(ctx, "id123", "12345", d).Result()
	// fmt.Printf("val:%s\nv:%s\n", val, v)
}

func ConnectMysqlDb() *gorm.DB {
	mysqlConfig := global.CONFIG.Mysql
	sqlDB, err := sql.Open("mysql", mysqlConfig.SqlDsn())
	if err != nil {
		global.LOGGER.Error("sql connect mysql database failed", zap.Error(err))
		return nil
	}
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", mysqlConfig.DbName)
	_, err = sqlDB.Exec(createSql)
	if err != nil {
		global.LOGGER.Error("create database failed", zap.Error(err))
		return nil
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		global.LOGGER.Error("gorm connect mysql database failed", zap.Error(err))
	}
	err = gormDB.Exec(fmt.Sprintf("use %s", mysqlConfig.DbName)).Error
	if err != nil {
		global.LOGGER.Error("select database failed", zap.Error(err))
	}

	return gormDB

}
