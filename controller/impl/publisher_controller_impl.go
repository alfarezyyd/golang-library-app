package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/publisher"
	"golang-library-app/usecase"
	"net/http"
	"strconv"
)

type PublisherControllerImpl struct {
	publisherUsecase usecase.PublisherUsecase
}

func NewPublisherControllerImpl(publisherUsecase usecase.PublisherUsecase) *PublisherControllerImpl {
	return &PublisherControllerImpl{publisherUsecase: publisherUsecase}
}

func (publisherController *PublisherControllerImpl) FindAll(ctx *gin.Context) {
	allPublisher := publisherController.publisherUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allPublisher))
}

func (publisherController *PublisherControllerImpl) FindByID(ctx *gin.Context) {
	publisherIdString := ctx.Param("id")
	publisherIdInt, err := strconv.Atoi(publisherIdString)
	helper.LogFatalIfError(err)
	dataPublisher := publisherController.publisherUsecase.FindById(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, dataPublisher)
}

func (publisherController *PublisherControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allPublisherDeleted := publisherController.publisherUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allPublisherDeleted))
}

func (publisherController *PublisherControllerImpl) Create(ctx *gin.Context) {
	var createPublisherRequest publisher.CreateRequestPublisher
	err := ctx.ShouldBindJSON(&createPublisherRequest)
	helper.LogFatalIfError(err)
	newPublisher := publisherController.publisherUsecase.Create(ctx, &createPublisherRequest)
	helper.WriteToWebResponse(ctx, newPublisher)
}

func (publisherController *PublisherControllerImpl) Update(ctx *gin.Context) {
	var updatePublisherRequest publisher.UpdateRequestPublisher
	err := ctx.ShouldBindJSON(&updatePublisherRequest)
	helper.LogFatalIfError(err)
	updatePublisher := publisherController.publisherUsecase.Update(ctx, &updatePublisherRequest)
	helper.WriteToWebResponse(ctx, updatePublisher)
}

func (publisherController *PublisherControllerImpl) Delete(ctx *gin.Context) {
	publisherIdString := ctx.Param("id")
	publisherIdInt, err := strconv.Atoi(publisherIdString)
	helper.LogFatalIfError(err)
	publisherController.publisherUsecase.Delete(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (publisherController *PublisherControllerImpl) PermanentDelete(ctx *gin.Context) {
	publisherIdString := ctx.Param("id")
	publisherIdInt, err := strconv.Atoi(publisherIdString)
	helper.LogFatalIfError(err)
	publisherController.publisherUsecase.PermanentDelete(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
