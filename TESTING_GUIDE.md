# EvtaarPro - Complete Testing Guide

## 🚀 Quick Start (3 Steps)

### Step 1: Setup
```bash
cd evtaarpro
bash scripts/setup.sh
```

### Step 2: Run Application
```bash
make run
```

### Step 3: Test
```bash
curl http://localhost:8080/health
```

If you see `{"status":"healthy","service":"evtaarpro"}`, you're ready!

---

## 📋 Complete Test Flow

### 1. Health Check
```bash
curl http://localhost:8080/health
```

**Expected Response:**
```json
{
  "status": "healthy",
  "service": "evtaarpro"
}
```

### 2. Readiness Check
```bash
curl http://localhost:8080/ready
```

**Expected Response:**
```json
{
  "status": "ready",
  "postgres": "connected",
  "redis": "connected"
}
```

---

## 🔐 Authentication Flow

### 1. Register a User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@company.com",
    "password": "SecurePass123!",
    "first_name": "John",
    "last_name": "Doe",
    "role": "employee"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user_id": "uuid-here",
    "email": "john.doe@company.com",
    "role": "employee"
  }
}
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@company.com",
    "password": "SecurePass123!"
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "user_id": "uuid-here",
    "email": "john.doe@company.com",
    "role": "employee",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Important:** Save the `access_token` for subsequent requests!

### 3. Set Token Variable
```bash
# Replace with your actual token
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

---

## 👤 User Management

### 1. Get Current User Profile
```bash
curl http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "message": "User retrieved successfully",
  "data": {
    "id": "uuid-here",
    "email": "john.doe@company.com",
    "first_name": "John",
    "last_name": "Doe",
    "full_name": "John Doe",
    "role": "employee",
    "is_active": true,
    "email_verified": false,
    "created_at": "2025-11-04T10:00:00Z",
    "updated_at": "2025-11-04T10:00:00Z"
  }
}
```

### 2. Update User Profile
```bash
curl -X PUT http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe Updated",
    "phone": "+1234567890",
    "department": "Engineering"
  }'
```

### 3. List All Users
```bash
curl "http://localhost:8080/api/v1/users?page=1&page_size=20" \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid1",
      "email": "john.doe@company.com",
      "first_name": "John",
      "last_name": "Doe",
      "full_name": "John Doe",
      "role": "employee"
    }
  ],
  "pagination": {
    "page": 1,
    "page_size": 20,
    "total_pages": 1,
    "total_items": 1
  }
}
```

### 4. Search Users
```bash
curl "http://localhost:8080/api/v1/users?search=john" \
  -H "Authorization: Bearer $TOKEN"
```

### 5. Get Specific User
```bash
# Replace USER_ID with actual ID
curl http://localhost:8080/api/v1/users/USER_ID \
  -H "Authorization: Bearer $TOKEN"
```

---

## 📅 Meetings Management

### 1. Create a Meeting
```bash
curl -X POST http://localhost:8080/api/v1/meetings \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Team Standup",
    "description": "Daily standup meeting for the engineering team",
    "start_time": "2025-12-01T10:00:00Z",
    "max_participants": 50
  }'
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Meeting created successfully",
  "data": {
    "id": "meeting-uuid",
    "room_id": "room-uuid",
    "title": "Team Standup",
    "description": "Daily standup meeting for the engineering team",
    "organizer_id": "user-uuid",
    "start_time": "2025-12-01T10:00:00Z",
    "status": "scheduled",
    "max_participants": 50,
    "created_at": "2025-11-04T10:00:00Z"
  }
}
```

**Important:** Save the `id` for joining the meeting!

### 2. List Meetings
```bash
curl "http://localhost:8080/api/v1/meetings?page=1&page_size=20" \
  -H "Authorization: Bearer $TOKEN"
```

### 3. Get Meeting Details
```bash
# Replace MEETING_ID with actual ID
curl http://localhost:8080/api/v1/meetings/MEETING_ID \
  -H "Authorization: Bearer $TOKEN"
```

### 4. Join Meeting (Get Jitsi Token)
```bash
# Replace MEETING_ID with actual ID
curl -X POST http://localhost:8080/api/v1/meetings/MEETING_ID/join \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Joined meeting successfully",
  "data": {
    "meeting_id": "meeting-uuid",
    "room_url": "https://meet.jit.si/room-uuid",
    "token": "jwt-token-for-jitsi"
  }
}
```

**To Join the Meeting:**
1. Open the `room_url` in your browser
2. Use the `token` for authenticated access
3. Or integrate the token in your Jitsi iframe

---

## 🔓 Logout

```bash
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Authorization: Bearer $TOKEN"
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Logout successful"
}
```

---

## 📊 Complete Test Script

Save this as `test_api.sh`:

```bash
#!/bin/bash

BASE_URL="http://localhost:8080"

echo "🧪 Testing EvtaarPro API"
echo "======================="

# Health check
echo "\n1. Health Check"
curl -s $BASE_URL/health | jq

