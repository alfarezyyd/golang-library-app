package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/publisher"
	"golang-library-app/model/web/response"
)

type PublisherUsecase interface {
	FindAll(ctx *gin.Context) []response.PublisherResponse
	FindById(ctx *gin.Context, publisherID *int) response.PublisherResponse
	FindAllDeleted(ctx *gin.Context) []response.PublisherResponse
	Create(ctx *gin.Context, publisherCreateRequest *publisher.CreateRequestPublisher) response.PublisherResponse
	Update(ctx *gin.Context, publisherUpdateRequest *publisher.UpdateRequestPublisher) response.PublisherResponse
	Delete(ctx *gin.Context, publisherID *int)
	PermanentDelete(ctx *gin.Context, publisherID *int)
}
