package meetings

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/meetings/infra/jitsi"
	"github.com/manab-pr/evtaarpro/modules/meetings/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/meetings/presentation/http/handlers"
	"github.com/manab-pr/evtaarpro/modules/meetings/presentation/http/routes"
)

// RegisterRoutes registers meetings module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Infrastructure
	meetingRepo := postgresql.NewMeetingRepository(pgStore.DB)
	jitsiAdapter := jitsi.NewJitsiAdapter(cfg.Jitsi.Domain, cfg.Jitsi.AppID, cfg.Jitsi.AppSecret)

	// Create a wrapper repository for join use case
	joinRepo := &joinMeetingRepoAdapter{repo: meetingRepo}

	// Use cases
	createMeetingUC := usecases.NewCreateMeetingUseCase(meetingRepo)
	getMeetingUC := usecases.NewGetMeetingUseCase(meetingRepo)
	listMeetingsUC := usecases.NewListMeetingsUseCase(meetingRepo)
	joinMeetingUC := usecases.NewJoinMeetingUseCase(joinRepo, jitsiAdapter)

	// Handlers
	meetingHandlers := handlers.NewMeetingHandlers(createMeetingUC, getMeetingUC, listMeetingsUC, joinMeetingUC)

	// Register routes
	routes.RegisterRoutes(rg, meetingHandlers, cfg.JWT.Secret)
}

// Adapter to bridge the meeting repository with join use case interface
type joinMeetingRepoAdapter struct {
	repo *postgresql.MeetingRepository
}

func (a *joinMeetingRepoAdapter) GetByID(ctx context.Context, id string) (*usecases.MeetingEntity, error) {
	meeting, err := a.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &usecases.MeetingEntity{
		ID:           meeting.ID,
		RoomID:       meeting.RoomID,
		OrganizerID:  meeting.OrganizerID,
		Status:       string(meeting.Status),
		JitsiRoomURL: meeting.JitsiRoomURL,
	}, nil
}

func (a *joinMeetingRepoAdapter) Update(ctx context.Context, entity *usecases.MeetingEntity) error {
	meeting, err := a.repo.GetByID(ctx, entity.ID)
	if err != nil {
		return err
	}

	meeting.Status = entities.MeetingStatus(entity.Status)
	meeting.JitsiRoomURL = entity.JitsiRoomURL

	return a.repo.Update(ctx, meeting)
}
