package datastore

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/manab-pr/evtaarpro/internal/config"
)

// RedisStore wraps a Redis client
type RedisStore struct {
	Client   *redis.Client
	Prefixes map[string]string
	TTL      map[string]int
}

// NewRedisStore creates a new Redis connection
func NewRedisStore(cfg *config.RedisConfig) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.GetAddr(),
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  cfg.Timeouts.Dial,
		ReadTimeout:  cfg.Timeouts.Read,
		WriteTimeout: cfg.Timeouts.Write,
		PoolSize:     cfg.Pool.MaxActive,
		MinIdleConns: cfg.Pool.MaxIdle,
		PoolTimeout:  cfg.Pool.IdleTimeout,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &RedisStore{
		Client:   client,
		Prefixes: cfg.Prefixes,
		TTL:      cfg.TTL,
	}, nil
}

// Close closes the Redis connection
func (s *RedisStore) Close() error {
	if s.Client != nil {
		return s.Client.Close()
	}
	return nil
}

// Health checks the Redis health
func (s *RedisStore) Health(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := s.Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis unhealthy: %w", err)
	}

	return nil
}

// GetKey returns a prefixed key
func (s *RedisStore) GetKey(prefix, key string) string {
	if p, ok := s.Prefixes[prefix]; ok {
		return p + key
	}
	return key
}

// GetTTL returns TTL for a given prefix
func (s *RedisStore) GetTTL(prefix string) time.Duration {
	if ttl, ok := s.TTL[prefix]; ok {
		return time.Duration(ttl) * time.Second
	}
	return 0
}

// Set sets a key-value pair with optional TTL
func (s *RedisStore) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.Client.Set(ctx, key, value, expiration).Err()
}

// Get gets a value by key
func (s *RedisStore) Get(ctx context.Context, key string) (string, error) {
	return s.Client.Get(ctx, key).Result()
}

// Delete deletes a key
func (s *RedisStore) Delete(ctx context.Context, keys ...string) error {
	return s.Client.Del(ctx, keys...).Err()
}

// Exists checks if a key exists
func (s *RedisStore) Exists(ctx context.Context, keys ...string) (int64, error) {
	return s.Client.Exists(ctx, keys...).Result()
}

// Expire sets TTL on a key
func (s *RedisStore) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return s.Client.Expire(ctx, key, expiration).Err()
}

// HSet sets a hash field
func (s *RedisStore) HSet(ctx context.Context, key string, values ...interface{}) error {
	return s.Client.HSet(ctx, key, values...).Err()
}

// HGet gets a hash field
func (s *RedisStore) HGet(ctx context.Context, key, field string) (string, error) {
	return s.Client.HGet(ctx, key, field).Result()
}

// HGetAll gets all hash fields
func (s *RedisStore) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return s.Client.HGetAll(ctx, key).Result()
}

// HDel deletes hash fields
func (s *RedisStore) HDel(ctx context.Context, key string, fields ...string) error {
	return s.Client.HDel(ctx, key, fields...).Err()
}
