package security

import (
	"time"

	"github.com/manab-pr/evtaarpro/pkg/jwt"
)

// JWTGenerator implements token generation using JWT
type JWTGenerator struct {
	secret             string
	issuer             string
	accessTokenExpiry  time.Duration
	refreshTokenExpiry time.Duration
}

// NewJWTGenerator creates a new JWTGenerator
func NewJWTGenerator(secret, issuer string, accessTokenExpiry, refreshTokenExpiry time.Duration) *JWTGenerator {
	return &JWTGenerator{
		secret:             secret,
		issuer:             issuer,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
	}
}

// GenerateAccessToken generates an access token
func (g *JWTGenerator) GenerateAccessToken(userID, email, role string) (string, error) {
	return jwt.GenerateToken(userID, email, role, g.issuer, g.secret, g.accessTokenExpiry)
}

// GenerateRefreshToken generates a refresh token
func (g *JWTGenerator) GenerateRefreshToken(userID, email, role string) (string, error) {
	return jwt.GenerateToken(userID, email, role, g.issuer, g.secret, g.refreshTokenExpiry)
}

// ValidateToken validates a token and returns claims
func (g *JWTGenerator) ValidateToken(token string) (userID, email, role string, err error) {
	claims, err := jwt.ValidateToken(token, g.secret)
	if err != nil {
		return "", "", "", err
	}
	return claims.UserID, claims.Email, claims.Role, nil
}
