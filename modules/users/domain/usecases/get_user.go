package usecases

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/users/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/users/domain/repository"
)

// GetUserUseCase handles retrieving a user
type GetUserUseCase struct {
	userRepo repository.UserRepository
}

// NewGetUserUseCase creates a new GetUserUseCase
func NewGetUserUseCase(userRepo repository.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
}

// Execute retrieves a user by ID
func (uc *GetUserUseCase) Execute(ctx context.Context, userID string) (*entities.User, error) {
	return uc.userRepo.GetByID(ctx, userID)
}
