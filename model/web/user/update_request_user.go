package user

type UpdateRequestUser struct {
	ID       int    `json:"id" binding:"required"`
	Username string `json:"username" binding:"max=30"`
	Password string `json:"password" binding:"max=50"`
}
