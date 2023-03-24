package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.User
	FindByID(ctx *gin.Context, tx *gorm.DB, userID *int) domain.User
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, userID *int) domain.User
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.User
	Create(ctx *gin.Context, tx *gorm.DB, userDomain *domain.User)
	Update(ctx *gin.Context, tx *gorm.DB, userDomain *domain.User)
	Delete(ctx *gin.Context, tx *gorm.DB, userID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, userID *int)
}
