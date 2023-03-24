package visitor

type CreateRequestVisitor struct {
	NIN      string `json:"nin,omitempty" binding:"required,max=15"`
	Name     string `json:"name,omitempty" binding:"required,max=50"`
	Instance string `json:"instance,omitempty" binding:"required,max=50"`
}
