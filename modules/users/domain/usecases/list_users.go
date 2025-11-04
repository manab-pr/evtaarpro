package usecases

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/users/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/users/domain/repository"
)

// ListUsersUseCase handles listing users
type ListUsersUseCase struct {
	userRepo repository.UserRepository
}

// NewListUsersUseCase creates a new ListUsersUseCase
func NewListUsersUseCase(userRepo repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{userRepo: userRepo}
}

// Execute lists users with pagination
func (uc *ListUsersUseCase) Execute(ctx context.Context, page, pageSize int) ([]*entities.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	return uc.userRepo.List(ctx, page, pageSize)
}
