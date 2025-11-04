# EvtaarPro - Quick Start Guide

Get EvtaarPro running in 5 minutes!

## Automatic Setup (Recommended)

```bash
cd evtaarpro

# Run the setup script
bash scripts/setup.sh

# Start the application
make run
```

The setup script will:
- ✅ Check prerequisites (Go, Docker, Docker Compose)
- ✅ Install Go dependencies
- ✅ Create .env file
- ✅ Start PostgreSQL and Redis
- ✅ Initialize database with migrations
- ✅ Verify everything is working

## Manual Setup

### 1. Prerequisites

```bash
# Check you have:
go version       # Should be 1.22+
docker --version
docker-compose --version
```

### 2. Install Dependencies

```bash
cd evtaarpro
go mod download
```

### 3. Configure Environment

```bash
cp .env.example .env

# Edit .env and set at minimum:
# JWT_SECRET=your-secret-key-here
```

### 4. Start Services

```bash
# Start PostgreSQL and Redis
docker-compose -f deploy/docker-compose.local.yml up -d postgres redis

# Wait for services (about 10 seconds)
sleep 10
```

### 5. Initialize Database

```bash
bash scripts/init_db.sh
```

### 6. Run the Application

```bash
make run
# or
go run cmd/server/main.go
```

## Testing the API

### 1. Health Check

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy",
  "service": "evtaarpro"
}
```

### 2. Register a User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "SecurePass123",
    "first_name": "John",
    "last_name": "Doe",
    "role": "employee"
  }'
```

### 3. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "SecurePass123"
  }'
```

Save the `access_token` from the response.

### 4. Test Protected Endpoint

```bash
# Replace YOUR_ACCESS_TOKEN with the token from login
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## View API Documentation

Open your browser to:
```
http://localhost:8080/swagger/index.html
```

## What's Available

✅ **Auth Module**: Complete with register, login, logout
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/logout`

🚧 **Other Modules**: Placeholder routes (ready for implementation)
- Users: `/api/v1/users`
- Meetings: `/api/v1/meetings`
- CRM: `/api/v1/crm`
- Payroll: `/api/v1/payroll`
- Notifications: `/api/v1/notifications`

## Common Commands

```bash
# Run the app
make run

# Run tests
make test

# Build binary
make build

# Start with Docker
docker-compose -f deploy/docker-compose.local.yml up

# View logs
docker-compose -f deploy/docker-compose.local.yml logs -f

# Stop services
docker-compose -f deploy/docker-compose.local.yml down

# Clean up
make clean
```

## Troubleshooting

### Port Already in Use

```bash
# Check what's using port 8080
lsof -i :8080

# Kill the process or change APP_PORT in .env
```

### Database Connection Error

```bash
# Check PostgreSQL is running
docker ps | grep postgres

# Restart PostgreSQL
docker-compose -f deploy/docker-compose.local.yml restart postgres
```

### Redis Connection Error

```bash
# Check Redis is running
docker ps | grep redis

# Restart Redis
docker-compose -f deploy/docker-compose.local.yml restart redis
```

## Next Steps

1. Read [PROJECT_SETUP.md](PROJECT_SETUP.md) for detailed architecture
2. Implement remaining modules (users, meetings, CRM, payroll, notifications)
3. Add Google OAuth integration
4. Deploy to AWS EKS

## Need Help?

- Check logs: `docker-compose -f deploy/docker-compose.local.yml logs`
- Verify database: `docker exec -it evtaarpro-postgres psql -U postgres -d evtaarpro -c "\dt"`
- Test Redis: `docker exec -it evtaarpro-redis redis-cli ping`

Happy coding! 🚀
