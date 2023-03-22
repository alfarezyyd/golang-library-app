package repository

import (
	"context"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(ctx context.Context, tx *gorm.DB) []domain.Book
	FindByID(ctx context.Context, tx *gorm.DB, bookID *int) (domain.Book, error)
	FindAllByPublisher(ctx context.Context, tx *gorm.DB, publisherID *int) []domain.Book
	FindDeletedByID(ctx context.Context, tx *gorm.DB, bookID *int) (domain.Book, error)
	FindAllDeleted(ctx context.Context, tx *gorm.DB) []domain.Book
	Create(ctx context.Context, tx *gorm.DB, bookDomain *domain.Book) error
	Update(ctx context.Context, tx *gorm.DB, bookDomain *domain.Book) error
	Delete(ctx context.Context, tx *gorm.DB, bookID *int) error
	PermanentDelete(ctx context.Context, tx *gorm.DB, bookID *int) error
}
