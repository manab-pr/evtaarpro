package dto

import "time"

// Employee DTOs
type CreateEmployeeRequest struct {
	UserID         string    `json:"user_id" binding:"required"`
	EmployeeCode   string    `json:"employee_code" binding:"required"`
	Department     string    `json:"department"`
	Designation    string    `json:"designation"`
	DateOfJoining  time.Time `json:"date_of_joining" binding:"required"`
	SalaryAmount   float64   `json:"salary_amount" binding:"required"`
	SalaryCurrency string    `json:"salary_currency"`
	BankAccount    *string   `json:"bank_account"`
	TaxID          *string   `json:"tax_id"`
}

type UpdateEmployeeRequest struct {
	EmployeeCode   string    `json:"employee_code"`
	Department     string    `json:"department"`
	Designation    string    `json:"designation"`
	DateOfJoining  time.Time `json:"date_of_joining"`
	SalaryAmount   float64   `json:"salary_amount"`
	SalaryCurrency string    `json:"salary_currency"`
	BankAccount    *string   `json:"bank_account"`
	TaxID          *string   `json:"tax_id"`
}

type EmployeeResponse struct {
	ID             string     `json:"id"`
	EmployeeCode   string     `json:"employee_code"`
	Department     string     `json:"department"`
	Designation    string     `json:"designation"`
	DateOfJoining  time.Time  `json:"date_of_joining"`
	SalaryAmount   float64    `json:"salary_amount"`
	SalaryCurrency string     `json:"salary_currency"`
	BankAccount    *string    `json:"bank_account,omitempty"`
	TaxID          *string    `json:"tax_id,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// Payroll Record DTOs
type CreatePayrollRecordRequest struct {
	EmployeeID  string    `json:"employee_id" binding:"required"`
	PeriodStart time.Time `json:"period_start" binding:"required"`
	PeriodEnd   time.Time `json:"period_end" binding:"required"`
	GrossSalary float64   `json:"gross_salary" binding:"required"`
	Deductions  float64   `json:"deductions"`
}

type UpdatePayrollRecordRequest struct {
	GrossSalary float64 `json:"gross_salary"`
	Deductions  float64 `json:"deductions"`
	Status      string  `json:"status"`
}

type PayrollRecordResponse struct {
	ID               string     `json:"id"`
	EmployeeID       string     `json:"employee_id"`
	PeriodStart      time.Time  `json:"period_start"`
	PeriodEnd        time.Time  `json:"period_end"`
	GrossSalary      float64    `json:"gross_salary"`
	Deductions       float64    `json:"deductions"`
	NetSalary        float64    `json:"net_salary"`
	Status           string     `json:"status"`
	PaidOn           *time.Time `json:"paid_on,omitempty"`
	PaymentReference *string    `json:"payment_reference,omitempty"`
	CreatedBy        string     `json:"created_by"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// Attendance DTOs
type CreateAttendanceRequest struct {
	EmployeeID string    `json:"employee_id" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
	CheckIn    *time.Time `json:"check_in"`
	CheckOut   *time.Time `json:"check_out"`
	Status     string    `json:"status" binding:"required"`
	Notes      *string   `json:"notes"`
}

type UpdateAttendanceRequest struct {
	CheckIn  *time.Time `json:"check_in"`
	CheckOut *time.Time `json:"check_out"`
	Status   string     `json:"status"`
	Notes    *string    `json:"notes"`
}

type AttendanceResponse struct {
	ID         string     `json:"id"`
	EmployeeID string     `json:"employee_id"`
	Date       time.Time  `json:"date"`
	CheckIn    *time.Time `json:"check_in,omitempty"`
	CheckOut   *time.Time `json:"check_out,omitempty"`
	Status     string     `json:"status"`
	Notes      *string    `json:"notes,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}

type ListResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}
