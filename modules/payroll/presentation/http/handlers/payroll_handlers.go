package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/payroll/domain/repository"
	"github.com/manab-pr/evtaarpro/modules/payroll/presentation/http/dto"
)

type PayrollHandlers struct {
	employeeRepo       repository.EmployeeRepository
	payrollRecordRepo  repository.PayrollRecordRepository
	attendanceRepo     repository.AttendanceRepository
}

func NewPayrollHandlers(
	employeeRepo repository.EmployeeRepository,
	payrollRecordRepo repository.PayrollRecordRepository,
	attendanceRepo repository.AttendanceRepository,
) *PayrollHandlers {
	return &PayrollHandlers{
		employeeRepo:      employeeRepo,
		payrollRecordRepo: payrollRecordRepo,
		attendanceRepo:    attendanceRepo,
	}
}

// Employee Handlers
func (h *PayrollHandlers) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := entities.NewEmployee(req.EmployeeCode, req.Department, req.Designation, req.DateOfJoining, req.SalaryAmount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.ID = req.UserID
	if req.SalaryCurrency != "" {
		employee.SalaryCurrency = req.SalaryCurrency
	}
	employee.BankAccount = req.BankAccount
	employee.TaxID = req.TaxID

	if err := h.employeeRepo.Create(c.Request.Context(), employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusCreated, mapEmployeeToResponse(employee))
}

func (h *PayrollHandlers) GetEmployee(c *gin.Context) {
	id := c.Param("id")
	employee, err := h.employeeRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, mapEmployeeToResponse(employee))
}

