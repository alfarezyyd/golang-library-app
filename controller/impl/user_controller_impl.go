package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/user"
	"golang-library-app/usecase"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	userUsecase usecase.UserUsecase
}

func NewUserControllerImpl(userUsecase usecase.UserUsecase) *UserControllerImpl {
	return &UserControllerImpl{userUsecase: userUsecase}
}

func (userController *UserControllerImpl) FindAll(ctx *gin.Context) {
	allUser := userController.userUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allUser))
}

func (userController *UserControllerImpl) FindByID(ctx *gin.Context) {
	userIdString := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataUser := userController.userUsecase.FindByID(ctx, &userIdInt)
	helper.WriteToWebResponse(ctx, dataUser)
}

func (userController *UserControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allUserDeleted := userController.userUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allUserDeleted))
}

func (userController *UserControllerImpl) Create(ctx *gin.Context) {
	var createUserRequest user.CreateRequestUser
	err := ctx.ShouldBindJSON(&createUserRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newUser := userController.userUsecase.Create(ctx, &createUserRequest)
	helper.WriteToWebResponse(ctx, newUser)
}

func (userController *UserControllerImpl) Update(ctx *gin.Context) {
	var updateUserRequest user.UpdateRequestUser
	err := ctx.ShouldBindJSON(&updateUserRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updateUser := userController.userUsecase.Update(ctx, &updateUserRequest)
	helper.WriteToWebResponse(ctx, updateUser)
}

func (userController *UserControllerImpl) Delete(ctx *gin.Context) {
	userIdString := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	userController.userUsecase.Delete(ctx, &userIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (userController *UserControllerImpl) PermanentDelete(ctx *gin.Context) {
	userIdString := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	userController.userUsecase.PermanentDelete(ctx, &userIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
