package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type EmployeeResponse struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Gender          string          `json:"gender"`
	Position        string          `json:"position"`
	TelephoneNumber string          `json:"telephone_number"`
	Address         string          `json:"address"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *utils.NullTime `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at"`
}
