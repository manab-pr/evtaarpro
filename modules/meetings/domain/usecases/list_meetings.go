package usecases

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/repository"
)

// ListMeetingsUseCase handles listing meetings
type ListMeetingsUseCase struct {
	meetingRepo repository.MeetingRepository
}

// NewListMeetingsUseCase creates a new ListMeetingsUseCase
func NewListMeetingsUseCase(meetingRepo repository.MeetingRepository) *ListMeetingsUseCase {
	return &ListMeetingsUseCase{meetingRepo: meetingRepo}
}

// Execute lists meetings with pagination
func (uc *ListMeetingsUseCase) Execute(ctx context.Context, userID string, page, pageSize int) ([]*entities.Meeting, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return uc.meetingRepo.List(ctx, page, pageSize, userID)
}
