package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/meetings/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/meetings/presentation/http/dto"
)

// MeetingHandlers contains meeting-related HTTP handlers
type MeetingHandlers struct {
	createMeetingUC *usecases.CreateMeetingUseCase
	getMeetingUC    *usecases.GetMeetingUseCase
	listMeetingsUC  *usecases.ListMeetingsUseCase
	joinMeetingUC   *usecases.JoinMeetingUseCase
}

// NewMeetingHandlers creates new MeetingHandlers
func NewMeetingHandlers(
	createMeetingUC *usecases.CreateMeetingUseCase,
	getMeetingUC *usecases.GetMeetingUseCase,
	listMeetingsUC *usecases.ListMeetingsUseCase,
	joinMeetingUC *usecases.JoinMeetingUseCase,
) *MeetingHandlers {
	return &MeetingHandlers{
		createMeetingUC: createMeetingUC,
		getMeetingUC:    getMeetingUC,
		listMeetingsUC:  listMeetingsUC,
		joinMeetingUC:   joinMeetingUC,
	}
}

// CreateMeeting creates a new meeting
// @Summary Create meeting
// @Description Create a new meeting
// @Tags meetings
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateMeetingRequest true "Meeting details"
// @Success 201 {object} response.Response{data=dto.MeetingResponse}
// @Router /meetings [post]
func (h *MeetingHandlers) CreateMeeting(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req dto.CreateMeetingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	meeting, err := h.createMeetingUC.Execute(c.Request.Context(), usecases.CreateInput{
		Title:           req.Title,
		Description:     req.Description,
		OrganizerID:     userID.(string),
		StartTime:       req.StartTime,
		MaxParticipants: req.MaxParticipants,
	})

	if err != nil {
		response.InternalServerError(c, "Failed to create meeting")
		return
	}

	response.Created(c, "Meeting created successfully", mapMeetingToResponse(meeting))
}

// GetMeeting retrieves a meeting
// @Summary Get meeting
// @Description Get meeting by ID
// @Tags meetings
// @Security BearerAuth
// @Produce json
// @Param id path string true "Meeting ID"
// @Success 200 {object} response.Response{data=dto.MeetingResponse}
// @Router /meetings/{id} [get]
func (h *MeetingHandlers) GetMeeting(c *gin.Context) {
	meetingID := c.Param("id")

	meeting, err := h.getMeetingUC.Execute(c.Request.Context(), meetingID)
	if err != nil {
		response.NotFound(c, "Meeting not found")
		return
	}

	response.OK(c, "Meeting retrieved successfully", mapMeetingToResponse(meeting))
}

// ListMeetings lists meetings
// @Summary List meetings
// @Description List meetings with pagination
// @Tags meetings
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page" default(1)
// @Param page_size query int false "Page size" default(20)
// @Success 200 {object} response.PaginatedResponse
// @Router /meetings [get]
func (h *MeetingHandlers) ListMeetings(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	meetings, total, err := h.listMeetingsUC.Execute(c.Request.Context(), userID.(string), page, pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to list meetings")
		return
	}

	meetingResponses := make([]dto.MeetingResponse, len(meetings))
	for i, meeting := range meetings {
		meetingResponses[i] = mapMeetingToResponse(meeting)
	}

	response.Paginated(c, meetingResponses, page, pageSize, total)
}

// JoinMeeting joins a meeting
// @Summary Join meeting
// @Description Join a meeting and get Jitsi token
// @Tags meetings
// @Security BearerAuth
// @Produce json
// @Param id path string true "Meeting ID"
// @Success 200 {object} response.Response{data=dto.JoinMeetingResponse}
// @Router /meetings/{id}/join [post]
func (h *MeetingHandlers) JoinMeeting(c *gin.Context) {
	meetingID := c.Param("id")
	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")

	// Get user name from context or use email
	userName := email.(string)

	output, err := h.joinMeetingUC.Execute(c.Request.Context(), meetingID, userID.(string), userName, email.(string))
	if err != nil {
		response.InternalServerError(c, "Failed to join meeting")
		return
	}

	response.OK(c, "Joined meeting successfully", dto.JoinMeetingResponse{
		MeetingID: output.MeetingID,
		RoomURL:   output.RoomURL,
		UserName:  output.UserName,
		UserEmail: output.UserEmail,
	})
}

func mapMeetingToResponse(meeting *entities.Meeting) dto.MeetingResponse {
	return dto.MeetingResponse{
		ID:              meeting.ID,
		RoomID:          meeting.RoomID,
		Title:           meeting.Title,
		Description:     meeting.Description,
		OrganizerID:     meeting.OrganizerID,
		StartTime:       meeting.StartTime,
		EndTime:         meeting.EndTime,
		Status:          string(meeting.Status),
		JitsiRoomURL:    meeting.JitsiRoomURL,
		RecordingURL:    meeting.RecordingURL,
		MaxParticipants: meeting.MaxParticipants,
		CreatedAt:       meeting.CreatedAt,
	}
}
