package impl

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/publisher"
	"golang-library-app/model/web/response"
	"golang-library-app/repository"
	"golang-library-app/usecase"
	"gorm.io/gorm"
	"time"
)

type PublisherUsecaseImpl struct {
	PublisherRepository repository.PublisherRepository
	BookUsecase         usecase.BookUsecase
	DB                  *gorm.DB
}

func NewPublisherUsecaseImpl(publisherRepository repository.PublisherRepository, DB *gorm.DB) *PublisherUsecaseImpl {
	return &PublisherUsecaseImpl{PublisherRepository: publisherRepository, DB: DB}
}

func (publisherUsecase *PublisherUsecaseImpl) FindAll(ctx *gin.Context) []response.PublisherResponse {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allPublisher := publisherUsecase.PublisherRepository.FindAll(ctx, tx)
	var allPublisherResponse []response.PublisherResponse
	for _, publisherData := range allPublisher {
		publisherResponse := helper.ConvertToPublisherResponse(&publisherData)
		allPublisherResponse = append(allPublisherResponse, publisherResponse)
	}
	return allPublisherResponse
}

func (publisherUsecase *PublisherUsecaseImpl) FindById(ctx *gin.Context, publisherID *int) response.PublisherResponse {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	publisherData := publisherUsecase.PublisherRepository.FindByID(ctx, tx, publisherID)
	return helper.ConvertToPublisherResponse(&publisherData)
}

func (publisherUsecase *PublisherUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.PublisherResponse {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allPublisher := publisherUsecase.PublisherRepository.FindAllDeleted(ctx, tx)
	var allPublisherResponse []response.PublisherResponse
	for _, publisherData := range allPublisher {
		publisherResponse := helper.ConvertToPublisherResponse(&publisherData)
		allPublisherResponse = append(allPublisherResponse, publisherResponse)
	}
	return allPublisherResponse
}

func (publisherUsecase *PublisherUsecaseImpl) FindAllBookByPublisher(ctx *gin.Context, publisherID *int) response.PublisherResponse {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	publisherData := publisherUsecase.PublisherRepository.FindAllBookByPublisher(ctx, tx, publisherID)
	return helper.ConvertToPublisherResponse(&publisherData)
}

func (publisherUsecase *PublisherUsecaseImpl) Create(ctx *gin.Context, publisherCreateRequest *publisher.CreateRequestPublisher) response.PublisherResponse {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	newPublisher := domain.Publisher{
		Name: publisherCreateRequest.Name,
		Description: sql.NullString{
			String: publisherCreateRequest.Description,
			Valid:  true,
		},
		CreatedAt: time.Now(),
	}
	publisherUsecase.PublisherRepository.Create(ctx, tx, &newPublisher)
	return helper.ConvertToPublisherResponse(&newPublisher)
}

func (publisherUsecase *PublisherUsecaseImpl) Update(ctx *gin.Context, publisherUpdateRequest *publisher.UpdateRequestPublisher) response.PublisherResponse {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	publisherData := publisherUsecase.PublisherRepository.FindByID(ctx, tx, &publisherUpdateRequest.ID)
	publisherData.Name = publisherUpdateRequest.Name
	publisherData.Description = sql.NullString{
		String: publisherUpdateRequest.Description,
		Valid:  true,
	}
	publisherUsecase.PublisherRepository.Update(ctx, tx, &publisherData)
	return helper.ConvertToPublisherResponse(&publisherData)
}

func (publisherUsecase *PublisherUsecaseImpl) Delete(ctx *gin.Context, publisherID *int) {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	publisherUsecase.PublisherRepository.FindByID(ctx, tx, publisherID)
	publisherUsecase.PublisherRepository.Delete(ctx, tx, publisherID)
}

func (publisherUsecase *PublisherUsecaseImpl) PermanentDelete(ctx *gin.Context, publisherID *int) {
	tx := publisherUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	publisherUsecase.PublisherRepository.FindDeletedByID(ctx, tx, publisherID)
	publisherUsecase.PublisherRepository.PermanentDelete(ctx, tx, publisherID)
}
