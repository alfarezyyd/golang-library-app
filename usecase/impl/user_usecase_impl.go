package impl

import (
	"github.com/gin-gonic/gin"
	"golang-library-app/helper"
	"golang-library-app/model/domain"
	"golang-library-app/model/web/response"
	"golang-library-app/model/web/user"
	"golang-library-app/repository"
	"gorm.io/gorm"
	"time"
)

type UserUsecaseImpl struct {
	UserRepository   repository.UserRepository
	MemberRepository repository.MemberRepository
	DB               *gorm.DB
}

func NewUserUsecaseImpl(userRepository repository.UserRepository, DB *gorm.DB) *UserUsecaseImpl {
	return &UserUsecaseImpl{UserRepository: userRepository, DB: DB}
}

func (userUsecase *UserUsecaseImpl) FindAll(ctx *gin.Context) []response.UserResponse {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allUser := userUsecase.UserRepository.FindAll(ctx, tx)
	var allUserResponse []response.UserResponse
	for _, userData := range allUser {
		userResponse := helper.ConvertToUserResponse(&userData)
		allUserResponse = append(allUserResponse, userResponse)
	}
	return allUserResponse
}

func (userUsecase *UserUsecaseImpl) FindByID(ctx *gin.Context, userID *int) response.UserResponse {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	userData := userUsecase.UserRepository.FindByID(ctx, tx, userID)
	return helper.ConvertToUserResponse(&userData)
}

func (userUsecase *UserUsecaseImpl) FindAllDeleted(ctx *gin.Context) []response.UserResponse {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	allDeletedUser := userUsecase.UserRepository.FindAllDeleted(ctx, tx)
	var allUserResponse []response.UserResponse
	for _, userData := range allDeletedUser {
		userResponse := helper.ConvertToUserResponse(&userData)
		allUserResponse = append(allUserResponse, userResponse)
	}
	return allUserResponse
}

func (userUsecase *UserUsecaseImpl) Create(ctx *gin.Context, userCreateRequest *user.CreateRequestUser) response.UserResponse {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	encryptedString := helper.HashPassword(&userCreateRequest.Password)
	userUsecase.MemberRepository.FindByID(ctx, tx, &userCreateRequest.MemberID)
	newUser := domain.User{
		MemberID:  userCreateRequest.MemberID,
		Username:  userCreateRequest.Username,
		Email:     userCreateRequest.Email,
		Password:  encryptedString,
		CreatedAt: time.Now(),
	}
	userUsecase.UserRepository.Create(ctx, tx, &newUser)
	return helper.ConvertToUserResponse(&newUser)
}

func (userUsecase *UserUsecaseImpl) Update(ctx *gin.Context, userUpdateRequest *user.UpdateRequestUser) response.UserResponse {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	encryptedString := helper.HashPassword(&userUpdateRequest.Password)
	userData := userUsecase.UserRepository.FindByID(ctx, tx, &userUpdateRequest.ID)
	userData.Username = userUpdateRequest.Username
	userData.Password = encryptedString
	return helper.ConvertToUserResponse(&userData)
}

func (userUsecase *UserUsecaseImpl) Delete(ctx *gin.Context, userID *int) {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	userUsecase.UserRepository.FindByID(ctx, tx, userID)
	userUsecase.UserRepository.Delete(ctx, tx, userID)
}

func (userUsecase *UserUsecaseImpl) PermanentDelete(ctx *gin.Context, userID *int) {
	tx := userUsecase.DB.Begin()
	defer helper.CommitOrRollback(tx)
	userUsecase.UserRepository.FindDeletedByID(ctx, tx, userID)
	userUsecase.UserRepository.Delete(ctx, tx, userID)
}
