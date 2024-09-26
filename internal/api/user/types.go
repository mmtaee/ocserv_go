package user

type CreateUserBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserPasswordBody struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type UpdateStaffPasswordBody struct {
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	IsAdmin  bool   `json:"is_admin"`
}

type CreateLoginBody struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}
