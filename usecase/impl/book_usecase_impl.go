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
	allBook := bookUsecase.BookRepository.FindAll(ctx.Request.Context(), tx)
	var allBookResponse []response.BookResponse
	for _, bookData := range allBook {
		bookResponse := helper.ConvertToBookResponse(&bookData)
		allBookResponse = append(allBookResponse, bookResponse)
	}
	return allBookResponse
}

func (bookUsecase *BookUsecaseImpl) FindById(ctx *gin.Context, bookID *int) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookData, err := bookUsecase.BookRepository.FindByID(ctx.Request.Context(), tx, bookID)
	helper.CreateApiErrorIfError(ctx, err, 400, "Not Found")
	return helper.ConvertToBookResponse(&bookData)
}

func (bookUsecase *BookUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allBook := bookUsecase.BookRepository.FindAllDeleted(ctx.Request.Context(), tx)
	var allBookResponse []response.BookResponse
	for _, bookData := range allBook {
		bookResponse := helper.ConvertToBookResponse(&bookData)
		allBookResponse = append(allBookResponse, bookResponse)
	}
	return allBookResponse
}

func (bookUsecase *BookUsecaseImpl) Create(ctx *gin.Context, bookCreateRequest *book.CreateRequestBook) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
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
	err := bookUsecase.BookRepository.Create(ctx.Request.Context(), tx, &newBook)
	helper.CreateApiErrorIfError(ctx, err, 500, "Internal Server Error")
	return helper.ConvertToBookResponse(&newBook)
}

func (bookUsecase *BookUsecaseImpl) Update(ctx *gin.Context, bookUpdateRequest *book.UpdateRequestBook) response.BookResponse {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	bookData, err := bookUsecase.BookRepository.FindByID(ctx.Request.Context(), tx, &bookUpdateRequest.ID)
	helper.CreateApiErrorIfError(ctx, err, 400, "Not Found")
	bookData.ISBN = bookUpdateRequest.ISBN
	bookData.Title = bookUpdateRequest.Title
	bookData.Author = bookUpdateRequest.Author
	bookData.PublicationYear = bookUpdateRequest.PublicationYear
	bookData.Amount = bookUpdateRequest.Amount
	bookData.Bookshelf = bookUpdateRequest.Bookshelf
	err = bookUsecase.BookRepository.Update(ctx.Request.Context(), tx, &bookData)
	helper.CreateApiErrorIfError(ctx, err, 500, "Internal Server Error")
	return helper.ConvertToBookResponse(&bookData)
}

func (bookUsecase *BookUsecaseImpl) Delete(ctx *gin.Context, bookID *int) {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	_, err := bookUsecase.BookRepository.FindByID(ctx.Request.Context(), tx, bookID)
	helper.CreateApiErrorIfError(ctx, err, 400, "Not Found")
	err = bookUsecase.BookRepository.Delete(ctx.Request.Context(), tx, bookID)
	helper.CreateApiErrorIfError(ctx, err, 400, "Not Found")
}

func (bookUsecase *BookUsecaseImpl) PermanentDelete(ctx *gin.Context, bookID *int) {
	tx := bookUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	_, err := bookUsecase.BookRepository.FindDeletedByID(ctx.Request.Context(), tx, bookID)
	helper.CreateApiErrorIfError(ctx, err, 400, "Not Found")
	err = bookUsecase.BookRepository.PermanentDelete(ctx.Request.Context(), tx, bookID)
	helper.CreateApiErrorIfError(ctx, err, 400, "Not Found")
}
