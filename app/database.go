package app

import (
	"database/sql"
	"golang-library-app/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewSetupDatabase() *gorm.DB {
	sqlDB, _ := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang-library-app?charset=utf8mb4&parseTime=True&loc=Local")
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	helper.LogFatalIfError(err)
	return gormDB
}
