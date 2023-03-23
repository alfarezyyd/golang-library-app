package controller

import "github.com/gin-gonic/gin"

type PublisherController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindAllDeleted(ctx *gin.Context)
	FindAllBookByPublisher(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	PermanentDelete(ctx *gin.Context)
}
