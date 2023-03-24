package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Book
	FindByID(ctx *gin.Context, tx *gorm.DB, bookID *int) domain.Book
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Book
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, bookID *int) domain.Book
	FindAllKindByBook(ctx *gin.Context, tx *gorm.DB, bookID *int) domain.Book
	DeleteAllKindByBook(ctx *gin.Context, tx *gorm.DB, bookID *int)
	Create(ctx *gin.Context, tx *gorm.DB, bookDomain *domain.Book)
	CreateBookKinds(ctx *gin.Context, tx *gorm.DB, bookId *int, allBookKindID []int)
	Update(ctx *gin.Context, tx *gorm.DB, bookDomain *domain.Book)
	Delete(ctx *gin.Context, tx *gorm.DB, bookID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, bookID *int)
}
