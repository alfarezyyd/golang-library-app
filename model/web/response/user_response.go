package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type UserResponse struct {
	ID              int             `json:"id"`
	MemberID        int             `json:"member_id"`
	Username        string          `json:"username"`
	Email           string          `json:"email"`
	Password        string          `json:"password"`
	EmailVerifiedAt *utils.NullTime `json:"email_verified_at"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *utils.NullTime `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at"`
}
