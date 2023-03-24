package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type VisitorResponse struct {
	ID        int            `json:"id,omitempty"`
	NIN       string         `json:"nin,omitempty"`
	Name      string         `json:"name,omitempty"`
	Instance  string         `json:"instance,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt utils.NullTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
