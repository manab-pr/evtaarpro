package usecases

import (
	"context"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/auth/domain/ports"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrUserNotActive      = errors.New("user account is not active")
)

// LoginUseCase handles user login
type LoginUseCase struct {
	userRepo       ports.UserRepository
	passwordHasher ports.PasswordHasher
	tokenGenerator ports.TokenGenerator
	sessionStore   ports.SessionStore
}

// NewLoginUseCase creates a new LoginUseCase
func NewLoginUseCase(
	userRepo ports.UserRepository,
	passwordHasher ports.PasswordHasher,
	tokenGenerator ports.TokenGenerator,
	sessionStore ports.SessionStore,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
		sessionStore:   sessionStore,
	}
}

// LoginInput represents login input
type LoginInput struct {
	Email    string
	Password string
}

// LoginOutput represents login output
type LoginOutput struct {
	UserID       string
	Email        string
	Role         string
	AccessToken  string
	RefreshToken string
}

// Execute executes the login use case
func (uc *LoginUseCase) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	// Get user by email
	user, err := uc.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Check if user is active
	if !user.IsActive {
		return nil, ErrUserNotActive
	}

	// Verify password
	if err := uc.passwordHasher.Compare(user.PasswordHash, input.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate tokens
	accessToken, err := uc.tokenGenerator.GenerateAccessToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	refreshToken, err := uc.tokenGenerator.GenerateRefreshToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		return nil, err
	}

	// Store session
	if err := uc.sessionStore.Create(ctx, user.ID, refreshToken, 0); err != nil {
		return nil, err
	}

	return &LoginOutput{
		UserID:       user.ID,
		Email:        user.Email,
		Role:         string(user.Role),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
