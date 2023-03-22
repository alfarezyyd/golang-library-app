package impl

import (
	"context"
	"errors"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepositoryImpl() *BookRepositoryImpl {
	return &BookRepositoryImpl{}
}

func (bookRepository *BookRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Book {
	var allBook []domain.Book
	resultQuery := tx.WithContext(ctx).Find(&allBook)
	helper.LogFatalIfError(resultQuery.Error)
	return allBook
}

func (bookRepository *BookRepositoryImpl) FindByID(ctx context.Context, tx *gorm.DB, bookID *int) (domain.Book, error) {
	var bookData domain.Book
	resultQuery := tx.WithContext(ctx).First(&bookData, bookID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		return bookData, errors.New("book not found")
	}
	return bookData, nil
}

func (bookRepository *BookRepositoryImpl) FindAllByPublisher(ctx context.Context, tx *gorm.DB, publisherID *int) []domain.Book {
	var allBook []domain.Book
	resultQuery := tx.WithContext(ctx).Debug().Find(&allBook, *publisherID)
	helper.LogFatalIfError(resultQuery.Error)
	return allBook
}

func (bookRepository *BookRepositoryImpl) FindAllDeleted(ctx context.Context, tx *gorm.DB) []domain.Book {
	var allBook []domain.Book
	resultQuery := tx.WithContext(ctx).Unscoped().Find(&allBook)
	helper.LogFatalIfError(resultQuery.Error)
	return allBook
}

func (bookRepository *BookRepositoryImpl) FindDeletedByID(ctx context.Context, tx *gorm.DB, bookID *int) (domain.Book, error) {
	var bookData domain.Book
	resultQuery := tx.WithContext(ctx).Unscoped().First(&bookData, bookID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		return bookData, errors.New("book not found")
	}
	return bookData, nil
}

func (bookRepository *BookRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, bookDomain *domain.Book) error {
	resultManipulation := tx.WithContext(ctx).Debug().Omit("updated_at").Create(&bookDomain)
	if resultManipulation.Error != nil {
		return errors.New(resultManipulation.Error.Error())
	}
	return nil
}

func (bookRepository *BookRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, bookDomain *domain.Book) error {
	resultManipulation := tx.WithContext(ctx).Debug().Save(&bookDomain)
	if resultManipulation.Error != nil {
		return errors.New(resultManipulation.Error.Error())
	}
	return nil
}

func (bookRepository *BookRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, bookID *int) error {
	resultSql := tx.WithContext(ctx).Delete(&domain.Book{}, bookID)
	if errors.Is(resultSql.Error, gorm.ErrRecordNotFound) {
		return errors.New("book not found")
	}
	return nil
}

func (bookRepository *BookRepositoryImpl) PermanentDelete(ctx context.Context, tx *gorm.DB, bookID *int) error {
	resultSql := tx.WithContext(ctx).Unscoped().Delete(&domain.Book{}, bookID)
	if errors.Is(resultSql.Error, gorm.ErrRecordNotFound) {
		return errors.New("book not found")
	}
	return nil
}
