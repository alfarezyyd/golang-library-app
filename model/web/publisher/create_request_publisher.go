package publisher

type CreateRequestPublisher struct {
	Name        string `json:"name" binding:"required,max=35"`
	Description string `json:"description"  binding:"required"`
}
