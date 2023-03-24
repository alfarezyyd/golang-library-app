package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/employee"
	"golang-library-app/model/web/response"
	"golang-library-app/repository"
	"gorm.io/gorm"
	"time"
)

type EmployeeUsecaseImpl struct {
	EmployeeRepository repository.EmployeeRepository
	DB                 *gorm.DB
}

func NewEmployeeUsecaseImpl(employeeRepository repository.EmployeeRepository, DB *gorm.DB) *EmployeeUsecaseImpl {
	return &EmployeeUsecaseImpl{EmployeeRepository: employeeRepository, DB: DB}
}

func (employeeUsecase *EmployeeUsecaseImpl) FindAll(ctx *gin.Context) []response.EmployeeResponse {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allEmployee := employeeUsecase.EmployeeRepository.FindAll(ctx, tx)
	var allEmployeeResponse []response.EmployeeResponse
	for _, employeeData := range allEmployee {
		employeeResponse := helper.ConvertToEmployeeResponse(&employeeData)
		allEmployeeResponse = append(allEmployeeResponse, employeeResponse)
	}
	return allEmployeeResponse
}

func (employeeUsecase *EmployeeUsecaseImpl) FindById(ctx *gin.Context, employeeID *int) response.EmployeeResponse {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	employeeData := employeeUsecase.EmployeeRepository.FindByID(ctx, tx, employeeID)
	return helper.ConvertToEmployeeResponse(&employeeData)
}

func (employeeUsecase *EmployeeUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.EmployeeResponse {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allDeletedEmployee := employeeUsecase.EmployeeRepository.FindAllDeleted(ctx, tx)
	var allEmployeeResponse []response.EmployeeResponse
	for _, employeeData := range allDeletedEmployee {
		employeeResponse := helper.ConvertToEmployeeResponse(&employeeData)
		allEmployeeResponse = append(allEmployeeResponse, employeeResponse)
	}
	return allEmployeeResponse

}

func (employeeUsecase *EmployeeUsecaseImpl) Create(ctx *gin.Context, employeeCreateRequest *employee.CreateRequestEmployee) response.EmployeeResponse {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	newEmployee := domain.Employee{
		Name:            employeeCreateRequest.Name,
		Gender:          employeeCreateRequest.Gender,
		Position:        employeeCreateRequest.Position,
		TelephoneNumber: employeeCreateRequest.TelephoneNumber,
		Address:         employeeCreateRequest.Address,
		CreatedAt:       time.Now(),
	}
	employeeUsecase.EmployeeRepository.Create(ctx, tx, &newEmployee)
	return helper.ConvertToEmployeeResponse(&newEmployee)

}

func (employeeUsecase *EmployeeUsecaseImpl) Update(ctx *gin.Context, employeeUpdateRequest *employee.UpdateRequestEmployee) response.EmployeeResponse {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	employeeData := employeeUsecase.EmployeeRepository.FindByID(ctx, tx, &employeeUpdateRequest.ID)
	employeeData.Name = employeeUpdateRequest.Name
	employeeData.Gender = employeeUpdateRequest.Gender
	employeeData.Position = employeeUpdateRequest.Position
	employeeData.TelephoneNumber = employeeUpdateRequest.TelephoneNumber
	employeeData.Address = employeeUpdateRequest.Address
	employeeUsecase.EmployeeRepository.Update(ctx, tx, &employeeData)
	return helper.ConvertToEmployeeResponse(&employeeData)
}

func (employeeUsecase *EmployeeUsecaseImpl) Delete(ctx *gin.Context, employeeID *int) {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	employeeUsecase.EmployeeRepository.FindByID(ctx, tx, employeeID)
	employeeUsecase.EmployeeRepository.Delete(ctx, tx, employeeID)
}

func (employeeUsecase *EmployeeUsecaseImpl) PermanentDelete(ctx *gin.Context, employeeID *int) {
	tx := employeeUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	employeeUsecase.EmployeeRepository.FindDeletedByID(ctx, tx, employeeID)
	employeeUsecase.EmployeeRepository.Delete(ctx, tx, employeeID)
}
