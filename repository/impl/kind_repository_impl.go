package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type KindRepositoryImpl struct {
}

func NewKindRepositoryImpl() *KindRepositoryImpl {
	return &KindRepositoryImpl{}
}

func (kindRepository *KindRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Kind {
	var allKind []domain.Kind
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().Find(&allKind)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allKind
}

func (kindRepository *KindRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, kindID *int) domain.Kind {
	var dataKind domain.Kind
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().First(&dataKind, kindID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return dataKind
}

func (kindRepository *KindRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, kindID *int) domain.Kind {
	var deletedKind domain.Kind
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().Unscoped().First(&deletedKind, kindID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	} else if resultQuery.Error != nil {
		helper.CheckInternalServerError(ctx, resultQuery.Error)
	}
	return deletedKind
}

func (kindRepository *KindRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Kind {
	var allDeletedKind []domain.Kind
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().Unscoped().Find(&allDeletedKind)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allDeletedKind
}

func (kindRepository *KindRepositoryImpl) FindAllBookByKind(ctx *gin.Context, tx *gorm.DB, kindID *int) domain.Kind {
	var kindData domain.Kind
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().Preload("Books").First(&kindData, kindID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return kindData
}

func (kindRepository *KindRepositoryImpl) DeleteBookByKind(ctx *gin.Context, tx *gorm.DB, kindID *int, bookID *int) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Exec("DELETE FROM books_kinds WHERE book_id = ? AND kind_id = ?", bookID, kindID)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (kindRepository *KindRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, domainData *domain.Kind) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Omit("updated_at").Create(&domainData)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (kindRepository *KindRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, domainData *domain.Kind) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Save(&domainData)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (kindRepository *KindRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, kindID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Debug().Delete(&domain.Kind{}, kindID)
	helper.CheckInternalServerError(ctx, resultSql.Error)

}

func (kindRepository *KindRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, kindID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.Kind{}, kindID)
	helper.CheckInternalServerError(ctx, resultSql.Error)

}
