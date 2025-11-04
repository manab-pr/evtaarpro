# EvtaarPro - Project Setup Guide

## What Has Been Created

This project has been scaffolded with a **Clean Architecture** pattern following the structure of enterprise-grade Go applications. Here's what's been implemented:

### ✅ Completed Components

#### 1. **Project Structure**
- Clean architecture with clear separation of concerns
- Modular design with independent business modules
- Proper dependency injection setup

#### 2. **Core Infrastructure**
- **Configuration Management** (`internal/config/`)
  - YAML-based configuration with environment variable expansion
  - Separate configs for app, PostgreSQL, and Redis

- **Database Layer** (`internal/datastore/`)
  - PostgreSQL connection with connection pooling
  - Redis connection with session management
  - Health check implementations

- **HTTP Layer** (`internal/httpx/`)
  - Gin router with middleware chain
  - Health and readiness endpoints
  - Module-based route registration

#### 3. **Middleware** (`internal/middleware/`)
- ✅ CORS handling
- ✅ Request logging with request IDs
- ✅ JWT authentication
- ✅ Role-based access control
- ✅ Panic recovery

#### 4. **Authentication Module** (`modules/auth/`)
- **Domain Layer:**
  - User entity with role management
  - Use cases: Register, Login, Logout
  - Port interfaces for repositories and services

- **Infrastructure Layer:**
  - Bcrypt password hasher
  - JWT token generator
  - PostgreSQL user repository
  - Redis session store

- **Presentation Layer:**
  - HTTP handlers for register, login, logout
  - DTOs for request/response
  - Route registration

**API Endpoints:**
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login and get JWT tokens
- `POST /api/v1/auth/logout` - Logout (requires authentication)

#### 5. **Database Migrations** (`migrations/`)
- ✅ Users table with roles and authentication
- ✅ Meetings and participants tables
- ✅ CRM tables (customers, interactions)
- ✅ Payroll tables (employees, payroll records, attendance)
- ✅ Notifications table
- ✅ Proper indexes and foreign keys
- ✅ Triggers for updated_at timestamps

#### 6. **CI/CD Pipeline** (`.github/workflows/`)
- ✅ Automated testing with PostgreSQL and Redis services
- ✅ Code linting with golangci-lint
- ✅ Docker image building and pushing to GitHub Container Registry
- ✅ AWS EKS deployment automation

#### 7. **Docker Setup** (`deploy/`)
- ✅ Multi-stage Dockerfile for production
- ✅ Docker Compose for local development
- ✅ Prometheus and Grafana for monitoring

#### 8. **Placeholder Modules**
- Users module (stub)
- Meetings module (stub) - For Jitsi integration
- CRM module (stub)
- Payroll module (stub)
- Notifications module (stub) - For WebSocket

### 📋 Module Placeholders Ready for Implementation

The following modules have container files ready but need full implementation:

1. **Users Module** - User profile management
2. **Meetings Module** - Jitsi/OpenVidu video conferencing
3. **CRM Module** - Customer relationship management
4. **Payroll Module** - Salary, attendance management
5. **Notifications Module** - Real-time WebSocket notifications

## Quick Start

### Prerequisites

```bash
# Required
- Go 1.22+
- Docker & Docker Compose
- PostgreSQL 15+
- Redis 7+

# Optional (for deployment)
- AWS CLI
- kubectl
- Helm
```

### Step 1: Clone and Setup

```bash
cd evtaarpro
cp .env.example .env

# Update .env with your settings
# Especially: JWT_SECRET, GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET
```

### Step 2: Start Dependencies

```bash
# Start PostgreSQL and Redis
docker-compose -f deploy/docker-compose.local.yml up -d postgres redis

# Wait for services to be healthy (check with docker ps)
```

### Step 3: Run Migrations

```bash
# Install migration tool (if not already installed)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run migrations
migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/evtaarpro?sslmode=disable" up

# Or manually apply each SQL file
psql -h localhost -U postgres -d evtaarpro -f migrations/001_create_users_table.sql
# ... repeat for each migration file
```

### Step 4: Install Dependencies

```bash
go mod download
go mod tidy
```

### Step 5: Run the Application

```bash
# Using Make
make run

# Or directly
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

### Step 6: Test the API

```bash
# Health check
curl http://localhost:8080/health

# Register a new user
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@evtaarpro.com",
    "password": "SecurePass123!",
    "first_name": "Admin",
    "last_name": "User",
    "role": "admin"
  }'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@evtaarpro.com",
    "password": "SecurePass123!"
  }'

# Use the access_token from the response for authenticated requests
```

## Development Workflow

### Running Tests

```bash
# All tests
make test

# With coverage
make test-coverage

# Integration tests
make test-integration
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Run vet
make vet

# All checks
make check
```

### Generate API Documentation

```bash
# Install swag if not present
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger docs
make swagger

# View at http://localhost:8080/swagger/index.html
```

### Docker Development

```bash
# Build Docker image
make docker-build

# Run with docker-compose
make docker-compose-up

# View logs
make docker-compose-logs

