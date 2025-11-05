package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type NotificationType string
type NotificationPriority string

const (
	TypeMeeting  NotificationType = "meeting"
	TypePayroll  NotificationType = "payroll"
	TypeCRM      NotificationType = "crm"
	TypeSystem   NotificationType = "system"

	PriorityLow    NotificationPriority = "low"
	PriorityMedium NotificationPriority = "medium"
	PriorityHigh   NotificationPriority = "high"
	PriorityUrgent NotificationPriority = "urgent"
)

type Notification struct {
	ID        string
	UserID    string
	Type      NotificationType
	Title     string
	Message   string
	Data      map[string]interface{}
	IsRead    bool
	ReadAt    *time.Time
	Priority  NotificationPriority
	CreatedAt time.Time
}

func NewNotification(userID string, notifType NotificationType, title, message string, priority NotificationPriority) (*Notification, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	if title == "" {
		return nil, errors.New("title is required")
	}
	if message == "" {
		return nil, errors.New("message is required")
	}

	return &Notification{
		ID:        uuid.New().String(),
		UserID:    userID,
		Type:      notifType,
		Title:     title,
		Message:   message,
		IsRead:    false,
		Priority:  priority,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Notification) MarkAsRead() {
	if !n.IsRead {
		n.IsRead = true
		readTime := time.Now()
		n.ReadAt = &readTime
	}
}

func (n *Notification) MarkAsUnread() {
	n.IsRead = false
	n.ReadAt = nil
}
