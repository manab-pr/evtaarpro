package usecases

import (
	"context"
	"time"

	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/repository"
)

// CreateMeetingUseCase handles meeting creation
type CreateMeetingUseCase struct {
	meetingRepo repository.MeetingRepository
}

// NewCreateMeetingUseCase creates a new CreateMeetingUseCase
func NewCreateMeetingUseCase(meetingRepo repository.MeetingRepository) *CreateMeetingUseCase {
	return &CreateMeetingUseCase{meetingRepo: meetingRepo}
}

// CreateInput represents meeting creation input
type CreateInput struct {
	Title          string
	Description    string
	OrganizerID    string
	StartTime      time.Time
	MaxParticipants int
}

// Execute creates a new meeting
func (uc *CreateMeetingUseCase) Execute(ctx context.Context, input CreateInput) (*entities.Meeting, error) {
	meeting, err := entities.NewMeeting(
		input.Title,
		input.Description,
		input.OrganizerID,
		input.StartTime,
	)
	if err != nil {
		return nil, err
	}

	if input.MaxParticipants > 0 {
		meeting.MaxParticipants = input.MaxParticipants
	}

	if err := uc.meetingRepo.Create(ctx, meeting); err != nil {
		return nil, err
	}

	return meeting, nil
}
