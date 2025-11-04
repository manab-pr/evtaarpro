package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/users/domain/repository"
)

var ErrUserNotFound = errors.New("user not found")

// UpdateUserUseCase handles user profile updates
type UpdateUserUseCase struct {
	userRepo repository.UserRepository
}

// NewUpdateUserUseCase creates a new UpdateUserUseCase
func NewUpdateUserUseCase(userRepo repository.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: userRepo}
}

// UpdateInput represents update input
type UpdateInput struct {
	UserID     string
	FirstName  string
	LastName   string
	Phone      string
	Department string
}

// Execute updates user information
func (uc *UpdateUserUseCase) Execute(ctx context.Context, input UpdateInput) error {
	user, err := uc.userRepo.GetByID(ctx, input.UserID)
	if err != nil {
		return ErrUserNotFound
	}

	user.Update(input.FirstName, input.LastName, input.Phone, input.Department)

	return uc.userRepo.Update(ctx, user)
}
