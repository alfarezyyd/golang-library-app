package repository

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Employee
	FindByID(ctx *gin.Context, tx *gorm.DB, employeeID *int) domain.Employee
	FindDeletedByID(ctx *gin.Context, tx *gorm.DB, employeeID *int) domain.Employee
	FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Employee
	Create(ctx *gin.Context, tx *gorm.DB, employeeDomain *domain.Employee)
	Update(ctx *gin.Context, tx *gorm.DB, employeeDomain *domain.Employee)
	Delete(ctx *gin.Context, tx *gorm.DB, employeeID *int)
	PermanentDelete(ctx *gin.Context, tx *gorm.DB, employeeID *int)
}
