package kind

type UpdateRequestKind struct {
	ID          int    `json:"id,omitempty"  binding:"required"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}
