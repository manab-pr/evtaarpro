package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/notifications/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/notifications/domain/ports"
	"github.com/manab-pr/evtaarpro/modules/notifications/presentation/http/dto"
)

// NotificationHandlers contains notification-related HTTP handlers
type NotificationHandlers struct {
	notificationRepo ports.NotificationRepository
}

// NewNotificationHandlers creates new NotificationHandlers
func NewNotificationHandlers(notificationRepo ports.NotificationRepository) *NotificationHandlers {
	return &NotificationHandlers{
		notificationRepo: notificationRepo,
	}
}

// CreateNotification creates a new notification
func (h *NotificationHandlers) CreateNotification(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	notification := &entities.Notification{
		ID:        uuid.New().String(),
		UserID:    req.UserID,
		Type:      req.Type,
		Title:     req.Title,
		Message:   req.Message,
		Data:      req.Data,
		Read:      false,
		CreatedAt: time.Now(),
	}

	if err := h.notificationRepo.Create(c.Request.Context(), notification); err != nil {
		response.InternalServerError(c, "Failed to create notification")
		return
	}

	response.Created(c, "Notification created successfully", mapNotificationToResponse(notification))
}

// ListNotifications lists notifications for the logged-in user
func (h *NotificationHandlers) ListNotifications(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	notifications, total, err := h.notificationRepo.ListByUser(c.Request.Context(), userID.(string), pageSize, (page-1)*pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to list notifications")
		return
	}

	notificationResponses := make([]dto.NotificationResponse, len(notifications))
	for i, notification := range notifications {
		notificationResponses[i] = mapNotificationToResponse(notification)
	}

	response.Paginated(c, notificationResponses, page, pageSize, int64(total))
}

// GetUnreadCount gets the count of unread notifications
func (h *NotificationHandlers) GetUnreadCount(c *gin.Context) {
	userID, _ := c.Get("user_id")

	count, err := h.notificationRepo.GetUnreadCount(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalServerError(c, "Failed to get unread count")
		return
	}

	response.OK(c, "Unread count retrieved successfully", gin.H{"count": count})
}

// MarkAsRead marks a notification as read
func (h *NotificationHandlers) MarkAsRead(c *gin.Context) {
	notificationID := c.Param("id")

	if err := h.notificationRepo.MarkAsRead(c.Request.Context(), notificationID); err != nil {
		response.InternalServerError(c, "Failed to mark notification as read")
		return
	}

	response.OK(c, "Notification marked as read", nil)
}

// MarkAllAsRead marks all notifications for the user as read
func (h *NotificationHandlers) MarkAllAsRead(c *gin.Context) {
	userID, _ := c.Get("user_id")

	if err := h.notificationRepo.MarkAllAsRead(c.Request.Context(), userID.(string)); err != nil {
		response.InternalServerError(c, "Failed to mark all notifications as read")
		return
	}

	response.OK(c, "All notifications marked as read", nil)
}

// DeleteNotification deletes a notification
func (h *NotificationHandlers) DeleteNotification(c *gin.Context) {
	notificationID := c.Param("id")

	if err := h.notificationRepo.Delete(c.Request.Context(), notificationID); err != nil {
		response.InternalServerError(c, "Failed to delete notification")
		return
	}

	response.OK(c, "Notification deleted successfully", nil)
}

func mapNotificationToResponse(notification *entities.Notification) dto.NotificationResponse {
	return dto.NotificationResponse{
		ID:        notification.ID,
		UserID:    notification.UserID,
		Type:      notification.Type,
		Title:     notification.Title,
		Message:   notification.Message,
		Data:      notification.Data,
		Read:      notification.Read,
		CreatedAt: notification.CreatedAt,
		ReadAt:    notification.ReadAt,
	}
}
