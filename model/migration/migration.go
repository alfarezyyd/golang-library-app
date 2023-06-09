package migration

import (
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.Publisher{},
		&domain.Book{},
		&domain.Kind{},
		&domain.Visitor{},
		&domain.Employee{},
		&domain.Member{},
		&domain.User{})
	helper.LogFatalIfError(err)
}
