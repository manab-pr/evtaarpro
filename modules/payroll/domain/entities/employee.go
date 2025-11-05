package entities

import (
	"errors"
	"time"
)

type Employee struct {
	ID             string
	EmployeeCode   string
	Department     string
	Designation    string
	DateOfJoining  time.Time
	SalaryAmount   float64
	SalaryCurrency string
	BankAccount    *string
	TaxID          *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewEmployee(employeeCode, department, designation string, dateOfJoining time.Time, salaryAmount float64) (*Employee, error) {
	if employeeCode == "" {
		return nil, errors.New("employee code is required")
	}
	if salaryAmount <= 0 {
		return nil, errors.New("salary amount must be positive")
	}

	return &Employee{
		EmployeeCode:   employeeCode,
		Department:     department,
		Designation:    designation,
		DateOfJoining:  dateOfJoining,
		SalaryAmount:   salaryAmount,
		SalaryCurrency: "USD",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}
