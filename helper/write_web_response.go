package helper

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/exception"
	"golang-library-app/model/web"
	"net/http"
)

func WriteToWebResponse(ctx *gin.Context, returnData interface{}) {
	for _, err := range ctx.Errors {
		switch err.Err.(type) {
		case *exception.APIError:
			apiError(ctx, err.Err.(*exception.APIError))
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
