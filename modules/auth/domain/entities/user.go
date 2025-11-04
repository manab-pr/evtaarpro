package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail    = errors.New("invalid email address")
	ErrInvalidPassword = errors.New("password must be at least 8 characters")
	ErrInvalidRole     = errors.New("invalid user role")
)

// Role represents user roles
type Role string

const (
	RoleAdmin    Role = "admin"
	RoleEmployee Role = "employee"
	RoleClient   Role = "client"
	RoleHR       Role = "hr"
)

// User represents a user entity
type User struct {
	ID           string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Role         Role
	IsActive     bool
	EmailVerified bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser creates a new user entity
func NewUser(email, passwordHash, firstName, lastName string, role Role) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}

	if !IsValidRole(role) {
		return nil, ErrInvalidRole
	}

	now := time.Now()
	return &User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
		Role:         role,
		IsActive:     true,
		EmailVerified: false,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// Activate activates the user account
func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

// Deactivate deactivates the user account
func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

// VerifyEmail marks email as verified
func (u *User) VerifyEmail() {
	u.EmailVerified = true
	u.UpdatedAt = time.Now()
}

// ChangeRole changes the user's role
func (u *User) ChangeRole(role Role) error {
	if !IsValidRole(role) {
		return ErrInvalidRole
	}
	u.Role = role
	u.UpdatedAt = time.Now()
	return nil
}

// IsValidRole checks if a role is valid
func IsValidRole(role Role) bool {
	switch role {
	case RoleAdmin, RoleEmployee, RoleClient, RoleHR:
		return true
	default:
		return false
	}
}
