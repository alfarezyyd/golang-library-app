package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type MemberRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Member
	FindByID(ctx *gin.Context, tx *gorm.DB, memberID *int) domain.Member
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, memberID *int) domain.Member
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Member
	Create(ctx *gin.Context, tx *gorm.DB, memberDomain *domain.Member)
	Update(ctx *gin.Context, tx *gorm.DB, memberDomain *domain.Member)
	Delete(ctx *gin.Context, tx *gorm.DB, memberID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, memberID *int)
}
