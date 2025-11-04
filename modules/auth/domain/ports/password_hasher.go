package ports

// PasswordHasher defines methods for password hashing
type PasswordHasher interface {
	// Hash hashes a plain password
	Hash(password string) (string, error)

	// Compare compares a plain password with a hash
	Compare(hashedPassword, password string) error
}
