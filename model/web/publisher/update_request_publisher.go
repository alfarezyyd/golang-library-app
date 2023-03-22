package publisher

type UpdateRequestPublisher struct {
	ID          int    `json:"ID" binding:"required"`
	Name        string `json:"name" binding:"max=35"`
	Description string `json:"description"`
}
