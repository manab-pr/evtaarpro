package dto

import "time"

// CreateEmployeeRequest represents an employee creation request
type CreateEmployeeRequest struct {
	UserID       *string   `json:"user_id"`
	EmployeeCode string    `json:"employee_code" binding:"required"`
	Department   string    `json:"department" binding:"required"`
	Designation  string    `json:"designation" binding:"required"`
	JoiningDate  time.Time `json:"joining_date" binding:"required"`
	SalaryAmount float64   `json:"salary_amount" binding:"required"`
}

// EmployeeResponse represents an employee response
type EmployeeResponse struct {
	ID           string    `json:"id"`
	UserID       *string   `json:"user_id,omitempty"`
	EmployeeCode string    `json:"employee_code"`
	Department   string    `json:"department"`
	Designation  string    `json:"designation"`
	JoiningDate  time.Time `json:"joining_date"`
	SalaryAmount float64   `json:"salary_amount"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// MarkAttendanceRequest represents an attendance request
type MarkAttendanceRequest struct {
	EmployeeID string     `json:"employee_id" binding:"required"`
	Date       time.Time  `json:"date" binding:"required"`
	CheckIn    *time.Time `json:"check_in"`
	CheckOut   *time.Time `json:"check_out"`
	Status     string     `json:"status" binding:"required"` // present, absent, halfday, leave
	Notes      string     `json:"notes"`
}

// AttendanceResponse represents an attendance response
type AttendanceResponse struct {
	ID          string     `json:"id"`
	EmployeeID  string     `json:"employee_id"`
	Date        time.Time  `json:"date"`
	CheckIn     *time.Time `json:"check_in,omitempty"`
	CheckOut    *time.Time `json:"check_out,omitempty"`
	Status      string     `json:"status"`
	HoursWorked float64    `json:"hours_worked"`
	Notes       string     `json:"notes"`
	CreatedAt   time.Time  `json:"created_at"`
}

// GeneratePayrollRequest represents a payroll generation request
type GeneratePayrollRequest struct {
	EmployeeID string  `json:"employee_id" binding:"required"`
	Month      int     `json:"month" binding:"required"`
	Year       int     `json:"year" binding:"required"`
	Allowances float64 `json:"allowances"`
	Deductions float64 `json:"deductions"`
	Notes      string  `json:"notes"`
}

// PayrollRecordResponse represents a payroll record response
type PayrollRecordResponse struct {
	ID            string     `json:"id"`
	EmployeeID    string     `json:"employee_id"`
	Month         int        `json:"month"`
	Year          int        `json:"year"`
	BasicSalary   float64    `json:"basic_salary"`
	Allowances    float64    `json:"allowances"`
	Deductions    float64    `json:"deductions"`
	NetSalary     float64    `json:"net_salary"`
	PaymentStatus string     `json:"payment_status"`
	PaymentDate   *time.Time `json:"payment_date,omitempty"`
	Notes         string     `json:"notes"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