func (h *PayrollHandlers) ListEmployees(c *gin.Context) {
	department := c.Query("department")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	employees, total, err := h.employeeRepo.List(c.Request.Context(), department, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	response := make([]dto.EmployeeResponse, len(employees))
	for i, emp := range employees {
		response[i] = mapEmployeeToResponse(emp)
	}

	c.JSON(http.StatusOK, dto.ListResponse{
		Data:  response,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

func (h *PayrollHandlers) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee, err := h.employeeRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	if req.EmployeeCode != "" {
		employee.EmployeeCode = req.EmployeeCode
	}
	if req.Department != "" {
		employee.Department = req.Department
	}
	if req.Designation != "" {
		employee.Designation = req.Designation
	}
	if !req.DateOfJoining.IsZero() {
		employee.DateOfJoining = req.DateOfJoining
	}
	if req.SalaryAmount > 0 {
		employee.SalaryAmount = req.SalaryAmount
	}
	if req.SalaryCurrency != "" {
		employee.SalaryCurrency = req.SalaryCurrency
	}
	employee.BankAccount = req.BankAccount
	employee.TaxID = req.TaxID
	employee.UpdatedAt = time.Now()

	if err := h.employeeRepo.Update(c.Request.Context(), employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, mapEmployeeToResponse(employee))
}

// Payroll Record Handlers
func (h *PayrollHandlers) CreatePayrollRecord(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var req dto.CreatePayrollRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := entities.NewPayrollRecord(
		req.EmployeeID,
		req.PeriodStart,
		req.PeriodEnd,
		req.GrossSalary,
		req.Deductions,
		userID.(string),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.payrollRecordRepo.Create(c.Request.Context(), record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payroll record"})
		return
	}

	c.JSON(http.StatusCreated, mapPayrollRecordToResponse(record))
}

func (h *PayrollHandlers) GetPayrollRecord(c *gin.Context) {
	id := c.Param("id")
	record, err := h.payrollRecordRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payroll record not found"})
		return
	}

	c.JSON(http.StatusOK, mapPayrollRecordToResponse(record))
}

func (h *PayrollHandlers) ListPayrollRecords(c *gin.Context) {
	employeeID := c.Query("employee_id")
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var payrollStatus entities.PayrollStatus
	if status != "" {
		payrollStatus = entities.PayrollStatus(status)
	}

	records, total, err := h.payrollRecordRepo.List(c.Request.Context(), employeeID, payrollStatus, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payroll records"})
		return
	}

	response := make([]dto.PayrollRecordResponse, len(records))
	for i, rec := range records {
		response[i] = mapPayrollRecordToResponse(rec)
	}

	c.JSON(http.StatusOK, dto.ListResponse{
		Data:  response,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

func (h *PayrollHandlers) ApprovePayrollRecord(c *gin.Context) {
	id := c.Param("id")
	record, err := h.payrollRecordRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payroll record not found"})
		return
	}

	if err := record.Approve(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.payrollRecordRepo.Update(c.Request.Context(), record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve payroll record"})
		return
	}

	c.JSON(http.StatusOK, mapPayrollRecordToResponse(record))
}

func (h *PayrollHandlers) MarkPayrollRecordAsPaid(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		PaymentReference string `json:"payment_reference" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := h.payrollRecordRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payroll record not found"})
		return
	}

	if err := record.MarkAsPaid(req.PaymentReference); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.payrollRecordRepo.Update(c.Request.Context(), record); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark payroll record as paid"})
		return
	}

	c.JSON(http.StatusOK, mapPayrollRecordToResponse(record))
}

// Attendance Handlers
func (h *PayrollHandlers) CreateAttendance(c *gin.Context) {
	var req dto.CreateAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance, err := entities.NewAttendance(req.EmployeeID, req.Date, entities.AttendanceStatus(req.Status))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendance.CheckIn = req.CheckIn
	attendance.CheckOut = req.CheckOut
	attendance.Notes = req.Notes

	if err := h.attendanceRepo.Create(c.Request.Context(), attendance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attendance record"})
		return
	}

	c.JSON(http.StatusCreated, mapAttendanceToResponse(attendance))
}

func (h *PayrollHandlers) GetAttendance(c *gin.Context) {
	id := c.Param("id")
	attendance, err := h.attendanceRepo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance record not found"})
		return
	}

	c.JSON(http.StatusOK, mapAttendanceToResponse(attendance))
}

func (h *PayrollHandlers) ListAttendance(c *gin.Context) {
	employeeID := c.Query("employee_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	attendances, total, err := h.attendanceRepo.List(c.Request.Context(), employeeID, startDate, endDate, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance records"})
		return
	}

	response := make([]dto.AttendanceResponse, len(attendances))
	for i, att := range attendances {
		response[i] = mapAttendanceToResponse(att)
	}

	c.JSON(http.StatusOK, dto.ListResponse{
		Data:  response,
		Total: total,
		Page:  page,
		Limit: limit,
	})
}

// Mapping functions
func mapEmployeeToResponse(emp *entities.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		ID:             emp.ID,
		EmployeeCode:   emp.EmployeeCode,
		Department:     emp.Department,
		Designation:    emp.Designation,
		DateOfJoining:  emp.DateOfJoining,
		SalaryAmount:   emp.SalaryAmount,
		SalaryCurrency: emp.SalaryCurrency,
		BankAccount:    emp.BankAccount,
		TaxID:          emp.TaxID,
		CreatedAt:      emp.CreatedAt,
		UpdatedAt:      emp.UpdatedAt,
	}
}

func mapPayrollRecordToResponse(record *entities.PayrollRecord) dto.PayrollRecordResponse {
	return dto.PayrollRecordResponse{
		ID:               record.ID,
		EmployeeID:       record.EmployeeID,
		PeriodStart:      record.PeriodStart,
		PeriodEnd:        record.PeriodEnd,
		GrossSalary:      record.GrossSalary,
		Deductions:       record.Deductions,
		NetSalary:        record.NetSalary,
		Status:           string(record.Status),
		PaidOn:           record.PaidOn,
		PaymentReference: record.PaymentReference,
		CreatedBy:        record.CreatedBy,
		CreatedAt:        record.CreatedAt,
		UpdatedAt:        record.UpdatedAt,
	}
}

func mapAttendanceToResponse(att *entities.Attendance) dto.AttendanceResponse {
	return dto.AttendanceResponse{
		ID:         att.ID,
		EmployeeID: att.EmployeeID,
		Date:       att.Date,
		CheckIn:    att.CheckIn,
		CheckOut:   att.CheckOut,
		Status:     string(att.Status),
		Notes:      att.Notes,
		CreatedAt:  att.CreatedAt,
	}
}
