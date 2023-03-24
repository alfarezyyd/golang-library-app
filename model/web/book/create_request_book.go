package book

type CreateRequestBook struct {
	PublisherID     int    `json:"publisher_id,omitempty" binding:"required,number"`
	ISBN            string `json:"isbn,omitempty" binding:"required"`
	Title           string `json:"title,omitempty" binding:"required"`
	Author          string `json:"author,omitempty" binding:"required"`
	PublicationYear string `json:"publication_year,omitempty" binding:"required"`
	Amount          int    `json:"amount,omitempty" binding:"required"`
	Bookshelf       string `json:"bookshelf,omitempty" binding:"required"`
	KindsID         []int  `json:"kinds_id,omitempty"`
}
