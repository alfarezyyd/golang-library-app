package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type VisitorRepositoryImpl struct {
}

func NewVisitorRepositoryImpl() *VisitorRepositoryImpl {
	return &VisitorRepositoryImpl{}
}

func (visitorRepository *VisitorRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Visitor {
	var allVisitor []domain.Visitor
	resultQuery := tx.WithContext(ctx.Request.Context()).Find(&allVisitor)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allVisitor
}

func (visitorRepository *VisitorRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, visitorID *int) domain.Visitor {
	var visitorData domain.Visitor
	resultQuery := tx.WithContext(ctx.Request.Context()).First(&visitorData, visitorID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return visitorData
}

func (visitorRepository *VisitorRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, visitorID *int) domain.Visitor {
	var visitorData domain.Visitor
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().First(&visitorData, visitorID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return visitorData

}

func (visitorRepository *VisitorRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Visitor {
	var allDeletedVisitor []domain.Visitor
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().Find(&allDeletedVisitor)
	if resultQuery.Error != nil {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allDeletedVisitor
}

func (visitorRepository *VisitorRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, visitorDomain *domain.Visitor) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Omit("updated_at").Create(&visitorDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (visitorRepository *VisitorRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, visitorDomain *domain.Visitor) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Save(&visitorDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)

}

func (visitorRepository *VisitorRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, visitorID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Debug().Delete(&domain.Publisher{}, visitorID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}

func (visitorRepository *VisitorRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, visitorID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.Publisher{}, visitorID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}
