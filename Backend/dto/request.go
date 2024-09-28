package dto

type SignupRes struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	LoginType string `json:"loginType"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UpdateUserRes struct {
	UserID      int    `json:"userId"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type GetUserDetailRes struct {
	UserID int `json:"userId" binding:"required"`
}

type DeleteUserRes struct {
	UserID int `json:"userId" binding:"required"`
}
