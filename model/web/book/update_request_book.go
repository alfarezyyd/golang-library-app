package book

type UpdateRequestBook struct {
	ID              int    `json:"id,omitempty" binding:"required"`
	IDPublisher     uint   `json:"id_publisher,omitempty"`
	ISBN            string `json:"isbn,omitempty"`
	Title           string `json:"title,omitempty"`
	Author          string `json:"author,omitempty"`
	PublicationYear string `json:"publication_year,omitempty"`
	Amount          int    `json:"amount,omitempty"`
	Bookshelf       string `json:"bookshelf,omitempty"`
	KindsID         []int  `json:"kinds_id"`
}
