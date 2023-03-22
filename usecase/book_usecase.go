package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/book"
	"golang-library-app/model/web/response"
)

type BookUsecase interface {
	FindAll(ctx *gin.Context) []response.BookResponse
	FindById(ctx *gin.Context, bookID *int) response.BookResponse
	FindAllDeleted(ctx *gin.Context) []response.BookResponse
	Create(ctx *gin.Context, bookCreateRequest *book.CreateRequestBook) response.BookResponse
	Update(ctx *gin.Context, bookUpdateRequest *book.UpdateRequestBook) response.BookResponse
	Delete(ctx *gin.Context, bookID *int)
	PermanentDelete(ctx *gin.Context, publisherID *int)
}
