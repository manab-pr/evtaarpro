package security

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHasher implements password hashing using bcrypt
type BcryptHasher struct {
	cost int
}

// NewBcryptHasher creates a new BcryptHasher
func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{
		cost: bcrypt.DefaultCost,
	}
}

// Hash hashes a plain password
func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compare compares a plain password with a hash
func (h *BcryptHasher) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
