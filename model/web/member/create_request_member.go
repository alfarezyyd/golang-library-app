package member

type CreateRequestMember struct {
	NIN             string `json:"nin,omitempty" binding:"required,max=15"`
	Name            string `json:"name,omitempty" binding:"required,max=50"`
	BirthPlace      string `json:"birth_place,omitempty" binding:"required,max=50"`
	BirthDate       string `json:"birth_date" binding:"required,birth_date"`
	Address         string `json:"address,omitempty" binding:"required,max=100"`
	Gender          string `json:"gender,omitempty" binding:"required,enum=Man-Woman"`
	TelephoneNumber string `json:"telephone_number,omitempty" binding:"required,max=15"`
}
