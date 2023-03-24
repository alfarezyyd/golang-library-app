package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type KindRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Kind
	FindByID(ctx *gin.Context, tx *gorm.DB, kindID *int) domain.Kind
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, kindID *int) domain.Kind
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Kind
	FindAllBookByKind(ctx *gin.Context, tx *gorm.DB, kindID *int) domain.Kind
	DeleteBookByKind(ctx *gin.Context, tx *gorm.DB, kindID *int, bookID *int)
	Create(ctx *gin.Context, tx *gorm.DB, kindDomain *domain.Kind)
	Update(ctx *gin.Context, tx *gorm.DB, kindDomain *domain.Kind)
	Delete(ctx *gin.Context, tx *gorm.DB, kindID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, kindID *int)
}
