package ports

import (
	"context"
	"github.com/manab-pr/evtaarpro/modules/notifications/domain/entities"
)

// NotificationRepository defines notification persistence operations
type NotificationRepository interface {
	Create(ctx context.Context, notification *entities.Notification) error
	GetByID(ctx context.Context, id string) (*entities.Notification, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*entities.Notification, int, error)
	MarkAsRead(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	Delete(ctx context.Context, id string) error
	GetUnreadCount(ctx context.Context, userID string) (int, error)
}
