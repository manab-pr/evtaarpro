package payroll

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/payroll/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/payroll/presentation/http/handlers"
)

// RegisterRoutes registers payroll module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Initialize repositories
	employeeRepo := postgresql.NewEmployeeRepository(pgStore.DB)
	payrollRecordRepo := postgresql.NewPayrollRecordRepository(pgStore.DB)
	attendanceRepo := postgresql.NewAttendanceRepository(pgStore.DB)

	// Initialize handlers
	payrollHandlers := handlers.NewPayrollHandlers(employeeRepo, payrollRecordRepo, attendanceRepo)

	// Register routes with auth middleware
	payroll := rg.Group("/payroll")
	payroll.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
	{
		// Employee routes
		payroll.POST("/employees", payrollHandlers.CreateEmployee)
		payroll.GET("/employees", payrollHandlers.ListEmployees)
		payroll.GET("/employees/:id", payrollHandlers.GetEmployee)
		payroll.PUT("/employees/:id", payrollHandlers.UpdateEmployee)

		// Payroll record routes
		payroll.POST("/records", payrollHandlers.CreatePayrollRecord)
		payroll.GET("/records", payrollHandlers.ListPayrollRecords)
		payroll.GET("/records/:id", payrollHandlers.GetPayrollRecord)
		payroll.POST("/records/:id/approve", payrollHandlers.ApprovePayrollRecord)
		payroll.POST("/records/:id/pay", payrollHandlers.MarkPayrollRecordAsPaid)

		// Attendance routes
		payroll.POST("/attendance", payrollHandlers.CreateAttendance)
		payroll.GET("/attendance", payrollHandlers.ListAttendance)
		payroll.GET("/attendance/:id", payrollHandlers.GetAttendance)
	}
}