# Stop services
make docker-compose-down
```

## Architecture Overview

```
┌─────────────────────────────────────────────────────┐
│                   HTTP Layer (Gin)                   │
│                   /api/v1/*                          │
└──────────────────────┬──────────────────────────────┘
                       │
┌──────────────────────▼──────────────────────────────┐
│                   Modules                            │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐          │
│  │  Auth    │  │ Meetings │  │   CRM    │  ...     │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘          │
│       │             │              │                 │
│  ┌────▼─────────────▼──────────────▼─────┐          │
│  │         Domain Layer                   │          │
│  │  (Entities, Use Cases, Ports)          │          │
│  └────────────────┬───────────────────────┘          │
│                   │                                  │
│  ┌────────────────▼───────────────────────┐          │
│  │      Infrastructure Layer              │          │
│  │  (PostgreSQL, Redis, Jitsi, AWS S3)    │          │
│  └────────────────────────────────────────┘          │
└─────────────────────────────────────────────────────┘
```

## Next Steps to Complete the Project

### 1. Implement Users Module
- User profile CRUD operations
- Profile picture upload to S3
- User preferences
- Password reset flow

### 2. Implement Meetings Module
- Jitsi API integration
- Meeting scheduling
- Participant management
- Recording to S3
- WebSocket for real-time updates

### 3. Implement CRM Module
- Customer management
- Lead tracking
- Interaction logs
- Pipeline management
- Reports and analytics

### 4. Implement Payroll Module
- Employee records
- Salary calculations
- Attendance tracking
- Payslip generation
- Tax calculations

### 5. Implement Notifications Module
- WebSocket server
- Push notifications
- Email notifications
- Real-time event streaming

### 6. Add Google OAuth
- Implement OAuth2 flow in auth module
- Handle Google callback
- Link social accounts

### 7. Deploy to AWS EKS
- Create EKS cluster
- Set up RDS for PostgreSQL
- Set up ElastiCache for Redis
- Configure S3 buckets
- Set up Load Balancer
- Configure GitHub secrets for CI/CD

### 8. Monitoring & Observability
- Set up Prometheus metrics
- Configure Grafana dashboards
- Add structured logging
- Set up alerts

## Useful Commands

```bash
# Development
make run                    # Run the application
make build                  # Build binary
make clean                  # Clean artifacts

# Testing
make test                   # Run tests
make test-coverage          # Run tests with coverage
make test-integration       # Run integration tests

# Code Quality
make lint                   # Run linter
make fmt                    # Format code
make check                  # Run all checks

# Docker
make docker-build           # Build Docker image
make docker-push            # Push to registry
make docker-compose-up      # Start with docker-compose
make docker-compose-down    # Stop services

# Kubernetes
make k8s-deploy             # Deploy to Kubernetes
make k8s-logs               # View logs
make k8s-restart            # Restart deployment

# Database
make migrate-up             # Run migrations
make migrate-down           # Rollback migrations
make seed                   # Seed database

# Documentation
make swagger                # Generate Swagger docs
```

## Environment Variables Reference

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `APP_ENV` | Environment (development/production) | No | development |
| `APP_PORT` | Server port | No | 8080 |
| `DB_HOST` | PostgreSQL host | Yes | localhost |
| `DB_PORT` | PostgreSQL port | No | 5432 |
| `DB_USER` | PostgreSQL user | Yes | postgres |
| `DB_PASSWORD` | PostgreSQL password | Yes | - |
| `DB_NAME` | Database name | Yes | evtaarpro |
| `REDIS_HOST` | Redis host | Yes | localhost |
| `REDIS_PORT` | Redis port | No | 6379 |
| `JWT_SECRET` | JWT signing secret | Yes | - |
| `GOOGLE_CLIENT_ID` | Google OAuth client ID | Yes | - |
| `GOOGLE_CLIENT_SECRET` | Google OAuth secret | Yes | - |
| `JITSI_API_URL` | Jitsi server URL | Yes | - |
| `AWS_S3_BUCKET` | S3 bucket for files | Yes | - |

## Project Statistics

- **Total Files Created**: ~90 files
- **Lines of Code**: ~5000+ lines
- **Modules**: 6 modules (1 fully implemented, 5 scaffolded)
- **Database Tables**: 10 tables with proper relationships
- **API Endpoints**: 3 auth endpoints + placeholders for others
- **Middleware**: 5 middleware functions
- **CI/CD**: Fully automated pipeline

## Support & Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Ensure PostgreSQL is running: `docker ps`
   - Check connection string in config files
   - Verify credentials in .env file

2. **Redis Connection Failed**
   - Ensure Redis is running: `docker ps`
   - Check Redis host and port

3. **JWT Validation Failed**
   - Ensure JWT_SECRET is set in .env
   - Check token expiration settings

4. **Migration Errors**
   - Run migrations in order (001, 002, etc.)
   - Check PostgreSQL user permissions
   - Verify database exists

## Contributing

1. Create a feature branch
2. Make your changes
3. Write tests
4. Run `make check`
5. Submit a pull request

## License

MIT License - See LICENSE file for details
