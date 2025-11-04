package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidMeetingData = errors.New("invalid meeting data")
	ErrMeetingNotActive   = errors.New("meeting is not active")
)

// MeetingStatus represents the status of a meeting
type MeetingStatus string

const (
	StatusScheduled MeetingStatus = "scheduled"
	StatusOngoing   MeetingStatus = "ongoing"
	StatusCompleted MeetingStatus = "completed"
	StatusCancelled MeetingStatus = "cancelled"
)

// Meeting represents a meeting entity
type Meeting struct {
	ID             string
	RoomID         string
	Title          string
	Description    string
	OrganizerID    string
	StartTime      time.Time
	EndTime        *time.Time
	Status         MeetingStatus
	JitsiRoomURL   string
	RecordingURL   string
	MaxParticipants int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewMeeting creates a new meeting entity
func NewMeeting(title, description, organizerID string, startTime time.Time) (*Meeting, error) {
	if title == "" || organizerID == "" {
		return nil, ErrInvalidMeetingData
	}

	roomID := uuid.New().String()
	now := time.Now()

	return &Meeting{
		ID:             uuid.New().String(),
		RoomID:         roomID,
		Title:          title,
		Description:    description,
		OrganizerID:    organizerID,
		StartTime:      startTime,
		Status:         StatusScheduled,
		MaxParticipants: 50,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

// Start starts the meeting
func (m *Meeting) Start(jitsiRoomURL string) error {
	if m.Status != StatusScheduled {
		return errors.New("meeting cannot be started")
	}

	m.Status = StatusOngoing
	m.JitsiRoomURL = jitsiRoomURL
	m.UpdatedAt = time.Now()
	return nil
}

// Complete completes the meeting
func (m *Meeting) Complete(recordingURL string) {
	now := time.Now()
	m.Status = StatusCompleted
	m.EndTime = &now
	if recordingURL != "" {
		m.RecordingURL = recordingURL
	}
	m.UpdatedAt = now
}

// Cancel cancels the meeting
func (m *Meeting) Cancel() error {
	if m.Status == StatusCompleted {
		return errors.New("cannot cancel completed meeting")
	}

	m.Status = StatusCancelled
	m.UpdatedAt = time.Now()
	return nil
}

// IsActive checks if the meeting is active
func (m *Meeting) IsActive() bool {
	return m.Status == StatusScheduled || m.Status == StatusOngoing
}
