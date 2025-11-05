package entities

import "time"

// Employee represents an employee entity
type Employee struct {
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

// Attendance represents attendance record
type Attendance struct {
	ID          string     `json:"id"`
	EmployeeID  string     `json:"employee_id"`
	Date        time.Time  `json:"date"`
	CheckIn     *time.Time `json:"check_in,omitempty"`
	CheckOut    *time.Time `json:"check_out,omitempty"`
	Status      string     `json:"status"` // present, absent, halfday, leave
	HoursWorked float64    `json:"hours_worked"`
	Notes       string     `json:"notes"`
	CreatedAt   time.Time  `json:"created_at"`
}

// PayrollRecord represents a payroll record
type PayrollRecord struct {
	ID            string     `json:"id"`
	EmployeeID    string     `json:"employee_id"`
	Month         int        `json:"month"`
	Year          int        `json:"year"`
	BasicSalary   float64    `json:"basic_salary"`
	Allowances    float64    `json:"allowances"`
	Deductions    float64    `json:"deductions"`
	NetSalary     float64    `json:"net_salary"`
	PaymentStatus string     `json:"payment_status"` // pending, paid, failed
	PaymentDate   *time.Time `json:"payment_date,omitempty"`
	Notes         string     `json:"notes"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
