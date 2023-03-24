package user

type CreateRequestUser struct {
	NIN      string `json:"nin" binding:"required,max=15"`
	MemberID int    `json:"member_id" binding:"required"`
	Username string `json:"username" binding:"required,max=30"`
	Email    string `json:"email" binding:"required,max=50"`
	Password string `json:"password" binding:"required,max=50"`
}
