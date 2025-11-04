package dto

import "time"

// CreateMeetingRequest represents a meeting creation request
type CreateMeetingRequest struct {
	Title           string    `json:"title" binding:"required"`
	Description     string    `json:"description"`
	StartTime       time.Time `json:"start_time" binding:"required"`
	MaxParticipants int       `json:"max_participants"`
}

// MeetingResponse represents a meeting response
type MeetingResponse struct {
	ID              string     `json:"id"`
	RoomID          string     `json:"room_id"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	OrganizerID     string     `json:"organizer_id"`
	StartTime       time.Time  `json:"start_time"`
	EndTime         *time.Time `json:"end_time,omitempty"`
	Status          string     `json:"status"`
	JitsiRoomURL    string     `json:"jitsi_room_url,omitempty"`
	RecordingURL    string     `json:"recording_url,omitempty"`
	MaxParticipants int        `json:"max_participants"`
	CreatedAt       time.Time  `json:"created_at"`
}

// JoinMeetingResponse represents join meeting response
type JoinMeetingResponse struct {
	MeetingID string `json:"meeting_id"`
	RoomURL   string `json:"room_url"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
