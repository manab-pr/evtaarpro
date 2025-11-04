package ports

// TokenGenerator defines methods for token generation
type TokenGenerator interface {
	// GenerateAccessToken generates an access token
	GenerateAccessToken(userID, email, role string) (string, error)

	// GenerateRefreshToken generates a refresh token
	GenerateRefreshToken(userID, email, role string) (string, error)

	// ValidateToken validates a token and returns claims
	ValidateToken(token string) (userID, email, role string, err error)
}
