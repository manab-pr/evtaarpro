package usecases

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/repository"
)

// GetMeetingUseCase handles retrieving a meeting
type GetMeetingUseCase struct {
	meetingRepo repository.MeetingRepository
}

// NewGetMeetingUseCase creates a new GetMeetingUseCase
func NewGetMeetingUseCase(meetingRepo repository.MeetingRepository) *GetMeetingUseCase {
	return &GetMeetingUseCase{meetingRepo: meetingRepo}
}

// Execute retrieves a meeting by ID
func (uc *GetMeetingUseCase) Execute(ctx context.Context, meetingID string) (*entities.Meeting, error) {
	return uc.meetingRepo.GetByID(ctx, meetingID)
}
