package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/web/employee"
	"golang-library-app/usecase"
	"net/http"
	"strconv"
)

type EmployeeControllerImpl struct {
	employeeUsecase usecase.EmployeeUsecase
}

func NewEmployeeControllerImpl(employeeUsecase usecase.EmployeeUsecase) *EmployeeControllerImpl {
	return &EmployeeControllerImpl{employeeUsecase: employeeUsecase}
}

func (employeeController *EmployeeControllerImpl) FindAll(ctx *gin.Context) {
	allEmployee := employeeController.employeeUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allEmployee))
}

func (employeeController *EmployeeControllerImpl) FindByID(ctx *gin.Context) {
	employeeIdString := ctx.Param("id")
	employeeIdInt, err := strconv.Atoi(employeeIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	dataEmployee := employeeController.employeeUsecase.FindByID(ctx, &employeeIdInt)
	helper.WriteToWebResponse(ctx, dataEmployee)
}

func (employeeController *EmployeeControllerImpl) FindAllDeleted(ctx *gin.Context) {
	allEmployeeDeleted := employeeController.employeeUsecase.FindAllDeleted(ctx)
	ctx.JSON(http.StatusOK, helper.ConvertToWebResponse(allEmployeeDeleted))
}

func (employeeController *EmployeeControllerImpl) Create(ctx *gin.Context) {
	var createEmployeeRequest employee.CreateRequestEmployee
	err := ctx.ShouldBindJSON(&createEmployeeRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	newEmployee := employeeController.employeeUsecase.Create(ctx, &createEmployeeRequest)
	helper.WriteToWebResponse(ctx, newEmployee)
}

func (employeeController *EmployeeControllerImpl) Update(ctx *gin.Context) {
	var updateEmployeeRequest employee.UpdateRequestEmployee
	err := ctx.ShouldBindJSON(&updateEmployeeRequest)
	webValidationError := helper.CheckIfValidationError(ctx, err)
	if webValidationError != nil {
		return
	}
	updateEmployee := employeeController.employeeUsecase.Update(ctx, &updateEmployeeRequest)
	helper.WriteToWebResponse(ctx, updateEmployee)
}

func (employeeController *EmployeeControllerImpl) Delete(ctx *gin.Context) {
	employeeIdString := ctx.Param("id")
	employeeIdInt, err := strconv.Atoi(employeeIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	employeeController.employeeUsecase.Delete(ctx, &employeeIdInt)
	helper.WriteToWebResponse(ctx, nil)
}

func (employeeController *EmployeeControllerImpl) PermanentDelete(ctx *gin.Context) {
	employeeIdString := ctx.Param("id")
	employeeIdInt, err := strconv.Atoi(employeeIdString)
	err = helper.CheckBadRequestError(ctx, err)
	if err != nil {
		return
	}
	employeeController.employeeUsecase.PermanentDelete(ctx, &employeeIdInt)
	helper.WriteToWebResponse(ctx, nil)
}
