package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/kind"
	"golang-library-app/usecase"
	"strconv"
)

type KindControllerImpl struct {
	kindUsecase usecase.KindUsecase
}

func NewKindControllerImpl(kindUsecase usecase.KindUsecase) *KindControllerImpl {
	return &KindControllerImpl{kindUsecase: kindUsecase}
}

func (kindController *KindControllerImpl) FindAll(ctx *gin.Context) {
	allKind := kindController.kindUsecase.FindAll(ctx)
	helper.WriteToWebResponse(ctx, allKind)
}

func (kindController *KindControllerImpl) FindByID(ctx *gin.Context) {
	kindIdString := ctx.Param("id")
	kindIdInt, err := strconv.Atoi(kindIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataPublisher := kindController.kindUsecase.FindByID(ctx, &kindIdInt)
	helper.WriteToWebResponse(ctx, dataPublisher)
}

func (kindController *KindControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allKind := kindController.kindUsecase.FindAllDeleted(ctx)
	helper.WriteToWebResponse(ctx, allKind)
}

func (kindController *KindControllerImpl) FindAllBookByKind(ctx *gin.Context) {
	kindIdString := ctx.Param("id")
	kindIdInt, err := strconv.Atoi(kindIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	kindData := kindController.kindUsecase.FindAllBookByKind(ctx, &kindIdInt)
	helper.WriteToWebResponse(ctx, kindData)
}

func (kindController *KindControllerImpl) DeleteBookByKind(ctx *gin.Context) {
	bookIdString := ctx.Param("book_id")
	kindIdString := ctx.Param("kind_id")
	bookIdInt, err := strconv.Atoi(bookIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	kindIdInt, err := strconv.Atoi(kindIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	kindController.kindUsecase.DeleteBookByKind(ctx, &kindIdInt, &bookIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (kindController *KindControllerImpl) Create(ctx *gin.Context) {
	var createKindRequest kind.CreateRequestKind
	err := ctx.ShouldBindJSON(&createKindRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newKind := kindController.kindUsecase.Create(ctx, &createKindRequest)
	helper.WriteToWebResponse(ctx, newKind)
}

func (kindController *KindControllerImpl) Update(ctx *gin.Context) {
	var updateKindRequest kind.UpdateRequestKind
	err := ctx.ShouldBindJSON(&updateKindRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updatedKind := kindController.kindUsecase.Update(ctx, &updateKindRequest)
	helper.WriteToWebResponse(ctx, updatedKind)
}

func (kindController *KindControllerImpl) Delete(ctx *gin.Context) {
	kindIdString := ctx.Param("id")
	kindIdInt, err := strconv.Atoi(kindIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	kindController.kindUsecase.Delete(ctx, &kindIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (kindController *KindControllerImpl) PermanentDelete(ctx *gin.Context) {
	kindIdString := ctx.Param("id")
	kindIdInt, err := strconv.Atoi(kindIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	kindController.kindUsecase.PermanentDelete(ctx, &kindIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
