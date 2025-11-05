package repository

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/notifications/domain/entities"
)

type NotificationRepository interface {
	Create(ctx context.Context, notification *entities.Notification) error
	GetByID(ctx context.Context, id string) (*entities.Notification, error)
	ListByUser(ctx context.Context, userID string, isRead *bool, notifType entities.NotificationType, offset, limit int) ([]*entities.Notification, int, error)
	Update(ctx context.Context, notification *entities.Notification) error
	Delete(ctx context.Context, id string) error
	MarkAllAsRead(ctx context.Context, userID string) error
	GetUnreadCount(ctx context.Context, userID string) (int, error)
}
