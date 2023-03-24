package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/member"
	"golang-library-app/model/web/response"
)

type MemberUsecase interface {
	FindAll(ctx *gin.Context) []response.MemberResponse
	FindByID(ctx *gin.Context, memberID *int) response.MemberResponse
	FindAllDeleted(ctx *gin.Context) []response.MemberResponse
	Create(ctx *gin.Context, memberCreateRequest *member.CreateRequestMember) response.MemberResponse
	Update(ctx *gin.Context, memberUpdateRequest *member.UpdateRequestMember) response.MemberResponse
	Delete(ctx *gin.Context, memberID *int)
	PermanentDelete(ctx *gin.Context, memberID *int)
}
