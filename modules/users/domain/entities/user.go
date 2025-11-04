package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserData = errors.New("invalid user data")
)

// User represents a user profile entity
type User struct {
	ID            string
	Email         string
	FirstName     string
	LastName      string
	Phone         string
	Avatar        string
	Role          string
	Department    string
	IsActive      bool
	EmailVerified bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// NewUser creates a new user entity
func NewUser(email, firstName, lastName, role string) (*User, error) {
	if email == "" || firstName == "" || lastName == "" {
		return nil, ErrInvalidUserData
	}

	now := time.Now()
	return &User{
		ID:            uuid.New().String(),
		Email:         email,
		FirstName:     firstName,
		LastName:      lastName,
		Role:          role,
		IsActive:      true,
		EmailVerified: false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// Update updates user information
func (u *User) Update(firstName, lastName, phone, department string) {
	if firstName != "" {
		u.FirstName = firstName
	}
	if lastName != "" {
		u.LastName = lastName
	}
	if phone != "" {
		u.Phone = phone
	}
	if department != "" {
		u.Department = department
	}
	u.UpdatedAt = time.Now()
}

// SetAvatar sets the user's avatar URL
func (u *User) SetAvatar(avatarURL string) {
	u.Avatar = avatarURL
	u.UpdatedAt = time.Now()
}
