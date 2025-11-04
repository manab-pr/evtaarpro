package usecases

import (
	"context"
	"errors"
)

// JitsiService defines Jitsi-related operations
type JitsiService interface {
	CreateRoomToken(roomName, userID, userName, userEmail string, moderator bool) (string, error)
	GetRoomURL(roomName string) string
}

// JoinMeetingUseCase handles joining a meeting
type JoinMeetingUseCase struct {
	meetingRepo  MeetingRepository
	jitsiService JitsiService
}

// MeetingRepository interface for this use case
type MeetingRepository interface {
	GetByID(ctx context.Context, id string) (*MeetingEntity, error)
	Update(ctx context.Context, meeting *MeetingEntity) error
}

// MeetingEntity represents a meeting (to avoid circular imports)
type MeetingEntity struct {
	ID           string
	RoomID       string
	OrganizerID  string
	Status       string
	JitsiRoomURL string
}

// NewJoinMeetingUseCase creates a new JoinMeetingUseCase
func NewJoinMeetingUseCase(meetingRepo MeetingRepository, jitsiService JitsiService) *JoinMeetingUseCase {
	return &JoinMeetingUseCase{
		meetingRepo:  meetingRepo,
		jitsiService: jitsiService,
	}
}

// JoinOutput represents join meeting output
type JoinOutput struct {
	MeetingID string
	RoomURL   string
	UserName  string
	UserEmail string
}

// Execute joins a meeting
func (uc *JoinMeetingUseCase) Execute(ctx context.Context, meetingID, userID, userName, userEmail string) (*JoinOutput, error) {
	meeting, err := uc.meetingRepo.GetByID(ctx, meetingID)
	if err != nil {
		return nil, err
	}

	if meeting.Status != "scheduled" && meeting.Status != "ongoing" {
		return nil, errors.New("meeting is not active")
	}

	// Get room URL (no JWT needed for public Jitsi)
	roomURL := uc.jitsiService.GetRoomURL(meeting.RoomID)

	// Update meeting status to ongoing if it's scheduled
	if meeting.Status == "scheduled" {
		meeting.Status = "ongoing"
		meeting.JitsiRoomURL = roomURL
		if err := uc.meetingRepo.Update(ctx, meeting); err != nil {
			return nil, err
		}
	}

	return &JoinOutput{
		MeetingID: meetingID,
		RoomURL:   roomURL,
		UserName:  userName,
		UserEmail: userEmail,
	}, nil
}
