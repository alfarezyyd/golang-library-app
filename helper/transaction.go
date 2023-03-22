package helper

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) {
	errorTransaction := recover()
	if errorTransaction != nil {
		errorRollback := tx.Rollback()
		LogFatalIfError(errorRollback.Error)
		panic(errorTransaction)
	} else {
		errorCommit := tx.Commit()
		LogFatalIfError(errorCommit.Error)
	}
}
