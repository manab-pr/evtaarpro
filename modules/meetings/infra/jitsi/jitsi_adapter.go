package jitsi

import (
	"github.com/manab-pr/evtaarpro/pkg/clients/jitsi"
)

// JitsiAdapter adapts the Jitsi client to use case needs
type JitsiAdapter struct {
	client *jitsi.Client
}

// NewJitsiAdapter creates a new JitsiAdapter
func NewJitsiAdapter(domain, appID, appSecret string) *JitsiAdapter {
	return &JitsiAdapter{
		client: jitsi.NewClient(domain, appID, appSecret),
	}
}

// CreateRoomToken generates a JWT token for a Jitsi room
func (a *JitsiAdapter) CreateRoomToken(roomName, userID, userName, userEmail string, moderator bool) (string, error) {
	return a.client.CreateRoomToken(roomName, userID, userName, userEmail, moderator)
}

// GetRoomURL returns the full URL for a Jitsi room
func (a *JitsiAdapter) GetRoomURL(roomName string) string {
	return a.client.GetRoomURL(roomName)
}
