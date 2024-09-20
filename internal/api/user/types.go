package user

type CreateData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateData struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
}

type UpdateStaffPasswordData struct {
	Password string `json:"password" binding:"required"`
}

type CreateResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	IsAdmin  bool   `json:"is_admin"`
}

type CreateStaffData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateLoginData struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}
