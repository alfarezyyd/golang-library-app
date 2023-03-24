package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/member"
	"golang-library-app/model/web/response"
	"golang-library-app/repository"
	"gorm.io/gorm"
	"time"
)

type MemberUsecaseImpl struct {
	MemberRepository repository.MemberRepository
	DB               *gorm.DB
}

func NewMemberUsecaseImpl(memberRepository repository.MemberRepository, DB *gorm.DB) *MemberUsecaseImpl {
	return &MemberUsecaseImpl{MemberRepository: memberRepository, DB: DB}
}

func (memberUsecase *MemberUsecaseImpl) FindAll(ctx *gin.Context) []response.MemberResponse {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allMember := memberUsecase.MemberRepository.FindAll(ctx, tx)
	var allMemberResponse []response.MemberResponse
	for _, memberData := range allMember {
		memberResponse := helper.ConvertToMemberResponse(&memberData)
		allMemberResponse = append(allMemberResponse, memberResponse)
	}
	return allMemberResponse
}

func (memberUsecase *MemberUsecaseImpl) FindByID(ctx *gin.Context, memberID *int) response.MemberResponse {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	memberData := memberUsecase.MemberRepository.FindByID(ctx, tx, memberID)
	return helper.ConvertToMemberResponse(&memberData)
}

func (memberUsecase *MemberUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.MemberResponse {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allDeletedMember := memberUsecase.MemberRepository.FindAllDeleted(ctx, tx)
	var allMemberResponse []response.MemberResponse
	for _, memberData := range allDeletedMember {
		memberResponse := helper.ConvertToMemberResponse(&memberData)
		allMemberResponse = append(allMemberResponse, memberResponse)
	}
	return allMemberResponse
}

func (memberUsecase *MemberUsecaseImpl) Create(ctx *gin.Context, memberCreateRequest *member.CreateRequestMember) response.MemberResponse {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	convertBirthDate, _ := time.Parse("02-01-2006", memberCreateRequest.BirthDate)
	newMember := domain.Member{
		NIN:             memberCreateRequest.NIN,
		Name:            memberCreateRequest.Name,
		BirthPlace:      memberCreateRequest.BirthPlace,
		BirthDate:       convertBirthDate,
		Address:         memberCreateRequest.Address,
		Gender:          memberCreateRequest.Gender,
		TelephoneNumber: memberCreateRequest.TelephoneNumber,
		CreatedAt:       time.Now(),
	}
	memberUsecase.MemberRepository.Create(ctx, tx, &newMember)
	return helper.ConvertToMemberResponse(&newMember)
}

func (memberUsecase *MemberUsecaseImpl) Update(ctx *gin.Context, memberUpdateRequest *member.UpdateRequestMember) response.MemberResponse {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	convertBirthDate, _ := time.Parse("02-01-2006", memberUpdateRequest.BirthDate)
	memberData := memberUsecase.MemberRepository.FindByID(ctx, tx, &memberUpdateRequest.ID)
	memberData.NIN = memberUpdateRequest.NIN
	memberData.Name = memberUpdateRequest.Name
	memberData.BirthPlace = memberUpdateRequest.BirthPlace
	memberData.BirthDate = convertBirthDate
	memberData.Address = memberUpdateRequest.Address
	memberData.Gender = memberUpdateRequest.Gender
	memberData.TelephoneNumber = memberUpdateRequest.TelephoneNumber
	memberUsecase.MemberRepository.Update(ctx, tx, &memberData)
	return helper.ConvertToMemberResponse(&memberData)
}

func (memberUsecase *MemberUsecaseImpl) Delete(ctx *gin.Context, memberID *int) {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	memberUsecase.MemberRepository.FindByID(ctx, tx, memberID)
	memberUsecase.MemberRepository.Delete(ctx, tx, memberID)
}

func (memberUsecase *MemberUsecaseImpl) PermanentDelete(ctx *gin.Context, memberID *int) {
	tx := memberUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	memberUsecase.MemberRepository.FindDeletedByID(ctx, tx, memberID)
	memberUsecase.MemberRepository.Delete(ctx, tx, memberID)
}
