package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/ports"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/crm/presentation/http/dto"
)

// CustomerHandlers contains CRM-related HTTP handlers
type CustomerHandlers struct {
	createCustomerUC    *usecases.CreateCustomerUseCase
	listCustomersUC     *usecases.ListCustomersUseCase
	addInteractionUC    *usecases.AddInteractionUseCase
	customerRepo        ports.CustomerRepository
}

// NewCustomerHandlers creates new CustomerHandlers
func NewCustomerHandlers(
	createCustomerUC *usecases.CreateCustomerUseCase,
	listCustomersUC *usecases.ListCustomersUseCase,
	addInteractionUC *usecases.AddInteractionUseCase,
	customerRepo ports.CustomerRepository,
) *CustomerHandlers {
	return &CustomerHandlers{
		createCustomerUC: createCustomerUC,
		listCustomersUC:  listCustomersUC,
		addInteractionUC: addInteractionUC,
		customerRepo:     customerRepo,
	}
}

// CreateCustomer creates a new customer
func (h *CustomerHandlers) CreateCustomer(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	customer, err := h.createCustomerUC.Execute(c.Request.Context(), usecases.CreateCustomerInput{
		CompanyID:  "default-company", // TODO: Get from user context
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Company:    req.Company,
		Status:     entities.CustomerStatus(req.Status),
		Source:     req.Source,
		AssignedTo: req.AssignedTo,
		Notes:      req.Notes,
		CreatedBy:  userID.(string),
	})

	if err != nil {
		response.InternalServerError(c, "Failed to create customer")
		return
	}

	response.Created(c, "Customer created successfully", mapCustomerToResponse(customer))
}

// ListCustomers lists customers with pagination
func (h *CustomerHandlers) ListCustomers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	customers, total, err := h.listCustomersUC.Execute(c.Request.Context(), "default-company", page, pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to list customers")
		return
	}

	customerResponses := make([]dto.CustomerResponse, len(customers))
	for i, customer := range customers {
		customerResponses[i] = mapCustomerToResponse(customer)
	}

	response.Paginated(c, customerResponses, page, pageSize, int64(total))
}

// GetCustomer retrieves a customer by ID
func (h *CustomerHandlers) GetCustomer(c *gin.Context) {
	customerID := c.Param("id")

	customer, err := h.customerRepo.GetByID(c.Request.Context(), customerID)
	if err != nil {
		response.NotFound(c, "Customer not found")
		return
	}

	response.OK(c, "Customer retrieved successfully", mapCustomerToResponse(customer))
}

// UpdateCustomer updates a customer
func (h *CustomerHandlers) UpdateCustomer(c *gin.Context) {
	customerID := c.Param("id")

	var req dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	customer, err := h.customerRepo.GetByID(c.Request.Context(), customerID)
	if err != nil {
		response.NotFound(c, "Customer not found")
		return
	}

	// Update fields
	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Email != "" {
		customer.Email = req.Email
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.Company != "" {
		customer.Company = req.Company
	}
	if req.Status != "" {
		customer.Status = entities.CustomerStatus(req.Status)
	}
	if req.Source != "" {
		customer.Source = req.Source
	}
	if req.AssignedTo != nil {
		customer.AssignedTo = req.AssignedTo
	}
	if req.Notes != "" {
		customer.Notes = req.Notes
	}

	if err := h.customerRepo.Update(c.Request.Context(), customer); err != nil {
		response.InternalServerError(c, "Failed to update customer")
		return
	}

	response.OK(c, "Customer updated successfully", mapCustomerToResponse(customer))
}

// DeleteCustomer deletes a customer
func (h *CustomerHandlers) DeleteCustomer(c *gin.Context) {
	customerID := c.Param("id")

	if err := h.customerRepo.Delete(c.Request.Context(), customerID); err != nil {
		response.InternalServerError(c, "Failed to delete customer")
		return
	}

	response.OK(c, "Customer deleted successfully", nil)
}

// AddInteraction adds an interaction to a customer
func (h *CustomerHandlers) AddInteraction(c *gin.Context) {
	customerID := c.Param("id")
	userID, _ := c.Get("user_id")

	var req dto.CreateInteractionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	interaction, err := h.addInteractionUC.Execute(c.Request.Context(), usecases.AddInteractionInput{
		CustomerID:  customerID,
		UserID:      userID.(string),
		Type:        req.Type,
		Subject:     req.Subject,
		Description: req.Description,
		ScheduledAt: req.ScheduledAt,
	})

	if err != nil {
		response.InternalServerError(c, "Failed to add interaction")
		return
	}

	response.Created(c, "Interaction added successfully", mapInteractionToResponse(interaction))
}

// GetInteractions retrieves interactions for a customer
func (h *CustomerHandlers) GetInteractions(c *gin.Context) {
	customerID := c.Param("id")

	interactions, err := h.customerRepo.GetInteractions(c.Request.Context(), customerID)
	if err != nil {
		response.InternalServerError(c, "Failed to get interactions")
		return
	}

	interactionResponses := make([]dto.InteractionResponse, len(interactions))
	for i, interaction := range interactions {
		interactionResponses[i] = mapInteractionToResponse(interaction)
	}

	response.OK(c, "Interactions retrieved successfully", interactionResponses)
}

func mapCustomerToResponse(customer *entities.Customer) dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:         customer.ID,
		CompanyID:  customer.CompanyID,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Company:    customer.Company,
		Status:     string(customer.Status),
		Source:     customer.Source,
		AssignedTo: customer.AssignedTo,
		Notes:      customer.Notes,
		CreatedBy:  customer.CreatedBy,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  customer.UpdatedAt,
	}
}

func mapInteractionToResponse(interaction *entities.CustomerInteraction) dto.InteractionResponse {
	return dto.InteractionResponse{
		ID:          interaction.ID,
		CustomerID:  interaction.CustomerID,
		UserID:      interaction.UserID,
		Type:        interaction.Type,
		Subject:     interaction.Subject,
		Description: interaction.Description,
		ScheduledAt: interaction.ScheduledAt,
		CompletedAt: interaction.CompletedAt,
		CreatedAt:   interaction.CreatedAt,
	}
}
