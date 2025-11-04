package dto

// AuthResponse represents an authentication response
type AuthResponse struct {
	UserID       string `json:"user_id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// UserResponse represents a user response
type UserResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
