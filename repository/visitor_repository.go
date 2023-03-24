package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type VisitorRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Visitor
	FindByID(ctx *gin.Context, tx *gorm.DB, visitorID *int) domain.Visitor
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, visitorID *int) domain.Visitor
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Visitor
	Create(ctx *gin.Context, tx *gorm.DB, visitorDomain *domain.Visitor)
	Update(ctx *gin.Context, tx *gorm.DB, visitorDomain *domain.Visitor)
	Delete(ctx *gin.Context, tx *gorm.DB, visitorID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, visitorID *int)
}
