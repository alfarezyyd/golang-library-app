package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/book"
	"golang-library-app/model/web/response"
	"golang-library-app/repository"
	"gorm.io/gorm"
	"time"
)

type BookUsecaseImpl struct {
	BookRepository repository.BookRepository
	DB             *gorm.DB
}

func NewBookUsecaseImpl(bookRepository repository.BookRepository, DB *gorm.DB) *BookUsecaseImpl {
	return &BookUsecaseImpl{BookRepository: bookRepository, DB: DB}
}

func (bookUsecase *BookUsecaseImpl) FindAll(ctx *gin.Context) []response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allBook := bookUsecase.BookRepository.FindAll(ctx, tx)
	var allBookResponse []response.BookResponse
	for _, bookData := range allBook {
		bookResponse := helper.ConvertToBookResponse(&bookData)
		allBookResponse = append(allBookResponse, bookResponse)
	}
	return allBookResponse
}

func (bookUsecase *BookUsecaseImpl) FindByID(ctx *gin.Context, bookID *int) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookData := bookUsecase.BookRepository.FindByID(ctx, tx, bookID)
	return helper.ConvertToBookResponse(&bookData)
}

func (bookUsecase *BookUsecaseImpl) FindAllKindByBook(ctx *gin.Context, bookID *int) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookData := bookUsecase.BookRepository.FindAllKindByBook(ctx, tx, bookID)
	return helper.ConvertToBookResponse(&bookData)
}

func (bookUsecase *BookUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allBook := bookUsecase.BookRepository.FindAllDeleted(ctx, tx)
	var allBookResponse []response.BookResponse
	for _, bookData := range allBook {
		bookResponse := helper.ConvertToBookResponse(&bookData)
		allBookResponse = append(allBookResponse, bookResponse)
	}
	return allBookResponse
}

func (bookUsecase *BookUsecaseImpl) Create(ctx *gin.Context, bookCreateRequest *book.CreateRequestBook) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	newBook := domain.Book{
		PublisherID:     bookCreateRequest.PublisherID,
		ISBN:            bookCreateRequest.ISBN,
		Title:           bookCreateRequest.Title,
		Author:          bookCreateRequest.Author,
		PublicationYear: bookCreateRequest.PublicationYear,
		Amount:          bookCreateRequest.Amount,
		Bookshelf:       bookCreateRequest.Bookshelf,
		CreatedAt:       time.Now(),
	}
	bookUsecase.BookRepository.Create(ctx, tx, &newBook)
	helper.CommitOrRollback(tx)
	tx = bookUsecase.DB.Begin()
	newBookIDInt := int(newBook.ID)
	bookUsecase.BookRepository.CreateBookKinds(ctx, tx, &newBookIDInt, bookCreateRequest.KindsID)
	defer helper.CommitOrRollback(tx)
	return helper.ConvertToBookResponse(&newBook)
}

func (bookUsecase *BookUsecaseImpl) Update(ctx *gin.Context, bookUpdateRequest *book.UpdateRequestBook) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookData := bookUsecase.BookRepository.FindByID(ctx, tx, &bookUpdateRequest.ID)
	bookData.ISBN = bookUpdateRequest.ISBN
	bookData.Title = bookUpdateRequest.Title
	bookData.Author = bookUpdateRequest.Author
	bookData.PublicationYear = bookUpdateRequest.PublicationYear
	bookData.Amount = bookUpdateRequest.Amount
	bookData.Bookshelf = bookUpdateRequest.Bookshelf
	bookUsecase.BookRepository.Update(ctx, tx, &bookData)
	return helper.ConvertToBookResponse(&bookData)
}

func (bookUsecase *BookUsecaseImpl) Delete(ctx *gin.Context, bookID *int) {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookUsecase.BookRepository.FindByID(ctx, tx, bookID)
	bookUsecase.BookRepository.Delete(ctx, tx, bookID)
}

func (bookUsecase *BookUsecaseImpl) PermanentDelete(ctx *gin.Context, bookID *int) {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookUsecase.BookRepository.FindDeletedByID(ctx, tx, bookID)
	bookUsecase.BookRepository.PermanentDelete(ctx, tx, bookID)
}
