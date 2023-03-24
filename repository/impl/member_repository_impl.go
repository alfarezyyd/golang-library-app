package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type MemberRepositoryImpl struct {
}

func NewMemberRepositoryImpl() *MemberRepositoryImpl {
	return &MemberRepositoryImpl{}
}

func (memberRepository *MemberRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.Member {
	var allMember []domain.Member
	resultQuery := tx.WithContext(ctx.Request.Context()).Find(&allMember)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allMember
}

func (memberRepository *MemberRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, memberID *int) domain.Member {
	var memberData domain.Member
	resultQuery := tx.WithContext(ctx.Request.Context()).First(&memberData, memberID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return memberData
}

func (memberRepository *MemberRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, memberID *int) domain.Member {
	var memberData domain.Member
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().First(&memberData, memberID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return memberData
}

func (memberRepository *MemberRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.Member {
	var allDeletedMember []domain.Member
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().Find(&allDeletedMember)
	if resultQuery.Error != nil {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allDeletedMember
}

func (memberRepository *MemberRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, memberDomain *domain.Member) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Omit("updated_at").Create(&memberDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (memberRepository *MemberRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, memberDomain *domain.Member) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Save(&memberDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (memberRepository *MemberRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, memberID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Debug().Delete(&domain.Member{}, memberID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}

func (memberRepository *MemberRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, memberID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.Member{}, memberID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}
