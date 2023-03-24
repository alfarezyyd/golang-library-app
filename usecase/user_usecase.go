package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/response"
	"golang-library-app/model/web/user"
)

type UserUsecase interface {
	FindAll(ctx *gin.Context) []response.UserResponse
	FindByID(ctx *gin.Context, userID *int) response.UserResponse
	FindAllDeleted(ctx *gin.Context) []response.UserResponse
	Create(ctx *gin.Context, userCreateRequest *user.CreateRequestUser) response.UserResponse
	Update(ctx *gin.Context, userUpdateRequest *user.UpdateRequestUser) response.UserResponse
	Delete(ctx *gin.Context, userID *int)
	PermanentDelete(ctx *gin.Context, userID *int)
}
