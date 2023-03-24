package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type VisitorResponse struct {
	ID        int             `json:"id"`
	NIN       string          `json:"nin"`
	Name      string          `json:"name"`
	Instance  string          `json:"instance"`
	CreatedAt *time.Time      `json:"created_at"`
	UpdatedAt *utils.NullTime `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
