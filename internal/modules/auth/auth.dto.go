package auth

type LoginDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=90"`
}

type RegisterDto struct {
	Name     string `json:"name" binding:"required,min=2,max=190"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=90"`
}

type RefreshDto struct {
	AccessToken  string `json:"access_token" binding:"required,min=30,max=50"`
	RefreshToken string `json:"refresh_token" binding:"required,min=30,max=50"`
}

type PasswordDto struct {
	CurrentPassword string `json:"current_password" binding:"required,min=8,max=90"`
	NewPassword     string `json:"new_password" binding:"required,min=8,max=90"`
}
