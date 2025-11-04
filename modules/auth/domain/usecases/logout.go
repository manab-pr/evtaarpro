package usecases

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/auth/domain/ports"
)

// LogoutUseCase handles user logout
type LogoutUseCase struct {
	sessionStore ports.SessionStore
}

// NewLogoutUseCase creates a new LogoutUseCase
func NewLogoutUseCase(sessionStore ports.SessionStore) *LogoutUseCase {
	return &LogoutUseCase{
		sessionStore: sessionStore,
	}
}

// Execute executes the logout use case
func (uc *LogoutUseCase) Execute(ctx context.Context, userID string) error {
	return uc.sessionStore.Delete(ctx, userID)
}
