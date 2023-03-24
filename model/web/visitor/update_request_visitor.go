package visitor

type UpdateRequestVisitor struct {
	ID       int    `json:"id,omitempty" binding:"required"`
	NIN      string `json:"nin,omitempty"`
	Name     string `json:"name,omitempty"`
	Instance string `json:"instance,omitempty"`
}
