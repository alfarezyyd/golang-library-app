package employee

type CreateRequestEmployee struct {
	Name            string `json:"name,omitempty" binding:"required,max=50"`
	Gender          string `json:"gender,omitempty" binding:"required,enum=male-female"`
	Position        string `json:"position,omitempty" binding:"required,max=50"`
	TelephoneNumber string `json:"telephone_number,omitempty" binding:"required,max=15"`
	Address         string `json:"address,omitempty" binding:"required,max=100"`
}
