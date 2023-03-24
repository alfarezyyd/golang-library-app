package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/visitor"
	"golang-library-app/usecase"
	"net/http"
	"strconv"
)

type VisitorControllerImpl struct {
	visitorUsecase usecase.VisitorUsecase
}

func NewVisitorControllerImpl(visitorUsecase usecase.VisitorUsecase) *VisitorControllerImpl {
	return &VisitorControllerImpl{visitorUsecase: visitorUsecase}
}

func (visitorController *VisitorControllerImpl) FindAll(ctx *gin.Context) {
	allVisitor := visitorController.visitorUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allVisitor))
}

func (visitorController *VisitorControllerImpl) FindByID(ctx *gin.Context) {
	visitorIdString := ctx.Param("id")
	visitorIdInt, err := strconv.Atoi(visitorIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataVisitor := visitorController.visitorUsecase.FindById(ctx, &visitorIdInt)
	helper.WriteToWebResponse(ctx, dataVisitor)
}

func (visitorController *VisitorControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allVisitorDeleted := visitorController.visitorUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allVisitorDeleted))
}

func (visitorController *VisitorControllerImpl) Create(ctx *gin.Context) {
	var createVisitorRequest visitor.CreateRequestVisitor
	err := ctx.ShouldBindJSON(&createVisitorRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newVisitor := visitorController.visitorUsecase.Create(ctx, &createVisitorRequest)
	helper.WriteToWebResponse(ctx, newVisitor)
}

func (visitorController *VisitorControllerImpl) Update(ctx *gin.Context) {
	var updateVisitorRequest visitor.UpdateRequestVisitor
	err := ctx.ShouldBindJSON(&updateVisitorRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updateVisitor := visitorController.visitorUsecase.Update(ctx, &updateVisitorRequest)
	helper.WriteToWebResponse(ctx, updateVisitor)
}

func (visitorController *VisitorControllerImpl) Delete(ctx *gin.Context) {
	visitorIdString := ctx.Param("id")
	visitorIdInt, err := strconv.Atoi(visitorIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	visitorController.visitorUsecase.Delete(ctx, &visitorIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (visitorController *VisitorControllerImpl) PermanentDelete(ctx *gin.Context) {
	visitorIdString := ctx.Param("id")
	visitorIdInt, err := strconv.Atoi(visitorIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	visitorController.visitorUsecase.PermanentDelete(ctx, &visitorIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
