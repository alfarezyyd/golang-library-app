package main

import (
	"golang-library-app/app"
	"golang-library-app/helper"
	"golang-library-app/model/migration"
)

func main() {
	gormDB := app.NewSetupDatabase()
	ginEngine := InitializedGinEngine(gormDB)
	migration.RunMigration(gormDB)
	err := ginEngine.Run("127.0.0.1:3000")
	helper.LogFatalIfError(err)
}
