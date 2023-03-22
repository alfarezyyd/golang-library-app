package helper

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/exception"
	"log"
)

func LogFatalIfError(err error) {
	if err != nil {
		log.Fatal("Error => ", err)
	}
}

func CreateApiErrorIfError(ctx *gin.Context, apiError error, statusError int, titleError string) {
	if apiError != nil {
		err := exception.NewAPIError(statusError, titleError, apiError.Error())
		ctx.Error(err)
	}
}
