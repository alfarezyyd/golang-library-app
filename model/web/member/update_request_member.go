package member

type UpdateRequestMember struct {
	ID              int    `json:"id" binding:"required"`
	NIN             string `json:"nin,omitempty" binding:"max=15"`
	Name            string `json:"name,omitempty" binding:"max=50"`
	BirthPlace      string `json:"birth_place,omitempty" binding:"max=50"`
	BirthDate       string `json:"birth_date" binding:"required"`
	Address         string `json:"address,omitempty" binding:"max=100"`
	Gender          string `json:"gender,omitempty" binding:"enum=male-female"`
	TelephoneNumber string `json:"telephone_number,omitempty" binding:"max=15"`
}
