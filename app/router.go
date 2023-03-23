package app

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/controller"
)

type (
	BookController      controller.BookController
	KindController      controller.KindController
	PublisherController controller.PublisherController
)

func NewRouter(bookController BookController, kindController KindController, publisherController PublisherController) *gin.Engine {
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	v1API := ginEngine.Group("/v1")
	// Book API
	v1API.GET("/books", bookController.FindAll)
	v1API.GET("/books/:id", bookController.FindByID)
	v1API.GET("/books/deleted", bookController.FindAllDeleted)
	v1API.GET("/books/kinds/:id", bookController.FindAllKindByBook)
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
	v1API.GET("/publishers/books/:id", publisherController.FindAllBookByPublisher)

	// Kind API
	v1API.GET("/kinds", kindController.FindAll)
	v1API.GET("/kinds/:id", kindController.FindByID)
	v1API.GET("/kinds/deleted", kindController.FindAllDeleted)
	v1API.GET("/kinds/books/:id", kindController.FindAllBookByKind)
	v1API.POST("/kinds", kindController.Create)
	v1API.PATCH("/kinds/:id", kindController.Update)
	v1API.DELETE("/kinds/:id", kindController.Delete)
	v1API.DELETE("/kinds/delete/:id", kindController.PermanentDelete)
	return ginEngine
}
