package app

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/controller"
)

type (
	BookController      controller.WebController
	PublisherController controller.WebController
)

func NewRouter(bookController BookController, publisherController PublisherController) *gin.Engine {
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	v1API := ginEngine.Group("/v1")
	// Book API
	v1API.GET("/books", bookController.FindAll)
	v1API.GET("/books/:id", bookController.FindByID)
	v1API.GET("/books/deleted", bookController.FindAllDeleted)
	v1API.POST("/books", bookController.Create)
	v1API.PATCH("/books/:id", bookController.Update)
	v1API.DELETE("/books/:id", bookController.Delete)
	v1API.DELETE("/books/delete/:id", bookController.PermanentDelete)

	// Publisher API
	v1API.GET("/publishers", publisherController.FindAll)
	v1API.GET("/publishers/:id", publisherController.FindByID)
	v1API.GET("/publishers/deleted", publisherController.FindAllDeleted)
	v1API.POST("/publishers", publisherController.Create)
	v1API.PATCH("/publishers/:id", publisherController.Update)
	v1API.DELETE("/publishers/:id", publisherController.Delete)
	v1API.DELETE("/publishers/delete/:id", publisherController.PermanentDelete)
	return ginEngine
}
