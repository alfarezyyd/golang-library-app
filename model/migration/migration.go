package migration

import (
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Publisher{}, &domain.Book{})
	helper.LogFatalIfError(err)
}
