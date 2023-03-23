package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/kind"
	"golang-library-app/model/web/response"
)

type KindUsecase interface {
	FindAll(ctx *gin.Context) []response.KindResponse
	FindById(ctx *gin.Context, kindID *int) response.KindResponse
	FindAllDeleted(ctx *gin.Context) []response.KindResponse
	FindAllBookByKind(ctx *gin.Context, kindID *int) response.KindResponse
	Create(ctx *gin.Context, kindCreateRequest *kind.CreateRequestKind) response.KindResponse
	Update(ctx *gin.Context, kindUpdateRequest *kind.UpdateRequestKind) response.KindResponse
	Delete(ctx *gin.Context, kindID *int)
	PermanentDelete(ctx *gin.Context, kindID *int)
}