package main

import (
	"golang-library-app/app"
	"golang-library-app/helper"
	"golang-library-app/model/migration"
	"golang-library-app/utils"
)

func main() {
	utils.RegisterCustomValidator()
	gormDB := app.NewSetupDatabase()
	ginEngine := InitializedGinEngine(gormDB)
	err := ginEngine.SetTrustedProxies([]string{"127.0.0.1"})
	helper.LogFatalIfError(err)
	migration.RunMigration(gormDB)
	err = ginEngine.Run("127.0.0.1:3000")
	helper.LogFatalIfError(err)
}
