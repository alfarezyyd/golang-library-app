package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type EmployeeResponse struct {
	ID              int            `json:"id,omitempty"`
	Name            string         `json:"name,omitempty"`
	Gender          string         `json:"gender,omitempty"`
	Position        string         `json:"position,omitempty"`
	TelephoneNumber string         `json:"telephone_number,omitempty"`
	Address         string         `json:"address,omitempty"`
	CreatedAt       time.Time      `json:"created_at,omitempty"`
	UpdatedAt       utils.NullTime `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}
