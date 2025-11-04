package jitsi

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Client handles Jitsi Meet API operations
type Client struct {
	domain    string
	appID     string
	appSecret string
}

// NewClient creates a new Jitsi client
func NewClient(domain, appID, appSecret string) *Client {
	return &Client{
		domain:    domain,
		appID:     appID,
		appSecret: appSecret,
	}
}

// CreateRoomToken generates a JWT token for a Jitsi room
func (c *Client) CreateRoomToken(roomName, userID, userName, userEmail string, moderator bool) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"iss":   c.appID,
		"sub":   c.domain,
		"aud":   c.appID,
		"room":  roomName,
		"exp":   now.Add(24 * time.Hour).Unix(),
		"nbf":   now.Unix(),
		"iat":   now.Unix(),
		"context": map[string]interface{}{
			"user": map[string]interface{}{
				"id":        userID,
				"name":      userName,
				"email":     userEmail,
				"moderator": moderator,
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.appSecret))
}

// GetRoomURL returns the full URL for a Jitsi room
func (c *Client) GetRoomURL(roomName string) string {
	return fmt.Sprintf("https://%s/%s", c.domain, roomName)
}

// ValidateRoomName checks if a room name is valid
func (c *Client) ValidateRoomName(roomName string) bool {
	return len(roomName) > 0 && len(roomName) <= 100
}
