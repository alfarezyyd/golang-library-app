package usecase

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/web/employee"
	"golang-library-app/model/web/response"
)

type EmployeeUsecase interface {
	FindAll(ctx *gin.Context) []response.EmployeeResponse
	FindById(ctx *gin.Context, employeeID *int) response.EmployeeResponse
	FindAllDeleted(ctx *gin.Context) []response.EmployeeResponse
	Create(ctx *gin.Context, employeeCreateRequest *employee.CreateRequestEmployee) response.EmployeeResponse
	Update(ctx *gin.Context, employeeUpdateRequest *employee.UpdateRequestEmployee) response.EmployeeResponse
	Delete(ctx *gin.Context, employeeID *int)
	PermanentDelete(ctx *gin.Context, employeeID *int)
}
