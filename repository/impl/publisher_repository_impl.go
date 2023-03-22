package impl

import (
	"context"
	"errors"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type PublisherRepositoryImpl struct {
}

func NewPublisherRepositoryImpl() *PublisherRepositoryImpl {
	return &PublisherRepositoryImpl{}
}

func (publisherRepository *PublisherRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Publisher {
	var allPublisher []domain.Publisher
	resultQuery := tx.WithContext(ctx).Find(&allPublisher)
	helper.LogFatalIfError(resultQuery.Error)
	return allPublisher
}

func (publisherRepository *PublisherRepositoryImpl) FindByID(ctx context.Context, tx *gorm.DB, publisherID *int) (domain.Publisher, error) {
	var publisherData domain.Publisher
	resultQuery := tx.WithContext(ctx).First(&publisherData, publisherID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		return publisherData, errors.New("publisher not found")
	}
	return publisherData, nil
}

func (publisherRepository *PublisherRepositoryImpl) FindDeletedByID(ctx context.Context, tx *gorm.DB, publisherID *int) (domain.Publisher, error) {
	var publisherData domain.Publisher
	resultQuery := tx.WithContext(ctx).Unscoped().First(&publisherData, publisherID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		return publisherData, errors.New("publisher not found")
	}
	return publisherData, nil
}

func (publisherRepository *PublisherRepositoryImpl) FindAllDeleted(ctx context.Context, tx *gorm.DB) []domain.Publisher {
	var allDeletedPublisher []domain.Publisher
	resultQuery := tx.WithContext(ctx).Unscoped().Find(&allDeletedPublisher)
	helper.LogFatalIfError(resultQuery.Error)
	return allDeletedPublisher

}

func (publisherRepository *PublisherRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, publisherDomain *domain.Publisher) error {
	resultManipulation := tx.WithContext(ctx).Debug().Omit("updated_at").Create(&publisherDomain)
	if resultManipulation.Error != nil {
		return errors.New(resultManipulation.Error.Error())
	}
	return nil
}

func (publisherRepository *PublisherRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, publisherDomain *domain.Publisher) error {
	resultManipulation := tx.WithContext(ctx).Save(&publisherDomain)
	if resultManipulation.Error != nil {
		return errors.New(resultManipulation.Error.Error())
	}
	return nil
}

func (publisherRepository *PublisherRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, publisherID *int) error {
	tx.WithContext(ctx).Debug().Delete(&domain.Publisher{}, publisherID)
	return nil
}

func (publisherRepository *PublisherRepositoryImpl) PermanentDelete(ctx context.Context, tx *gorm.DB, publisherID *int) error {
	tx.WithContext(ctx).Unscoped().Delete(&domain.Publisher{}, publisherID)
	return nil
}
