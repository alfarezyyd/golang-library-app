package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/response"
	"golang-library-app/model/web/visitor"
	"golang-library-app/repository"
	"gorm.io/gorm"
	"time"
)

type VisitorUsecaseImpl struct {
	VisitorRepository repository.VisitorRepository
	DB                *gorm.DB
}

func NewVisitorUsecaseImpl(visitorRepository repository.VisitorRepository, DB *gorm.DB) *VisitorUsecaseImpl {
	return &VisitorUsecaseImpl{VisitorRepository: visitorRepository, DB: DB}
}

func (visitorUsecase *VisitorUsecaseImpl) FindAll(ctx *gin.Context) []response.VisitorResponse {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allVisitor := visitorUsecase.VisitorRepository.FindAll(ctx, tx)
	var allVisitorResponse []response.VisitorResponse
	for _, visitorData := range allVisitor {
		visitorResponse := helper.ConvertToVisitorResponse(&visitorData)
		allVisitorResponse = append(allVisitorResponse, visitorResponse)
	}
	return allVisitorResponse
}

func (visitorUsecase *VisitorUsecaseImpl) FindById(ctx *gin.Context, visitorID *int) response.VisitorResponse {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	visitorData := visitorUsecase.VisitorRepository.FindByID(ctx, tx, visitorID)
	return helper.ConvertToVisitorResponse(&visitorData)
}

func (visitorUsecase *VisitorUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.VisitorResponse {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allDeletedVisitor := visitorUsecase.VisitorRepository.FindAllDeleted(ctx, tx)
	var allVisitorResponse []response.VisitorResponse
	for _, visitorData := range allDeletedVisitor {
		visitorResponse := helper.ConvertToVisitorResponse(&visitorData)
		allVisitorResponse = append(allVisitorResponse, visitorResponse)
	}
	return allVisitorResponse
}

func (visitorUsecase *VisitorUsecaseImpl) Create(ctx *gin.Context, visitorCreateRequest *visitor.CreateRequestVisitor) response.VisitorResponse {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	newVisitor := domain.Visitor{
		NIN:       visitorCreateRequest.NIN,
		Name:      visitorCreateRequest.Name,
		Instance:  visitorCreateRequest.Instance,
		CreatedAt: time.Now(),
	}
	visitorUsecase.VisitorRepository.Create(ctx, tx, &newVisitor)
	return helper.ConvertToVisitorResponse(&newVisitor)
}

func (visitorUsecase *VisitorUsecaseImpl) Update(ctx *gin.Context, visitorUpdateRequest *visitor.UpdateRequestVisitor) response.VisitorResponse {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	visitorData := visitorUsecase.VisitorRepository.FindByID(ctx, tx, &visitorUpdateRequest.ID)
	visitorData.NIN = visitorUpdateRequest.NIN
	visitorData.Name = visitorUpdateRequest.Name
	visitorData.Instance = visitorUpdateRequest.Instance
	visitorUsecase.VisitorRepository.Update(ctx, tx, &visitorData)
	return helper.ConvertToVisitorResponse(&visitorData)
}

func (visitorUsecase *VisitorUsecaseImpl) Delete(ctx *gin.Context, visitorID *int) {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	visitorUsecase.VisitorRepository.FindByID(ctx, tx, visitorID)
	visitorUsecase.VisitorRepository.Delete(ctx, tx, visitorID)
}

func (visitorUsecase *VisitorUsecaseImpl) PermanentDelete(ctx *gin.Context, visitorID *int) {
	tx := visitorUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	visitorUsecase.VisitorRepository.FindDeletedByID(ctx, tx, visitorID)
	visitorUsecase.VisitorRepository.Delete(ctx, tx, visitorID)
}
