package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLeave   AttendanceStatus = "leave"
	StatusHoliday AttendanceStatus = "holiday"
)

type Attendance struct {
	ID         string
	EmployeeID string
	Date       time.Time
	CheckIn    *time.Time
	CheckOut   *time.Time
	Status     AttendanceStatus
	Notes      *string
	CreatedAt  time.Time
}

func NewAttendance(employeeID string, date time.Time, status AttendanceStatus) (*Attendance, error) {
	if employeeID == "" {
		return nil, errors.New("employee ID is required")
	}

	return &Attendance{
		ID:         uuid.New().String(),
		EmployeeID: employeeID,
		Date:       date,
		Status:     status,
		CreatedAt:  time.Now(),
	}, nil
}

func (a *Attendance) RecordCheckIn(checkInTime time.Time) {
	a.CheckIn = &checkInTime
	a.Status = StatusPresent
}

func (a *Attendance) RecordCheckOut(checkOutTime time.Time) error {
	if a.CheckIn == nil {
		return errors.New("cannot check out without checking in first")
	}
	if checkOutTime.Before(*a.CheckIn) {
		return errors.New("check out time cannot be before check in time")
	}
	a.CheckOut = &checkOutTime
	return nil
}
