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
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataPublisher := publisherController.publisherUsecase.FindById(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, dataPublisher)
}

func (publisherController *PublisherControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allPublisherDeleted := publisherController.publisherUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allPublisherDeleted))
}

func (publisherController *PublisherControllerImpl) FindAllBookByPublisher(ctx *gin.Context) {
	publisherIdString := ctx.Param("id")
	publisherIdInt, err := strconv.Atoi(publisherIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	publisherWithBookData := publisherController.publisherUsecase.FindAllBookByPublisher(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, publisherWithBookData)
}

func (publisherController *PublisherControllerImpl) Create(ctx *gin.Context) {
	var createPublisherRequest publisher.CreateRequestPublisher
	err := ctx.ShouldBindJSON(&createPublisherRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newPublisher := publisherController.publisherUsecase.Create(ctx, &createPublisherRequest)
	helper.WriteToWebResponse(ctx, newPublisher)
}

func (publisherController *PublisherControllerImpl) Update(ctx *gin.Context) {
	var updatePublisherRequest publisher.UpdateRequestPublisher
	err := ctx.ShouldBindJSON(&updatePublisherRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updatePublisher := publisherController.publisherUsecase.Update(ctx, &updatePublisherRequest)
	helper.WriteToWebResponse(ctx, updatePublisher)
}

func (publisherController *PublisherControllerImpl) Delete(ctx *gin.Context) {
	publisherIdString := ctx.Param("id")
	publisherIdInt, err := strconv.Atoi(publisherIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	publisherController.publisherUsecase.Delete(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (publisherController *PublisherControllerImpl) PermanentDelete(ctx *gin.Context) {
	publisherIdString := ctx.Param("id")
	publisherIdInt, err := strconv.Atoi(publisherIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	publisherController.publisherUsecase.PermanentDelete(ctx, &publisherIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
