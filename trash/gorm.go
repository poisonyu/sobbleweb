package main

import (
	"database/sql"
	"fmt"

	"github.com/cyansobble/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO 连接数据库前判断数据库是否创建，如果没创建就创建

func ConnectDatabase() *gorm.DB {
	mysqlConfig := global.CONFIG.Mysql
	sqlDB, err := sql.Open("mysql", mysqlConfig.SqlDsn())
	if err != nil {
		global.LOGGER.Error("sql connect mysql database failed", zap.Error(err))
		return nil
	}
	// defer func(sqlDB *sql.DB) {
	// 	err := sqlDB.Close()
	// 	if err != nil {
	// 		LOGGER.Error("sql close database failed", zap.Error(err))
	// 	}
	// }(sqlDB)
	//sqlDB.Ping()
	// CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
	// create database if not exists `%s` default character set utf8mb4 default collate utf8mb4_general_ci;
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
	//defer sqlDB.Close()
	return gormDB
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN: mysqlConfig.Dsn(),
	// }), &gorm.Config{})
	// if err != nil {
	// 	LOGGER.Error("connect mysql database failed", zap.Error(err))
	// }

	//return db

}
