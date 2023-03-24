package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
}

func NewBookRepositoryImpl() *BookRepositoryImpl {
	return &BookRepositoryImpl{}
}

func (bookRepository *BookRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Book {
	var allBook []domain.Book
	resultQuery := tx.WithContext(ctx.Request.Context()).Find(&allBook)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allBook
}

func (bookRepository *BookRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, bookID *int) domain.Book {
	var bookData domain.Book
	resultQuery := tx.WithContext(ctx.Request.Context()).First(&bookData, bookID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return bookData
}

func (bookRepository *BookRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Book {
	var allBook []domain.Book
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().Find(&allBook)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allBook
}

func (bookRepository *BookRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, bookID *int) domain.Book {
	var bookData domain.Book
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().First(&bookData, bookID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return bookData
}

func (bookRepository *BookRepositoryImpl) FindAllKindByBook(ctx *gin.Context, tx *gorm.DB, bookID *int) domain.Book {
	var bookData domain.Book
	resultQuery := tx.WithContext(ctx.Request.Context()).Debug().Preload("Kinds").First(&bookData, bookID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return bookData
}

func (bookRepository *BookRepositoryImpl) DeleteAllKindByBook(ctx *gin.Context, tx *gorm.DB, bookID *int) {
	resultManipulation := tx.WithContext(ctx).Exec("DELETE FROM books_kinds WHERE book_id = ?", bookID)
	if errors.Is(resultManipulation.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultManipulation.Error)
	}
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (bookRepository *BookRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, bookDomain *domain.Book) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Omit("updated_at").Create(&bookDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (bookRepository *BookRepositoryImpl) CreateBookKinds(ctx *gin.Context, tx *gorm.DB, bookID *int, allBookKindID []int) {
	for _, bookKindID := range allBookKindID {
		resultManipulation := tx.WithContext(ctx).Exec("INSERT INTO books_kinds(kind_id, book_id) VALUES (?,?)", bookKindID, bookID)
		helper.CheckInternalServerError(ctx, resultManipulation.Error)
	}
}

func (bookRepository *BookRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, bookDomain *domain.Book) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Save(&bookDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (bookRepository *BookRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, bookID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Delete(&domain.Book{}, bookID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}

func (bookRepository *BookRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, bookID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.Book{}, bookID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}
