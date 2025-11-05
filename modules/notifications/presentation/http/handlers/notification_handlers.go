package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/modules/notifications/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/notifications/domain/repository"
	"github.com/manab-pr/evtaarpro/modules/notifications/presentation/http/dto"
)

type NotificationHandlers struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationHandlers(notificationRepo repository.NotificationRepository) *NotificationHandlers {
	return &NotificationHandlers{
		notificationRepo: notificationRepo,
	}
}

func (h *NotificationHandlers) CreateNotification(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	priority := entities.PriorityMedium
	if req.Priority != "" {
		priority = entities.NotificationPriority(req.Priority)
	}

	notification, err := entities.NewNotification(
		req.UserID,
		entities.NotificationType(req.Type),
		req.Title,
		req.Message,
		priority,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification.Data = req.Data

	if err := h.notificationRepo.Create(c.Request.Context(), notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	c.JSON(http.StatusCreated, mapNotificationToResponse(notification))
}

func (h *NotificationHandlers) GetNotification(c *gin.Context) {
	id := c.Param("id")
	notification, err := h.notificationRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	c.JSON(http.StatusOK, mapNotificationToResponse(notification))
}

func (h *NotificationHandlers) ListMyNotifications(c *gin.Context) {
	userID, _ := c.Get("userID")
	notifType := c.Query("type")
	isReadStr := c.Query("is_read")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit

	var isRead *bool
	if isReadStr != "" {
		val := isReadStr == "true"
		isRead = &val
	}

	var notificationT entities.NotificationType
	if notifType != "" {
		notificationT = entities.NotificationType(notifType)
	}

	notifications, total, err := h.notificationRepo.ListByUser(
		c.Request.Context(),
		userID.(string),
		isRead,
		notificationT,
		offset,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	response := make([]dto.NotificationResponse, len(notifications))
	for i, notif := range notifications {
		response[i] = mapNotificationToResponse(notif)
	}

	c.JSON(http.StatusOK, dto.ListResponse{
		Data:  response,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

func (h *NotificationHandlers) MarkAsRead(c *gin.Context) {
	id := c.Param("id")
	notification, err := h.notificationRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	notification.MarkAsRead()

	if err := h.notificationRepo.Update(c.Request.Context(), notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as read"})
		return
	}

	c.JSON(http.StatusOK, mapNotificationToResponse(notification))
}

func (h *NotificationHandlers) MarkAsUnread(c *gin.Context) {
	id := c.Param("id")
	notification, err := h.notificationRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notification not found"})
		return
	}

	notification.MarkAsUnread()

	if err := h.notificationRepo.Update(c.Request.Context(), notification); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark notification as unread"})
		return
	}

	c.JSON(http.StatusOK, mapNotificationToResponse(notification))
}

func (h *NotificationHandlers) MarkAllAsRead(c *gin.Context) {
	userID, _ := c.Get("userID")

	if err := h.notificationRepo.MarkAllAsRead(c.Request.Context(), userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark all notifications as read"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All notifications marked as read"})
}

func (h *NotificationHandlers) GetUnreadCount(c *gin.Context) {
	userID, _ := c.Get("userID")

	count, err := h.notificationRepo.GetUnreadCount(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get unread count"})
		return
	}

	c.JSON(http.StatusOK, dto.UnreadCountResponse{Count: count})
}

func (h *NotificationHandlers) DeleteNotification(c *gin.Context) {
	id := c.Param("id")

	if err := h.notificationRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete notification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notification deleted successfully"})
}

func mapNotificationToResponse(notif *entities.Notification) dto.NotificationResponse {
	return dto.NotificationResponse{
		ID:        notif.ID,
		UserID:    notif.UserID,
		Type:      string(notif.Type),
		Title:     notif.Title,
		Message:   notif.Message,
		Data:      notif.Data,
		IsRead:    notif.IsRead,
		ReadAt:    notif.ReadAt,
		Priority:  string(notif.Priority),
		CreatedAt: notif.CreatedAt,
	}
}
