package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type BookResponse struct {
	ID              int             `json:"id"`
	PublisherID     int             `json:"publisher_id"`
	ISBN            string          `json:"isbn"`
	Title           string          `json:"title"`
	Author          string          `json:"author"`
	PublicationYear string          `json:"publication_year"`
	Amount          int             `json:"amount"`
	Bookshelf       string          `json:"bookshelf"`
	CreatedAt       *time.Time      `json:"created_at"`
	UpdatedAt       *utils.NullTime `json:"updated_at"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at"`
	Kinds           []KindResponse  `json:"kinds"`
}
