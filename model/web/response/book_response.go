package response

import (
	"golang-library-app/utils"
	"gorm.io/gorm"
	"time"
)

type BookResponse struct {
	ID              int             `json:"id,omitempty"`
	PublisherID     int             `json:"publisher_id,omitempty"`
	ISBN            string          `json:"isbn,omitempty"`
	Title           string          `json:"title,omitempty"`
	Author          string          `json:"author,omitempty"`
	PublicationYear string          `json:"publication_year,omitempty"`
	Amount          int             `json:"amount,omitempty"`
	Bookshelf       string          `json:"bookshelf,omitempty"`
	CreatedAt       *time.Time      `json:"created_at,omitempty"`
	UpdatedAt       *utils.NullTime `json:"updated_at,omitempty"`
	DeletedAt       *gorm.DeletedAt `json:"deleted_at,omitempty"`
	Kinds           []KindResponse  `json:"kinds"`
}
