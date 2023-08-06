package models

// SignIn is ...
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

// Renew is ...
type Renew struct {
	RefreshToken string `json:"refresh_token"`
}
