package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang-library-app/exception"
	"golang-library-app/model/web"
	"net/http"
)

func WriteToWebResponse(ctx *gin.Context, returnData interface{}) {
	if len(ctx.Errors) != 0 {
		switch ctx.Errors[0].Err.(type) {
		case *exception.APIError:
			apiError(ctx, ctx.Errors[0].Err.(*exception.APIError))
			return
		case validator.ValidationErrors:
			validationError(ctx, ctx.Errors[0].Err.(validator.ValidationErrors))
			return
		default:
			internalServerError(ctx)
			return
		}
	}
	ctx.JSON(http.StatusOK, ConvertToWebResponse(returnData))
}

func apiError(ctx *gin.Context, err *exception.APIError) {
	ctx.AbortWithStatusJSON(err.Status, web.ResponseWeb{
		Code:    err.Status,
		Message: err.Title,
		Data:    err.Details,
	})
}

func validationError(ctx *gin.Context, err validator.ValidationErrors) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ResponseWeb{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Data:    err.Error(),
	})
}

func internalServerError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, web.ResponseWeb{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    nil,
	})
}
