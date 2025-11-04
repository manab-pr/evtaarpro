package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/manab-pr/evtaarpro/internal/config"
)

// PostgresStore wraps a PostgreSQL database connection
type PostgresStore struct {
	DB *sql.DB
}

// NewPostgresStore creates a new PostgreSQL connection
func NewPostgresStore(cfg *config.PostgresConfig) (*PostgresStore, error) {
	dsn := cfg.GetDSN()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(cfg.Pool.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Pool.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.Pool.ConnMaxLifetime)
	db.SetConnMaxIdleTime(cfg.Pool.ConnMaxIdleTime)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresStore{DB: db}, nil
}

// Close closes the database connection
func (s *PostgresStore) Close() error {
	if s.DB != nil {
		return s.DB.Close()
	}
	return nil
}

// Health checks the database health
func (s *PostgresStore) Health(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := s.DB.PingContext(ctx); err != nil {
		return fmt.Errorf("database unhealthy: %w", err)
	}

	return nil
}

// Stats returns database statistics
func (s *PostgresStore) Stats() sql.DBStats {
	return s.DB.Stats()
}

// BeginTx starts a new transaction
func (s *PostgresStore) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return s.DB.BeginTx(ctx, opts)
}

// WithTransaction executes a function within a database transaction
func (s *PostgresStore) WithTransaction(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := s.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
