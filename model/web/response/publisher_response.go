package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type PublisherResponse struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	CreatedAt   *time.Time      `json:"created_at"`
	UpdatedAt   *utils.NullTime `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
	BooksData   []BookResponse  `json:"books_data"`
}
