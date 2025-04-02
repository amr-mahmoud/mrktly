package userDto

type AuthInput struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type AuthOutput struct {
	ID           string `json:"id"`
	Email        string `json:"email,omitempty"`  // Omit if empty
	Phone        string `json:"phone,omitempty"`  // Omit if empty
	Verified     bool   `json:"verified"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}