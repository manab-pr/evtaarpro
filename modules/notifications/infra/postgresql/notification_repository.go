package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/manab-pr/evtaarpro/modules/notifications/domain/entities"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(ctx context.Context, notification *entities.Notification) error {
	var dataJSON []byte
	var err error
	if notification.Data != nil {
		dataJSON, err = json.Marshal(notification.Data)
		if err != nil {
			return err
		}
	}

	query := `
		INSERT INTO notifications (id, user_id, type, title, message, data, is_read, read_at, priority, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err = r.db.ExecContext(ctx, query,
		notification.ID, notification.UserID, notification.Type, notification.Title,
		notification.Message, dataJSON, notification.IsRead, notification.ReadAt,
		notification.Priority, notification.CreatedAt,
	)
	return err
}

func (r *NotificationRepository) GetByID(ctx context.Context, id string) (*entities.Notification, error) {
	query := `
		SELECT id, user_id, type, title, message, data, is_read, read_at, priority, created_at
		FROM notifications WHERE id = $1
	`
	notification := &entities.Notification{}
	var dataJSON []byte
	var readAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&notification.ID, &notification.UserID, &notification.Type, &notification.Title,
		&notification.Message, &dataJSON, &notification.IsRead, &readAt,
		&notification.Priority, &notification.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("notification not found")
	}
	if err != nil {
		return nil, err
	}

	if dataJSON != nil {
		if err := json.Unmarshal(dataJSON, &notification.Data); err != nil {
			return nil, err
		}
	}

	if readAt.Valid {
		notification.ReadAt = &readAt.Time
	}

	return notification, nil
}

func (r *NotificationRepository) ListByUser(ctx context.Context, userID string, isRead *bool, notifType entities.NotificationType, offset, limit int) ([]*entities.Notification, int, error) {
	var countQuery string
	var query string
	var args []interface{}

	baseCondition := `WHERE user_id = $1`
	args = append(args, userID)
	argPos := 2

	if isRead != nil {
		baseCondition += ` AND is_read = $` + string(rune(argPos+'0'))
		args = append(args, *isRead)
		argPos++
	}

	if notifType != "" {
		baseCondition += ` AND type = $` + string(rune(argPos+'0'))
		args = append(args, notifType)
		argPos++
	}

	countQuery = `SELECT COUNT(*) FROM notifications ` + baseCondition

	query = `
		SELECT id, user_id, type, title, message, data, is_read, read_at, priority, created_at
		FROM notifications ` + baseCondition + `
		ORDER BY created_at DESC
		LIMIT $` + string(rune(argPos+'0')) + ` OFFSET $` + string(rune(argPos+1+'0'))

	countArgs := args
	args = append(args, limit, offset)

	var total int
	err := r.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	notifications := []*entities.Notification{}
	for rows.Next() {
		notification := &entities.Notification{}
		var dataJSON []byte
		var readAt sql.NullTime

		err := rows.Scan(
			&notification.ID, &notification.UserID, &notification.Type, &notification.Title,
			&notification.Message, &dataJSON, &notification.IsRead, &readAt,
			&notification.Priority, &notification.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if dataJSON != nil {
			if err := json.Unmarshal(dataJSON, &notification.Data); err != nil {
				return nil, 0, err
			}
		}

		if readAt.Valid {
			notification.ReadAt = &readAt.Time
		}

		notifications = append(notifications, notification)
	}

	return notifications, total, nil
}

func (r *NotificationRepository) Update(ctx context.Context, notification *entities.Notification) error {
	var dataJSON []byte
	var err error
	if notification.Data != nil {
		dataJSON, err = json.Marshal(notification.Data)
		if err != nil {
			return err
		}
	}

	query := `
		UPDATE notifications SET
			is_read = $2, read_at = $3, data = $4
		WHERE id = $1
	`
	_, err = r.db.ExecContext(ctx, query,
		notification.ID, notification.IsRead, notification.ReadAt, dataJSON,
	)
	return err
}

func (r *NotificationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM notifications WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *NotificationRepository) MarkAllAsRead(ctx context.Context, userID string) error {
	query := `UPDATE notifications SET is_read = true, read_at = $1 WHERE user_id = $2 AND is_read = false`
	_, err := r.db.ExecContext(ctx, query, time.Now(), userID)
	return err
}

func (r *NotificationRepository) GetUnreadCount(ctx context.Context, userID string) (int, error) {
	query := `SELECT COUNT(*) FROM notifications WHERE user_id = $1 AND is_read = false`
	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	return count, err
}
