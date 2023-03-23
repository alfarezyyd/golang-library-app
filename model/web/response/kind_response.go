package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type KindResponse struct {
	ID          uint            `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description"`
	CreatedAt   *time.Time      `json:"createdAt"`
	UpdatedAt   *utils.NullTime `json:"updatedAt"`
	Deleted     *gorm.DeletedAt `json:"deleted"`
	Books       []BookResponse  `json:"books,omitempty"`
}
