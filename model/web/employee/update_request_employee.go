package employee

type UpdateRequestEmployee struct {
	ID              int    `json:"id,omitempty" binding:"required"`
	Name            string `json:"name,omitempty"`
	Gender          string `json:"gender,omitempty"`
	Position        string `json:"position,omitempty"`
	TelephoneNumber string `json:"telephone_number,omitempty"`
	Address         string `json:"address,omitempty"`
}
