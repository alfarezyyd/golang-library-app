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

func CreateNotFoundError(ctx *gin.Context, apiError error) {
	CreateApiErrorIfError(ctx, apiError, 404, "Not Found")
}

func CheckInternalServerError(ctx *gin.Context, apiError error) {
	CreateApiErrorIfError(ctx, apiError, 500, "Internal Server Error")
}

func CheckBadRequestError(ctx *gin.Context, apiError error) error {
	if apiError != nil {
		err := exception.NewAPIError(400, "Bad Request", apiError.Error())
		ctx.Error(err)
		WriteToWebResponse(ctx, nil)
		return err
	}
	return nil
}

func CheckIfValidationError(ctx *gin.Context, webError error) error {
	if webError != nil {
		ctx.Error(webError)
		WriteToWebResponse(ctx, nil)
		return webError
	}
	return nil
}
