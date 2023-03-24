package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/member"
	"golang-library-app/usecase"
	"net/http"
	"strconv"
)

type MemberControllerImpl struct {
	memberUsecase usecase.MemberUsecase
}

func NewMemberControllerImpl(memberUsecase usecase.MemberUsecase) *MemberControllerImpl {
	return &MemberControllerImpl{memberUsecase: memberUsecase}
}

func (memberController *MemberControllerImpl) FindAll(ctx *gin.Context) {
	allMember := memberController.memberUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allMember))
}

func (memberController *MemberControllerImpl) FindByID(ctx *gin.Context) {
	memberIdString := ctx.Param("id")
	memberIdInt, err := strconv.Atoi(memberIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataMember := memberController.memberUsecase.FindByID(ctx, &memberIdInt)
	helper.WriteToWebResponse(ctx, dataMember)
}

func (memberController *MemberControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allMemberDeleted := memberController.memberUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allMemberDeleted))
}

func (memberController *MemberControllerImpl) Create(ctx *gin.Context) {
	var createMemberRequest member.CreateRequestMember
	err := ctx.ShouldBindJSON(&createMemberRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newMember := memberController.memberUsecase.Create(ctx, &createMemberRequest)
	helper.WriteToWebResponse(ctx, newMember)
}

func (memberController *MemberControllerImpl) Update(ctx *gin.Context) {
	var updateMemberRequest member.UpdateRequestMember
	err := ctx.ShouldBindJSON(&updateMemberRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updateMember := memberController.memberUsecase.Update(ctx, &updateMemberRequest)
	helper.WriteToWebResponse(ctx, updateMember)
}

func (memberController *MemberControllerImpl) Delete(ctx *gin.Context) {
	memberIdString := ctx.Param("id")
	memberIdInt, err := strconv.Atoi(memberIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	memberController.memberUsecase.Delete(ctx, &memberIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (memberController *MemberControllerImpl) PermanentDelete(ctx *gin.Context) {
	memberIdString := ctx.Param("id")
	memberIdInt, err := strconv.Atoi(memberIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	memberController.memberUsecase.PermanentDelete(ctx, &memberIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
