package usecases

import (
	"context"
	"errors"
	"strings"

	"github.com/manab-pr/evtaarpro/modules/auth/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/auth/domain/ports"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrRegistrationFailed = errors.New("registration failed")
)

// RegisterUseCase handles user registration
type RegisterUseCase struct {
	userRepo       ports.UserRepository
	passwordHasher ports.PasswordHasher
}

// NewRegisterUseCase creates a new RegisterUseCase
func NewRegisterUseCase(userRepo ports.UserRepository, passwordHasher ports.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

// RegisterInput represents registration input
type RegisterInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      entities.Role
}

// RegisterOutput represents registration output
type RegisterOutput struct {
	UserID string
	Email  string
}

// Execute executes the register use case
func (uc *RegisterUseCase) Execute(ctx context.Context, input RegisterInput) (*RegisterOutput, error) {
	// Check if user already exists
	exists, err := uc.userRepo.Exists(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := uc.passwordHasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	// Create user entity
	user, err := entities.NewUser(
		input.Email,
		hashedPassword,
		input.FirstName,
		input.LastName,
		input.Role,
	)
	if err != nil {
		return nil, err
	}

	// Save user
	if err := uc.userRepo.Create(ctx, user); err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, ErrUserAlreadyExists
		}
		return nil, ErrRegistrationFailed
	}

	return &RegisterOutput{
		UserID: user.ID,
		Email:  user.Email,
	}, nil
}