# Register user
echo "\n2. Registering User"
REGISTER_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "TestPass123!",
    "first_name": "Test",
    "last_name": "User",
    "role": "employee"
  }')
echo $REGISTER_RESPONSE | jq

# Login
echo "\n3. Logging In"
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "TestPass123!"
  }')
echo $LOGIN_RESPONSE | jq

TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.data.access_token')
echo "\n📝 Token: $TOKEN"

# Get user profile
echo "\n4. Getting User Profile"
curl -s $BASE_URL/api/v1/users/me \
  -H "Authorization: Bearer $TOKEN" | jq

# Update profile
echo "\n5. Updating Profile"
curl -s -X PUT $BASE_URL/api/v1/users/me \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Test",
    "last_name": "User Updated",
    "phone": "+1234567890",
    "department": "Engineering"
  }' | jq

# Create meeting
echo "\n6. Creating Meeting"
MEETING_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/meetings \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Meeting",
    "description": "This is a test meeting",
    "start_time": "2025-12-01T10:00:00Z"
  }')
echo $MEETING_RESPONSE | jq

MEETING_ID=$(echo $MEETING_RESPONSE | jq -r '.data.id')
echo "\n📝 Meeting ID: $MEETING_ID"

# List meetings
echo "\n7. Listing Meetings"
curl -s "$BASE_URL/api/v1/meetings?page=1&page_size=10" \
  -H "Authorization: Bearer $TOKEN" | jq

# Join meeting
echo "\n8. Joining Meeting"
curl -s -X POST $BASE_URL/api/v1/meetings/$MEETING_ID/join \
  -H "Authorization: Bearer $TOKEN" | jq

# Logout
echo "\n9. Logging Out"
curl -s -X POST $BASE_URL/api/v1/auth/logout \
  -H "Authorization: Bearer $TOKEN" | jq

echo "\n✅ All tests completed!"
```

**Run the script:**
```bash
chmod +x test_api.sh
./test_api.sh
```

---

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Check what's using port 8080
lsof -i :8080

# Kill the process or change APP_PORT in .env
```

### Database Connection Error
```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Check logs
docker logs evtaarpro-postgres

# Restart
docker-compose -f deploy/docker-compose.local.yml restart postgres
```

### Redis Connection Error
```bash
# Check if Redis is running
docker ps | grep redis

# Test connection
docker exec -it evtaarpro-redis redis-cli ping

# Should return: PONG
```

### JWT Token Expired
```bash
# Login again to get a new token
# Access tokens expire after 15 minutes by default
```

### Migration Errors
```bash
# Re-run migrations
bash scripts/init_db.sh

# Or manually
psql -h localhost -U postgres -d evtaarpro -f migrations/001_create_users_table.sql
```

---

## 🌐 Using Postman

1. Import the collection: `sparto-api.postman_collection.json` (if exists)
2. Or create manually:

**Setup Environment:**
```
base_url: http://localhost:8080
token: (leave empty initially)
```

**Auth Requests:**
1. Register → Save user details
2. Login → Save `access_token` to environment variable `token`
3. Use `{{token}}` in Authorization header for other requests

---

## 📈 Performance Testing

```bash
# Install hey
go install github.com/rakyll/hey@latest

# Test health endpoint
hey -n 1000 -c 10 http://localhost:8080/health

# Test authenticated endpoint (replace TOKEN)
hey -n 1000 -c 10 \
  -H "Authorization: Bearer TOKEN" \
  http://localhost:8080/api/v1/users/me
```

---

## 🎯 What's Working

✅ **Authentication**
- Register with email/password
- Login with JWT tokens
- Logout with session invalidation
- Token-based authentication

✅ **User Management**
- Get current user profile
- Update profile
- List all users
- Search users
- Get user by ID

✅ **Meetings**
- Create scheduled meetings
- List meetings
- Get meeting details
- Join meeting with Jitsi token

✅ **Security**
- Password hashing with bcrypt
- JWT with expiration (15min access, 7day refresh)
- Session management via Redis
- Role-based access control

---

## 📝 API Documentation

View Swagger documentation at:
```
http://localhost:8080/swagger/index.html
```

(Note: Run `make swagger` to generate docs if not present)

---

## 🎉 Success Criteria

Your API is working correctly if you can:

1. ✅ Register a new user
2. ✅ Login and receive JWT tokens
3. ✅ Get and update user profile
4. ✅ Create a meeting
5. ✅ Join a meeting and receive Jitsi token
6. ✅ Logout successfully

All 6 steps should work without errors!

---

## 🚀 Next Steps

1. **Implement Google OAuth** - Social login
2. **Add File Uploads** - Avatar images, meeting recordings
3. **Complete CRM Module** - Customer management
4. **Complete Payroll Module** - Salary and attendance
5. **Complete Notifications Module** - WebSocket real-time updates
6. **Add Tests** - Unit, integration, E2E tests
7. **Deploy to AWS** - EKS deployment

---

Happy Testing! 🎉
