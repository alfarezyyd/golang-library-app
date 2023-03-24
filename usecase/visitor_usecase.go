package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/response"
	"golang-library-app/model/web/visitor"
)

type VisitorUsecase interface {
	FindAll(ctx *gin.Context) []response.VisitorResponse
	FindByID(ctx *gin.Context, visitorID *int) response.VisitorResponse
	FindAllDeleted(ctx *gin.Context) []response.VisitorResponse
	Create(ctx *gin.Context, visitorCreateRequest *visitor.CreateRequestVisitor) response.VisitorResponse
	Update(ctx *gin.Context, visitorUpdateRequest *visitor.UpdateRequestVisitor) response.VisitorResponse
	Delete(ctx *gin.Context, visitorID *int)
	PermanentDelete(ctx *gin.Context, visitorID *int)
}
