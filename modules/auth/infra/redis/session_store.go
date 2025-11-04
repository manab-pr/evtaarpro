package redis

import (
	"context"
	"time"

	"github.com/manab-pr/evtaarpro/internal/datastore"
)

// SessionStore implements ports.SessionStore using Redis
type SessionStore struct {
	redis *datastore.RedisStore
	ttl   time.Duration
}

// NewSessionStore creates a new SessionStore
func NewSessionStore(redis *datastore.RedisStore, ttl time.Duration) *SessionStore {
	return &SessionStore{
		redis: redis,
		ttl:   ttl,
	}
}

// Create creates a new session
func (s *SessionStore) Create(ctx context.Context, userID string, token string, ttl time.Duration) error {
	key := s.redis.GetKey("session", userID)
	if ttl == 0 {
		ttl = s.ttl
	}
	return s.redis.Set(ctx, key, token, ttl)
}

// Get retrieves a session
func (s *SessionStore) Get(ctx context.Context, userID string) (string, error) {
	key := s.redis.GetKey("session", userID)
	return s.redis.Get(ctx, key)
}

// Delete deletes a session
func (s *SessionStore) Delete(ctx context.Context, userID string) error {
	key := s.redis.GetKey("session", userID)
	return s.redis.Delete(ctx, key)
}

// Exists checks if a session exists
func (s *SessionStore) Exists(ctx context.Context, userID string) (bool, error) {
	key := s.redis.GetKey("session", userID)
	count, err := s.redis.Exists(ctx, key)
	return count > 0, err
}
