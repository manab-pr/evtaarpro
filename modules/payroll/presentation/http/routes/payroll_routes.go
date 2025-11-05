package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/payroll/presentation/http/handlers"
)

// RegisterRoutes registers payroll routes
func RegisterRoutes(rg *gin.RouterGroup, payrollHandlers *handlers.PayrollHandlers, jwtSecret string) {
	payroll := rg.Group("/payroll")
	payroll.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// Employee routes
		payroll.POST("/employees", payrollHandlers.CreateEmployee)
		payroll.GET("/employees", payrollHandlers.ListEmployees)
		payroll.GET("/employees/:id", payrollHandlers.GetEmployee)

		// Attendance routes
		payroll.POST("/attendance", payrollHandlers.MarkAttendance)
		payroll.GET("/attendance", payrollHandlers.GetAttendance)

		// Payroll routes
		payroll.POST("/records", payrollHandlers.GeneratePayroll)
		payroll.GET("/records", payrollHandlers.ListPayrollRecords)
		payroll.GET("/records/:id", payrollHandlers.GetPayrollRecord)
	}
}
