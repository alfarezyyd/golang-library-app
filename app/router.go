package app

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/controller"
)

func NewRouter(
	bookController controller.BookController,
	kindController controller.KindController,
	publisherController controller.PublisherController,
	visitorController controller.VisitorController,
	employeeController controller.EmployeeController,
	memberController controller.MemberController,
	userController controller.UserController) *gin.Engine {
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
	v1API.DELETE("/kinds/books/:kind_id/:book_id", kindController.DeleteBookByKind)

	// Visitor API
	v1API.GET("/visitors", visitorController.FindAll)
	v1API.GET("/visitors/:id", visitorController.FindByID)
	v1API.GET("/visitors/deleted", visitorController.FindAllDeleted)
	v1API.POST("/visitors", visitorController.Create)
	v1API.PATCH("/visitors/:id", visitorController.Update)
	v1API.DELETE("/visitors/:id", visitorController.Delete)
	v1API.DELETE("/visitors/delete/:id", visitorController.PermanentDelete)

	// Employee API
	v1API.GET("/employees", employeeController.FindAll)
	v1API.GET("/employees/:id", employeeController.FindByID)
	v1API.GET("/employees/deleted", employeeController.FindAllDeleted)
	v1API.POST("/employees", employeeController.Create)
	v1API.PATCH("/employees/:id", employeeController.Update)
	v1API.DELETE("/employees/:id", employeeController.Delete)
	v1API.DELETE("/employees/delete/:id", employeeController.PermanentDelete)

	// Member API
	v1API.GET("/members", memberController.FindAll)
	v1API.GET("/members/:id", memberController.FindByID)
	v1API.GET("/members/deleted", memberController.FindAllDeleted)
	v1API.POST("/members", memberController.Create)
	v1API.PATCH("/members/:id", memberController.Update)
	v1API.DELETE("/members/:id", memberController.Delete)
	v1API.DELETE("/members/delete/:id", memberController.PermanentDelete)

	// User API
	v1API.GET("/users", userController.FindAll)
	v1API.GET("/users/:id", userController.FindByID)
	v1API.GET("/users/deleted", userController.FindAllDeleted)
	v1API.POST("/users", userController.Create)
	v1API.PATCH("/users/:id", userController.Update)
	v1API.DELETE("/users/:id", userController.Delete)
	v1API.DELETE("/users/delete/:id", userController.PermanentDelete)
	return ginEngine
}
