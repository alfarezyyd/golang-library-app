package repository

import (
	"context"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type PublisherRepository interface {
	FindAll(ctx context.Context, tx *gorm.DB) []domain.Publisher
	FindByID(ctx context.Context, tx *gorm.DB, publisherID *int) (domain.Publisher, error)
	FindDeletedByID(ctx context.Context, tx *gorm.DB, publisherID *int) (domain.Publisher, error)
	FindAllDeleted(ctx context.Context, tx *gorm.DB) []domain.Publisher
	Create(ctx context.Context, tx *gorm.DB, publisherDomain *domain.Publisher) error
	Update(ctx context.Context, tx *gorm.DB, publisherDomain *domain.Publisher) error
	Delete(ctx context.Context, tx *gorm.DB, publisherID *int) error
	PermanentDelete(ctx context.Context, tx *gorm.DB, publisherID *int) error
}
