package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PayrollStatus string

const (
	StatusDraft    PayrollStatus = "draft"
	StatusApproved PayrollStatus = "approved"
	StatusPaid     PayrollStatus = "paid"
)

type PayrollRecord struct {
	ID                string
	EmployeeID        string
	PeriodStart       time.Time
	PeriodEnd         time.Time
	GrossSalary       float64
	Deductions        float64
	NetSalary         float64
	Status            PayrollStatus
	PaidOn            *time.Time
	PaymentReference  *string
	CreatedBy         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewPayrollRecord(employeeID string, periodStart, periodEnd time.Time, grossSalary, deductions float64, createdBy string) (*PayrollRecord, error) {
	if employeeID == "" {
		return nil, errors.New("employee ID is required")
	}
	if grossSalary <= 0 {
		return nil, errors.New("gross salary must be positive")
	}
	if periodStart.After(periodEnd) {
		return nil, errors.New("period start must be before period end")
	}

	netSalary := grossSalary - deductions
	if netSalary < 0 {
		netSalary = 0
	}

	return &PayrollRecord{
		ID:          uuid.New().String(),
		EmployeeID:  employeeID,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		GrossSalary: grossSalary,
		Deductions:  deductions,
		NetSalary:   netSalary,
		Status:      StatusDraft,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (pr *PayrollRecord) Approve() error {
	if pr.Status != StatusDraft {
		return errors.New("only draft records can be approved")
	}
	pr.Status = StatusApproved
	pr.UpdatedAt = time.Now()
	return nil
}

func (pr *PayrollRecord) MarkAsPaid(paymentReference string) error {
	if pr.Status != StatusApproved {
		return errors.New("only approved records can be marked as paid")
	}
	pr.Status = StatusPaid
	paidTime := time.Now()
	pr.PaidOn = &paidTime
	pr.PaymentReference = &paymentReference
	pr.UpdatedAt = time.Now()
	return nil
}
