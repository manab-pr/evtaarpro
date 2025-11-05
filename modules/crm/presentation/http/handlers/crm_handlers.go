package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/repository"
	"github.com/manab-pr/evtaarpro/modules/crm/presentation/http/dto"
)

type CRMHandlers struct {
	customerRepo    repository.CustomerRepository
	interactionRepo repository.CustomerInteractionRepository
}

func NewCRMHandlers(
	customerRepo repository.CustomerRepository,
	interactionRepo repository.CustomerInteractionRepository,
) *CRMHandlers {
	return &CRMHandlers{
		customerRepo:    customerRepo,
		interactionRepo: interactionRepo,
	}
}

// Customer Handlers
func (h *CRMHandlers) CreateCustomer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := entities.NewCustomer(
		req.CompanyName,
		req.ContactName,
		req.Industry,
		entities.CustomerStatus(req.Status),
		userID.(string),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.Email = req.Email
	customer.Phone = req.Phone
	customer.Address = req.Address
	customer.AssignedTo = req.AssignedTo

	if err := h.customerRepo.Create(c.Request.Context(), customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, mapCustomerToResponse(customer))
}

func (h *CRMHandlers) GetCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	customer, err := h.customerRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, mapCustomerToResponse(customer))
}

func (h *CRMHandlers) ListCustomers(c *gin.Context) {
	status := c.Query("status")
	assignedTo := c.Query("assigned_to")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var customerStatus entities.CustomerStatus
	if status != "" {
		customerStatus = entities.CustomerStatus(status)
	}

	customers, total, err := h.customerRepo.List(c.Request.Context(), customerStatus, assignedTo, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}

	response := make([]dto.CustomerResponse, len(customers))
	for i, cust := range customers {
		response[i] = mapCustomerToResponse(cust)
	}

	c.JSON(http.StatusOK, dto.ListResponse{
		Data:  response,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

func (h *CRMHandlers) UpdateCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	var req dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.customerRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	if req.CompanyName != "" {
		customer.CompanyName = req.CompanyName
	}
	if req.ContactName != "" {
		customer.ContactName = req.ContactName
	}
	if req.Industry != "" {
		customer.Industry = req.Industry
	}
	if req.Status != "" {
		customer.UpdateStatus(entities.CustomerStatus(req.Status))
	}
	customer.Email = req.Email
	customer.Phone = req.Phone
	customer.Address = req.Address
	customer.AssignedTo = req.AssignedTo
	customer.UpdatedAt = time.Now()

	if err := h.customerRepo.Update(c.Request.Context(), customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(http.StatusOK, mapCustomerToResponse(customer))
}

func (h *CRMHandlers) DeleteCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	if err := h.customerRepo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

// Interaction Handlers
func (h *CRMHandlers) CreateInteraction(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req dto.CreateInteractionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interaction, err := entities.NewCustomerInteraction(
		req.CustomerID,
		userID.(string),
		entities.InteractionType(req.InteractionType),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interaction.Subject = req.Subject
	interaction.Notes = req.Notes
	if req.InteractionDate != nil {
		interaction.InteractionDate = *req.InteractionDate
	}

	if err := h.interactionRepo.Create(c.Request.Context(), interaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create interaction"})
		return
	}

	c.JSON(http.StatusCreated, mapInteractionToResponse(interaction))
}

func (h *CRMHandlers) GetInteraction(c *gin.Context) {
	id := c.Param("id")
	interaction, err := h.interactionRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Interaction not found"})
		return
	}

	c.JSON(http.StatusOK, mapInteractionToResponse(interaction))
}

func (h *CRMHandlers) ListInteractionsByCustomer(c *gin.Context) {
	customerID := c.Param("customer_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	interactions, total, err := h.interactionRepo.ListByCustomer(c.Request.Context(), customerID, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch interactions"})
		return
	}

	response := make([]dto.InteractionResponse, len(interactions))
	for i, inter := range interactions {
		response[i] = mapInteractionToResponse(inter)
	}

	c.JSON(http.StatusOK, dto.ListResponse{
		Data:  response,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

// Mapping functions
func mapCustomerToResponse(cust *entities.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          cust.ID,
		CompanyName: cust.CompanyName,
		ContactName: cust.ContactName,
		Email:       cust.Email,
		Phone:       cust.Phone,
		Address:     cust.Address,
		Industry:    cust.Industry,
		Status:      string(cust.Status),
		AssignedTo:  cust.AssignedTo,
		CreatedBy:   cust.CreatedBy,
		CreatedAt:   cust.CreatedAt,
		UpdatedAt:   cust.UpdatedAt,
	}
}

func mapInteractionToResponse(inter *entities.CustomerInteraction) dto.InteractionResponse {
	return dto.InteractionResponse{
		ID:              inter.ID,
		CustomerID:      inter.CustomerID,
		UserID:          inter.UserID,
		InteractionType: string(inter.InteractionType),
		Subject:         inter.Subject,
		Notes:           inter.Notes,
		InteractionDate: inter.InteractionDate,
		CreatedAt:       inter.CreatedAt,
	}
}
