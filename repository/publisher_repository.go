package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type PublisherRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Publisher
	FindByID(ctx *gin.Context, tx *gorm.DB, publisherID *int) domain.Publisher
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, publisherID *int) domain.Publisher
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Publisher
	FindAllBookByPublisher(ctx *gin.Context, tx *gorm.DB, publisherID *int) domain.Publisher
	Create(ctx *gin.Context, tx *gorm.DB, publisherDomain *domain.Publisher)
	Update(ctx *gin.Context, tx *gorm.DB, publisherDomain *domain.Publisher)
	Delete(ctx *gin.Context, tx *gorm.DB, publisherID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, publisherID *int)
}
