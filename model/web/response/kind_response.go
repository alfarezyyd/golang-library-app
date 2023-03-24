package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type KindResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   *time.Time      `json:"created_at"`
	UpdatedAt   *utils.NullTime `json:"updated_at"`
	Deleted     *gorm.DeletedAt `json:"deleted_at"`
	Books       []BookResponse  `json:"books"`
}
