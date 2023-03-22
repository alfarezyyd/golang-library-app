package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type PublisherResponse struct {
	ID          uint            `json:"id,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *utils.NullTime `json:"updated_at,omitempty"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at,omitempty"`
	BooksData   []BookResponse  `json:"books_data,omitempty"`
}
