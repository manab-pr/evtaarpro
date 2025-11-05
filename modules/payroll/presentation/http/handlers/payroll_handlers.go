package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/payroll/domain/ports"
	"github.com/manab-pr/evtaarpro/modules/payroll/presentation/http/dto"
)

// PayrollHandlers contains payroll-related HTTP handlers
type PayrollHandlers struct {
	payrollRepo ports.PayrollRepository
}

// NewPayrollHandlers creates new PayrollHandlers
func NewPayrollHandlers(payrollRepo ports.PayrollRepository) *PayrollHandlers {
	return &PayrollHandlers{
		payrollRepo: payrollRepo,
	}
}

// CreateEmployee creates a new employee
func (h *PayrollHandlers) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	employee := &entities.Employee{
		ID:           uuid.New().String(),
		UserID:       req.UserID,
		EmployeeCode: req.EmployeeCode,
		Department:   req.Department,
		Designation:  req.Designation,
		JoiningDate:  req.JoiningDate,
		SalaryAmount: req.SalaryAmount,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := h.payrollRepo.CreateEmployee(c.Request.Context(), employee); err != nil {
		response.InternalServerError(c, "Failed to create employee")
		return
	}

	response.Created(c, "Employee created successfully", mapEmployeeToResponse(employee))
}

// ListEmployees lists employees with pagination
func (h *PayrollHandlers) ListEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	employees, total, err := h.payrollRepo.ListEmployees(c.Request.Context(), pageSize, (page-1)*pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to list employees")
		return
	}

	employeeResponses := make([]dto.EmployeeResponse, len(employees))
	for i, employee := range employees {
		employeeResponses[i] = mapEmployeeToResponse(employee)
	}

	response.Paginated(c, employeeResponses, page, pageSize, int64(total))
}

// GetEmployee retrieves an employee by ID
func (h *PayrollHandlers) GetEmployee(c *gin.Context) {
	employeeID := c.Param("id")

	employee, err := h.payrollRepo.GetEmployee(c.Request.Context(), employeeID)
	if err != nil {
		response.NotFound(c, "Employee not found")
		return
	}

	response.OK(c, "Employee retrieved successfully", mapEmployeeToResponse(employee))
}

// MarkAttendance marks attendance for an employee
func (h *PayrollHandlers) MarkAttendance(c *gin.Context) {
	var req dto.MarkAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Calculate hours worked if both check-in and check-out are provided
	var hoursWorked float64
	if req.CheckIn != nil && req.CheckOut != nil {
		duration := req.CheckOut.Sub(*req.CheckIn)
		hoursWorked = duration.Hours()
	}

	attendance := &entities.Attendance{
		ID:          uuid.New().String(),
		EmployeeID:  req.EmployeeID,
		Date:        req.Date,
		CheckIn:     req.CheckIn,
		CheckOut:    req.CheckOut,
		Status:      req.Status,
		HoursWorked: hoursWorked,
		Notes:       req.Notes,
		CreatedAt:   time.Now(),
	}

	if err := h.payrollRepo.CreateAttendance(c.Request.Context(), attendance); err != nil {
		response.InternalServerError(c, "Failed to mark attendance")
		return
	}

	response.Created(c, "Attendance marked successfully", mapAttendanceToResponse(attendance))
}

// GetAttendance retrieves attendance records
func (h *PayrollHandlers) GetAttendance(c *gin.Context) {
	employeeID := c.Query("employee_id")
	month, _ := strconv.Atoi(c.Query("month"))
	year, _ := strconv.Atoi(c.Query("year"))

	if employeeID == "" || month == 0 || year == 0 {
		response.BadRequest(c, "employee_id, month, and year are required")
		return
	}

	records, err := h.payrollRepo.GetAttendance(c.Request.Context(), employeeID, month, year)
	if err != nil {
		response.InternalServerError(c, "Failed to get attendance")
		return
	}

	attendanceResponses := make([]dto.AttendanceResponse, len(records))
	for i, record := range records {
		attendanceResponses[i] = mapAttendanceToResponse(record)
	}

	response.OK(c, "Attendance retrieved successfully", attendanceResponses)
}

