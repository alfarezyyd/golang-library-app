package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepositoryImpl() *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{}
}

func (employeeRepository *EmployeeRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Employee {
	var allEmployee []domain.Employee
	resultQuery := tx.WithContext(ctx.Request.Context()).Find(&allEmployee)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allEmployee
}

func (employeeRepository *EmployeeRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, employeeID *int) domain.Employee {
	var employeeData domain.Employee
	resultQuery := tx.WithContext(ctx.Request.Context()).First(&employeeData, employeeID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return employeeData
}

func (employeeRepository *EmployeeRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, employeeID *int) domain.Employee {
	var employeeData domain.Employee
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().First(&employeeData, employeeID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return employeeData

}

func (employeeRepository *EmployeeRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Employee {
	var allDeletedEmployee []domain.Employee
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().Find(&allDeletedEmployee)
	if resultQuery.Error != nil {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allDeletedEmployee
}

func (employeeRepository *EmployeeRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, employeeDomain *domain.Employee) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Omit("updated_at").Create(&employeeDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (employeeRepository *EmployeeRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, employeeDomain *domain.Employee) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Save(&employeeDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)

}

func (employeeRepository *EmployeeRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, employeeID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Debug().Delete(&domain.Employee{}, employeeID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}

func (employeeRepository *EmployeeRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, employeeID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.Employee{}, employeeID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}
