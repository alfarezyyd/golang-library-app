package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/book"
	"golang-library-app/usecase"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	bookUsecase usecase.BookUsecase
}

func NewBookControllerImpl(bookUsecase usecase.BookUsecase) *BookControllerImpl {
	return &BookControllerImpl{bookUsecase: bookUsecase}
}

func (bookController *BookControllerImpl) FindAll(ctx *gin.Context) {
	allBook := bookController.bookUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allBook))
}

func (bookController *BookControllerImpl) FindByID(ctx *gin.Context) {
	bookIdString := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataBook := bookController.bookUsecase.FindByID(ctx, &bookIdInt)
	helper.WriteToWebResponse(ctx, dataBook)
}

func (bookController *BookControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allBookDeleted := bookController.bookUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allBookDeleted))
}

func (bookController *BookControllerImpl) FindAllKindByBook(ctx *gin.Context) {
	bookIdString := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	bookData := bookController.bookUsecase.FindAllKindByBook(ctx, &bookIdInt)
	helper.WriteToWebResponse(ctx, bookData)
}

func (bookController *BookControllerImpl) Create(ctx *gin.Context) {
	var createBookRequest book.CreateRequestBook
	err := ctx.ShouldBindJSON(&createBookRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newBook := bookController.bookUsecase.Create(ctx, &createBookRequest)
	helper.WriteToWebResponse(ctx, newBook)
}

func (bookController *BookControllerImpl) Update(ctx *gin.Context) {
	var updateBookRequest book.UpdateRequestBook
	err := ctx.ShouldBindJSON(&updateBookRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updateBook := bookController.bookUsecase.Update(ctx, &updateBookRequest)
	helper.WriteToWebResponse(ctx, updateBook)
}

func (bookController *BookControllerImpl) Delete(ctx *gin.Context) {
	bookIdString := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	bookController.bookUsecase.Delete(ctx, &bookIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (bookController *BookControllerImpl) PermanentDelete(ctx *gin.Context) {
	bookIdString := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	bookController.bookUsecase.PermanentDelete(ctx, &bookIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
