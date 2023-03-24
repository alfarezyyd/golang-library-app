package impl

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/kind"
	"golang-library-app/model/web/response"
	"golang-library-app/repository"
	"gorm.io/gorm"
	"time"
)

type KindUsecaseImpl struct {
	KindRepository repository.KindRepository
	BookRepository repository.BookRepository
	DB             *gorm.DB
}

func NewKindUsecaseImpl(kindRepository repository.KindRepository, bookRepository repository.BookRepository, DB *gorm.DB) *KindUsecaseImpl {
	return &KindUsecaseImpl{KindRepository: kindRepository, BookRepository: bookRepository, DB: DB}
}

func (kindUsecase *KindUsecaseImpl) FindAll(ctx *gin.Context) []response.KindResponse {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allKind := kindUsecase.KindRepository.FindAll(ctx, tx)
	var allKindResponse []response.KindResponse
	for _, kindData := range allKind {
		kindResponse := helper.ConvertToKindResponse(&kindData)
		allKindResponse = append(allKindResponse, kindResponse)
	}
	return allKindResponse
}

func (kindUsecase *KindUsecaseImpl) FindByID(ctx *gin.Context, kindID *int) response.KindResponse {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	kindData := kindUsecase.KindRepository.FindByID(ctx, tx, kindID)
	return helper.ConvertToKindResponse(&kindData)
}

func (kindUsecase *KindUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.KindResponse {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allKindData := kindUsecase.KindRepository.FindAllDeleted(ctx, tx)

	var allKindResponse []response.KindResponse
	for _, kindData := range allKindData {
		kindResponse := helper.ConvertToKindResponse(&kindData)
		allKindResponse = append(allKindResponse, kindResponse)
	}
	return allKindResponse
}

func (kindUsecase *KindUsecaseImpl) FindAllBookByKind(ctx *gin.Context, kindID *int) response.KindResponse {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	kindData := kindUsecase.KindRepository.FindAllBookByKind(ctx, tx, kindID)
	return helper.ConvertToKindResponse(&kindData)
}

func (kindUsecase *KindUsecaseImpl) DeleteBookByKind(ctx *gin.Context, kindID *int, bookID *int) {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	kindUsecase.KindRepository.FindByID(ctx, tx, kindID)
	kindUsecase.BookRepository.FindByID(ctx, tx, bookID)
	kindUsecase.KindRepository.DeleteBookByKind(ctx, tx, kindID, bookID)
}

func (kindUsecase *KindUsecaseImpl) Create(ctx *gin.Context, kindCreateRequest *kind.CreateRequestKind) response.KindResponse {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	newKind := domain.Kind{
		Name: kindCreateRequest.Name,
		Description: sql.NullString{
			String: kindCreateRequest.Description,
			Valid:  true,
		},
		CreatedAt: time.Now(),
	}
	kindUsecase.KindRepository.Create(ctx, tx, &newKind)
	return helper.ConvertToKindResponse(&newKind)
}

func (kindUsecase *KindUsecaseImpl) Update(ctx *gin.Context, kindUpdateRequest *kind.UpdateRequestKind) response.KindResponse {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	kindData := kindUsecase.KindRepository.FindByID(ctx, tx, &kindUpdateRequest.ID)

	kindData.Name = kindUpdateRequest.Name
	kindData.Description = sql.NullString{String: kindUpdateRequest.Description, Valid: true}
	kindUsecase.KindRepository.Update(ctx, tx, &kindData)

	return helper.ConvertToKindResponse(&kindData)
}

func (kindUsecase *KindUsecaseImpl) Delete(ctx *gin.Context, kindID *int) {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	kindUsecase.KindRepository.FindByID(ctx, tx, kindID)
	kindUsecase.KindRepository.Delete(ctx, tx, kindID)
}

func (kindUsecase *KindUsecaseImpl) PermanentDelete(ctx *gin.Context, kindID *int) {
	tx := kindUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	kindUsecase.KindRepository.FindDeletedByID(ctx, tx, kindID)
	kindUsecase.KindRepository.PermanentDelete(ctx, tx, kindID)
}
