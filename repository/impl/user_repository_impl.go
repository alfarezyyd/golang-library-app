package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB) []domain.User {
	var allUser []domain.User
	resultQuery := tx.WithContext(ctx.Request.Context()).Find(&allUser)
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allUser
}

func (userRepository *UserRepositoryImpl) FindByID(ctx *gin.Context, tx *gorm.DB, userID *int) domain.User {
	var userData domain.User
	resultQuery := tx.WithContext(ctx.Request.Context()).First(&userData, userID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return userData
}

func (userRepository *UserRepositoryImpl) FindDeletedByID(ctx *gin.Context, tx *gorm.DB, userID *int) domain.User {
	var userData domain.User
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().First(&userData, userID)
	if errors.Is(resultQuery.Error, gorm.ErrRecordNotFound) {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return userData
}

func (userRepository *UserRepositoryImpl) FindAllDeleted(ctx *gin.Context, tx *gorm.DB) []domain.User {
	var allDeletedUser []domain.User
	resultQuery := tx.WithContext(ctx.Request.Context()).Unscoped().Find(&allDeletedUser)
	if resultQuery.Error != nil {
		helper.CreateNotFoundError(ctx, resultQuery.Error)
	}
	helper.CheckInternalServerError(ctx, resultQuery.Error)
	return allDeletedUser
}

func (userRepository *UserRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, userDomain *domain.User) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Debug().Omit("updated_at").Create(&userDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (userRepository *UserRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, userDomain *domain.User) {
	resultManipulation := tx.WithContext(ctx.Request.Context()).Save(&userDomain)
	helper.CheckInternalServerError(ctx, resultManipulation.Error)
}

func (userRepository *UserRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, userID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Debug().Delete(&domain.User{}, userID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}

func (userRepository *UserRepositoryImpl) PermanentDelete(ctx *gin.Context, tx *gorm.DB, userID *int) {
	resultSql := tx.WithContext(ctx.Request.Context()).Unscoped().Delete(&domain.User{}, userID)
	helper.CheckInternalServerError(ctx, resultSql.Error)
}
