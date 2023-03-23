package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type PublisherRepositoryImpl struct {
}

func NewPublisherRepositoryImpl() *PublisherRepositoryImpl {
	return &PublisherRepositoryImpl{}
}

func (publisherRepository *PublisherRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Publisher {
	var allPublisher []domain.Publisher
	resultQuery := tx.WithContext(ctx.Request.Context()).Find(&allPublisher)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allPublisher
}

func (publisherRepository *PublisherRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, publisherID *int) domain.Publisher {
	var publisherData domain.Publisher
	resultQuery := tx.WithContext(ctx.Request.Context()).First(&publisherData, publisherID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return publisherData
}

func (publisherRepository *PublisherRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, publisherID *int) domain.Publisher {
	var publisherData domain.Publisher
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().First(&publisherData, publisherID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return publisherData
}

func (publisherRepository *PublisherRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Publisher {
	var allDeletedPublisher []domain.Publisher
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().Find(&allDeletedPublisher)
	if resultQuery.Error != nil {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allDeletedPublisher
}

func (publisherRepository *PublisherRepositoryImpl) FindAllBookByPublisher(ctx *gin.Context, tx *gorm.DB, publisherID *int) domain.Publisher {
	var publisherData domain.Publisher
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().Preload("Books").Find(&publisherData, publisherID)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return publisherData
}

func (publisherRepository *PublisherRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, publisherDomain *domain.Publisher) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Omit("updated_at").Create(&publisherDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (publisherRepository *PublisherRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, publisherDomain *domain.Publisher) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Save(&publisherDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (publisherRepository *PublisherRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, publisherID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Debug().Delete(&domain.Publisher{}, publisherID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}

func (publisherRepository *PublisherRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, publisherID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.Publisher{}, publisherID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}
