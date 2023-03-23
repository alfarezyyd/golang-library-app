package kind

type CreateRequestKind struct {
	Name        string `json:"name,omitempty" binding:"required,max=35"`
	Description string `json:"description"  binding:"required"`
}