// GeneratePayroll generates payroll for an employee
func (h *PayrollHandlers) GeneratePayroll(c *gin.Context) {
	var req dto.GeneratePayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Get employee to retrieve basic salary
	employee, err := h.payrollRepo.GetEmployee(c.Request.Context(), req.EmployeeID)
	if err != nil {
		response.NotFound(c, "Employee not found")
		return
	}

	// Calculate net salary
	netSalary := employee.SalaryAmount + req.Allowances - req.Deductions

	payrollRecord := &entities.PayrollRecord{
		ID:            uuid.New().String(),
		EmployeeID:    req.EmployeeID,
		Month:         req.Month,
		Year:          req.Year,
		BasicSalary:   employee.SalaryAmount,
		Allowances:    req.Allowances,
		Deductions:    req.Deductions,
		NetSalary:     netSalary,
		PaymentStatus: "pending",
		Notes:         req.Notes,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := h.payrollRepo.CreatePayrollRecord(c.Request.Context(), payrollRecord); err != nil {
		response.InternalServerError(c, "Failed to generate payroll")
		return
	}

	response.Created(c, "Payroll generated successfully", mapPayrollRecordToResponse(payrollRecord))
}

// ListPayrollRecords lists payroll records with pagination
func (h *PayrollHandlers) ListPayrollRecords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	records, total, err := h.payrollRepo.ListPayrollRecords(c.Request.Context(), pageSize, (page-1)*pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to list payroll records")
		return
	}

	payrollResponses := make([]dto.PayrollRecordResponse, len(records))
	for i, record := range records {
		payrollResponses[i] = mapPayrollRecordToResponse(record)
	}

	response.Paginated(c, payrollResponses, page, pageSize, int64(total))
}

// GetPayrollRecord retrieves a payroll record by ID
func (h *PayrollHandlers) GetPayrollRecord(c *gin.Context) {
	recordID := c.Param("id")

	record, err := h.payrollRepo.GetPayrollRecord(c.Request.Context(), recordID)
	if err != nil {
		response.NotFound(c, "Payroll record not found")
		return
	}

	response.OK(c, "Payroll record retrieved successfully", mapPayrollRecordToResponse(record))
}

func mapEmployeeToResponse(employee *entities.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		ID:           employee.ID,
		UserID:       employee.UserID,
		EmployeeCode: employee.EmployeeCode,
		Department:   employee.Department,
		Designation:  employee.Designation,
		JoiningDate:  employee.JoiningDate,
		SalaryAmount: employee.SalaryAmount,
		IsActive:     employee.IsActive,
		CreatedAt:    employee.CreatedAt,
		UpdatedAt:    employee.UpdatedAt,
	}
}

func mapAttendanceToResponse(attendance *entities.Attendance) dto.AttendanceResponse {
	return dto.AttendanceResponse{
		ID:          attendance.ID,
		EmployeeID:  attendance.EmployeeID,
		Date:        attendance.Date,
		CheckIn:     attendance.CheckIn,
		CheckOut:    attendance.CheckOut,
		Status:      attendance.Status,
		HoursWorked: attendance.HoursWorked,
		Notes:       attendance.Notes,
		CreatedAt:   attendance.CreatedAt,
	}
}

func mapPayrollRecordToResponse(record *entities.PayrollRecord) dto.PayrollRecordResponse {
	return dto.PayrollRecordResponse{
		ID:            record.ID,
		EmployeeID:    record.EmployeeID,
		Month:         record.Month,
		Year:          record.Year,
		BasicSalary:   record.BasicSalary,
		Allowances:    record.Allowances,
		Deductions:    record.Deductions,
		NetSalary:     record.NetSalary,
		PaymentStatus: record.PaymentStatus,
		PaymentDate:   record.PaymentDate,
		Notes:         record.Notes,
		CreatedAt:     record.CreatedAt,
		UpdatedAt:     record.UpdatedAt,
	}
}
