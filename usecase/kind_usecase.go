package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/kind"
	"golang-library-app/model/web/response"
)

type KindUsecase interface {
	FindAll(ctx *gin.Context) []response.KindResponse
	FindByID(ctx *gin.Context, kindID *int) response.KindResponse
	FindAllDeleted(ctx *gin.Context) []response.KindResponse
	FindAllBookByKind(ctx *gin.Context, kindID *int) response.KindResponse
	DeleteBookByKind(ctx *gin.Context, kindID *int, bookID *int)
	Create(ctx *gin.Context, kindCreateRequest *kind.CreateRequestKind) response.KindResponse
	Update(ctx *gin.Context, kindUpdateRequest *kind.UpdateRequestKind) response.KindResponse
	Delete(ctx *gin.Context, kindID *int)
	PermanentDelete(ctx *gin.Context, kindID *int)
}
