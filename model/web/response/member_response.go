package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type MemberResponse struct {
	ID              int             `json:"id"`
	NIN             string          `json:"nin"`
	Name            string          `json:"name"`
	BirthPlace      string          `json:"birth_place"`
	BirthDate       *time.Time      `json:"birth_date"`
	Address         string          `json:"address"`
	Gender          string          `json:"gender"`
	TelephoneNumber string          `json:"telephone_number"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *utils.NullTime `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at"`
	User            UserResponse    `json:"user"`
}
