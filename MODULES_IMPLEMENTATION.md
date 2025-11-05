# EvtaarPro - Complete Modules Implementation Guide

## Overview
This document provides the complete implementation guide for CRM, Payroll, and Notifications modules.

## Architecture
All modules follow Clean Architecture with:
- **Domain Layer**: Entities, Use Cases, Ports (interfaces)
- **Infrastructure Layer**: Database repositories, External services
- **Presentation Layer**: HTTP handlers, DTOs, Routes

---

## 1. CRM MODULE

### Features
- Customer management (CRUD)
- Lead tracking and conversion
- Customer interactions (calls, emails, meetings, notes)
- Customer assignment to sales reps
- Pipeline management

### Database Tables (Already Created)
```sql
customers (
  id, company_id, name, email, phone, company, status,
  source, assigned_to, notes, created_by, created_at, updated_at
)

customer_interactions (
  id, customer_id, user_id, type, subject, description,
  scheduled_at, completed_at, created_at
)
```

### API Endpoints
```
POST   /api/v1/crm/customers          - Create customer
GET    /api/v1/crm/customers          - List customers
GET    /api/v1/crm/customers/:id      - Get customer
PUT    /api/v1/crm/customers/:id      - Update customer
DELETE /api/v1/crm/customers/:id      - Delete customer

POST   /api/v1/crm/customers/:id/interactions  - Add interaction
GET    /api/v1/crm/customers/:id/interactions  - Get interactions
```

### Frontend Pages
- `/crm` - Customer list with filters and search
- `/crm/customers/:id` - Customer details with interaction timeline
- `/crm/customers/new` - Create new customer form

---

## 2. PAYROLL MODULE

### Features
- Employee salary management
- Attendance tracking
- Automatic payroll calculations
- Payslip generation
- Payment history

### Database Tables (Already Created)
```sql
employees (
  id, user_id, employee_code, department, designation,
  joining_date, salary_amount, is_active, created_at, updated_at
)

attendance (
  id, employee_id, date, check_in, check_out, status,
  hours_worked, notes, created_at
)

payroll_records (
  id, employee_id, month, year, basic_salary, allowances,
  deductions, net_salary, payment_status, payment_date,
  notes, created_at, updated_at
)
```

### API Endpoints
```
POST   /api/v1/payroll/employees           - Add employee
GET    /api/v1/payroll/employees           - List employees
GET    /api/v1/payroll/employees/:id       - Get employee

POST   /api/v1/payroll/attendance          - Mark attendance
GET    /api/v1/payroll/attendance          - Get attendance records
PUT    /api/v1/payroll/attendance/:id      - Update attendance

POST   /api/v1/payroll/records             - Generate payroll
GET    /api/v1/payroll/records             - List payroll records
GET    /api/v1/payroll/records/:id         - Get payroll details
GET    /api/v1/payroll/records/:id/payslip - Download payslip (PDF)
```

### Frontend Pages
- `/payroll` - Payroll dashboard with monthly summary
- `/payroll/employees` - Employee list
- `/payroll/attendance` - Attendance tracking interface
- `/payroll/records` - Payroll records history
- `/payroll/generate` - Generate monthly payroll

---

## 3. NOTIFICATIONS MODULE

### Features
- Real-time WebSocket notifications
- Meeting reminders
- Payroll alerts
- CRM activity notifications
- Push notifications to browser

### Database Tables (Already Created)
```sql
notifications (
  id, user_id, type, title, message, data,
  read, created_at, read_at
)
```

### WebSocket Protocol
```
Connection: ws://localhost:8080/api/v1/notifications/ws

Message Format:
{
  "type": "meeting_reminder|payroll_generated|customer_assigned",
  "title": "Notification Title",
  "message": "Notification message",
  "data": { ... },
  "timestamp": "2025-11-04T10:00:00Z"
}
```

### API Endpoints
```
GET    /api/v1/notifications/ws        - WebSocket connection
GET    /api/v1/notifications           - Get user notifications
PUT    /api/v1/notifications/:id/read  - Mark as read
DELETE /api/v1/notifications/:id       - Delete notification
POST   /api/v1/notifications/read-all  - Mark all as read
```

### Frontend Integration
- WebSocket connection in AuthContext
- Notification bell icon in header
- Toast notifications for real-time alerts
- Notification center dropdown

---

## Implementation Status

### ✅ Completed
- Database migrations for all tables
- Authentication module (JWT + OAuth)
- Meeting module (Jitsi integration)
- User management module

### 🚧 In Progress
- CRM module entities and repositories
- API route structure

### 📋 To Do
- Complete CRM backend implementation
- Complete Payroll backend implementation
- Complete Notifications WebSocket implementation
- Build all frontend pages
- End-to-end testing

---

## Quick Start Guide

### Step 1: Complete Backend Implementation

Run the implementation script:
```bash
./scripts/implement-modules.sh
```

This will generate all necessary files for CRM, Payroll, and Notifications.

### Step 2: Restart Backend
```bash
# Kill existing backend
pkill -f "go run cmd/server/main.go"

# Start fresh
go mod tidy
go run cmd/server/main.go
```

### Step 3: Test APIs
```bash
# Test CRM
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/crm/customers

# Test Payroll
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/payroll/employees

# Test Notifications
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/notifications
```

### Step 4: Build Frontend Pages
The frontend pages should be created in:
- `frontend/src/pages/CRM.jsx`
- `frontend/src/pages/Payroll.jsx`
- `frontend/src/pages/Notifications.jsx`

---

## Integration Points

### CRM → Notifications
When a customer is assigned to a sales rep, send a notification.

### Payroll → Notifications
When payroll is generated, notify HR and employees.

### Meetings → CRM
Link meetings to customers for interaction tracking.

### Attendance → Payroll
Attendance data feeds into payroll calculations.

---

## Testing Checklist

- [ ] Create a customer via CRM API
- [ ] Add interaction to customer
- [ ] Add employee to payroll system
- [ ] Mark attendance for employee
- [ ] Generate monthly payroll
- [ ] Receive real-time notification via WebSocket
- [ ] View notification history
- [ ] Mark notification as read

---

## Deployment Considerations

### Environment Variables
```env
# Already configured
DB_HOST=localhost
DB_PORT=5432
REDIS_HOST=localhost
REDIS_PORT=6379

# WebSocket
WS_ENABLED=true
WS_PORT=8080
```

### Docker Compose
All services (postgres, redis, backend, frontend) should run together.

### AWS EKS Deployment
- Each module can be deployed as a separate microservice
- Use Kubernetes Services for inter-service communication
- Use AWS ALB for load balancing
- Use AWS RDS for PostgreSQL
- Use AWS ElastiCache for Redis

---

## Next Steps

1. **Complete Implementation**: Run the generation scripts
2. **API Testing**: Use Postman/curl to test all endpoints
3. **Frontend Development**: Build React pages for each module
4. **Integration Testing**: Test cross-module features
5. **Documentation**: Update API documentation with Swagger
6. **Deployment**: Containerize and deploy to EKS

---

**Generated**: 2025-11-04
**Version**: 1.0
**Author**: EvtaarPro Team
