package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/manab-pr/evtaarpro/modules/notifications/domain/entities"
)

// NotificationRepository implements notification persistence
type NotificationRepository struct {
	db *sql.DB
}

// NewNotificationRepository creates a new repository
func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// Create creates a new notification
func (r *NotificationRepository) Create(ctx context.Context, notification *entities.Notification) error {
	query := `
		INSERT INTO notifications (id, user_id, type, title, message, data, read, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		notification.ID, notification.UserID, notification.Type, notification.Title,
		notification.Message, notification.Data, notification.Read, notification.CreatedAt,
	)
	return err
}

// GetByID retrieves a notification by ID
func (r *NotificationRepository) GetByID(ctx context.Context, id string) (*entities.Notification, error) {
	query := `
		SELECT id, user_id, type, title, message, data, read, created_at, read_at
		FROM notifications WHERE id = $1
	`
	notification := &entities.Notification{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&notification.ID, &notification.UserID, &notification.Type, &notification.Title,
		&notification.Message, &notification.Data, &notification.Read,
		&notification.CreatedAt, &notification.ReadAt,
	)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

// ListByUser retrieves notifications for a user
func (r *NotificationRepository) ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entities.Notification, int, error) {
	query := `
		SELECT id, user_id, type, title, message, data, read, created_at, read_at
		FROM notifications
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var notifications []*entities.Notification
	for rows.Next() {
		notification := &entities.Notification{}
		err := rows.Scan(
			&notification.ID, &notification.UserID, &notification.Type, &notification.Title,
			&notification.Message, &notification.Data, &notification.Read,
			&notification.CreatedAt, &notification.ReadAt,
		)
		if err != nil {
			return nil, 0, err
		}
		notifications = append(notifications, notification)
	}

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM notifications WHERE user_id = $1`
	err = r.db.QueryRowContext(ctx, countQuery, userID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}

// MarkAsRead marks a notification as read
func (r *NotificationRepository) MarkAsRead(ctx context.Context, id string) error {
	query := `
		UPDATE notifications
		SET read = true, read_at = $2
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id, time.Now())
	return err
}

// MarkAllAsRead marks all notifications for a user as read
func (r *NotificationRepository) MarkAllAsRead(ctx context.Context, userID string) error {
	query := `
		UPDATE notifications
		SET read = true, read_at = $2
		WHERE user_id = $1 AND read = false
	`
	_, err := r.db.ExecContext(ctx, query, userID, time.Now())
	return err
}

// Delete deletes a notification
func (r *NotificationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM notifications WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GetUnreadCount gets the count of unread notifications for a user
func (r *NotificationRepository) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND read = false`
	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
