package request

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserOTPRequest struct {
	Email string `json:"email" binding:"required"`
}

type UserOTPVerifyRequest struct {
	Email    string `json:"email" binding:"required"`
	AuthCode string `json:"auth_code" binding:"required"`
}
