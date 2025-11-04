# EvtaarPro - Complete Full Stack Setup Guide

## 🚀 Quick Start (3 Commands)

```bash
# 1. Setup everything
bash scripts/setup.sh

# 2. Start backend
make run

# 3. In a new terminal, start frontend
cd frontend && npm install && npm run dev
```

Then open `http://localhost:3000` in your browser!

---

## 📋 Complete Setup Instructions

### Prerequisites

**Required**:
- Go 1.22+
- Node.js 18+
- Docker & Docker Compose
- PostgreSQL 15+
- Redis 7+

**Optional**:
- Make
- Git

### Step 1: Clone and Setup Backend

```bash
cd evtaarpro

# Setup backend (installs dependencies, starts DB, runs migrations)
bash scripts/setup.sh

# This will:
# ✓ Check prerequisites
# ✓ Install Go dependencies
# ✓ Create .env file
# ✓ Start PostgreSQL and Redis
# ✓ Run database migrations
```

### Step 2: Setup Frontend

```bash
cd frontend

# Install dependencies
npm install

# This installs:
# ✓ React 18
# ✓ React Router
# ✓ TailwindCSS
# ✓ Axios
# ✓ Icons and utilities
```

### Step 3: Start Backend Server

```bash
# From evtaarpro directory
make run

# OR
go run cmd/server/main.go
```

**Backend will run on**: `http://localhost:8080`

**Check it's running**:
```bash
curl http://localhost:8080/health
```

### Step 4: Start Frontend Server

```bash
# Open a NEW terminal
cd evtaarpro/frontend

npm run dev
```

**Frontend will run on**: `http://localhost:3000`

---

## 🎯 Testing the Full Application

### 1. Open Browser

Navigate to: `http://localhost:3000`

### 2. Register a New User

1. Click "Sign up" on login page
2. Fill in details:
   - First Name: John
   - Last Name: Doe
   - Email: john.doe@company.com
   - Password: SecurePass123!
   - Role: Employee
3. Click "Create Account"

### 3. Login

1. Use the credentials you just created
2. Click "Sign In"
3. You'll be redirected to the Dashboard

### 4. Explore Dashboard

- View statistics
- See welcome message
- Check recent meetings
- Use quick actions

### 5. Update Your Profile

1. Click "Profile" in sidebar
2. Click "Edit Profile"
3. Update:
   - Phone number
   - Department
4. Click "Save Changes"

### 6. View Team Members

1. Click "Users" in sidebar
2. See all registered users
3. Try search functionality
4. View user details

### 7. Create a Meeting

1. Click "Meetings" in sidebar
2. Click "New Meeting" button
3. Fill in:
   - Title: "Team Standup"
   - Description: "Daily standup meeting"
   - Date & Time: Tomorrow at 10:00 AM
   - Max Participants: 50
4. Click "Create Meeting"

### 8. Join a Meeting

1. Go to Meetings page
2. Click on your meeting
3. Click "Join Meeting"
4. **Jitsi will open in a new window!**
5. Allow camera and microphone access
6. Start your video call!

---

## 🏗️ Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    Browser (localhost:3000)                  │
│                   React Frontend (Vite)                      │
└─────────────────────┬───────────────────────────────────────┘
                      │ HTTP/REST API
                      │ JWT Authentication
┌─────────────────────▼───────────────────────────────────────┐
│                Backend API (localhost:8080)                  │
│                    Go (Gin Framework)                        │
│                                                              │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐   │
│  │   Auth   │  │  Users   │  │ Meetings │  │  More... │   │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘   │
└──────────┬────────────────────┬──────────────┬─────────────┘
           │                    │              │
    ┌──────▼──────┐      ┌──────▼──────┐      │
    │ PostgreSQL  │      │    Redis    │      │
    │  (5432)     │      │   (6379)    │      │
    └─────────────┘      └─────────────┘      │
                                               │
                                         ┌─────▼──────┐
                                         │   Jitsi    │
                                         │ (External) │
                                         └────────────┘
```

---

## 🔌 API Integration

The frontend makes API calls to the backend:

### Authentication Flow

```
1. User registers → POST /api/v1/auth/register
2. User logs in → POST /api/v1/auth/login → Receives JWT token
3. Token stored in localStorage
4. All subsequent requests include: Authorization: Bearer <token>
5. Token validated by backend middleware
6. User data returned
```

### Meetings Flow

```
1. Create meeting → POST /api/v1/meetings
2. List meetings → GET /api/v1/meetings
3. View details → GET /api/v1/meetings/:id
4. Join meeting → POST /api/v1/meetings/:id/join
   ↓
5. Backend generates Jitsi JWT token
6. Returns room URL + token
7. Frontend opens Jitsi in new window
8. User joins video call!
```

---

## 🎨 Frontend Features

### Pages

1. **Login** (`/login`)
   - Email/password form
   - Link to registration
   - JWT token storage

2. **Register** (`/register`)
   - User signup form
   - Role selection
   - Validation

3. **Dashboard** (`/dashboard`)
   - Statistics cards
   - Recent meetings
   - Quick actions
   - Welcome message

4. **Profile** (`/profile`)
   - View user info
   - Edit profile
   - Account stats

5. **Users** (`/users`)
   - List all users
   - Search functionality
   - Pagination
   - User cards

6. **Meetings** (`/meetings`)
   - List meetings
   - Status indicators
   - Pagination
   - Create button

7. **Create Meeting** (`/meetings/new`)
   - Meeting form
   - Date/time picker
   - Validation

8. **Meeting Details** (`/meetings/:id`)
   - Full details
   - Join button
   - Meeting info
   - Tips & info

### Components

- **Layout**: Sidebar navigation, user info, logout
- **AuthContext**: Global auth state management
- **API Service**: Axios instance with interceptors
- **Toast Notifications**: Success/error messages

---

## 🛠️ Development Workflow

### Recommended Setup

**Terminal 1** - Backend:
```bash
cd evtaarpro
make run
```

**Terminal 2** - Frontend:
```bash
cd evtaarpro/frontend
npm run dev
```

**Terminal 3** - Database:
```bash
# View PostgreSQL logs
docker logs evtaarpro-postgres -f

# Or Redis logs
docker logs evtaarpro-redis -f
```

### Hot Reload

Both frontend and backend support hot reload:
- **Frontend**: Vite watches for changes, instant reload
- **Backend**: Use `make run` or restart manually

### Making Changes

1. **Backend Changes**:
   - Edit Go files
   - Restart `make run`
   - Test with `curl` or frontend

2. **Frontend Changes**:
   - Edit React components
   - Vite auto-reloads
   - See changes instantly

---

## 🐛 Troubleshooting

### Backend Issues

**Problem**: Port 8080 already in use
```bash
# Find and kill process
lsof -i :8080
kill -9 <PID>
```

**Problem**: Database connection failed
```bash
# Restart PostgreSQL
docker-compose -f deploy/docker-compose.local.yml restart postgres

# Check logs
docker logs evtaarpro-postgres
```

**Problem**: Redis connection failed
```bash
# Restart Redis
docker-compose -f deploy/docker-compose.local.yml restart redis

# Test connection
docker exec -it evtaarpro-redis redis-cli ping
```

### Frontend Issues

**Problem**: Port 3000 already in use
```bash
# Kill process on port 3000
lsof -i :3000
kill -9 <PID>

# Or use different port
npm run dev -- --port 3001
```

**Problem**: API calls failing (CORS)
- Backend must be running on port 8080
- Check proxy in `vite.config.js`
- Verify CORS settings in backend

**Problem**: Module not found
```bash
# Clean install
rm -rf node_modules package-lock.json
npm install
```

### Common Issues

**Issue**: JWT token expired (401 errors)
**Solution**: Login again to get new token

**Issue**: Cannot join meeting
**Solution**:
1. Check backend logs
2. Verify Jitsi configuration in .env
3. Test API endpoint: `curl -X POST http://localhost:8080/api/v1/meetings/:id/join -H "Authorization: Bearer TOKEN"`

---

## 📊 Testing Checklist

Use this checklist to verify everything works:

### ✅ Backend
- [ ] Health check responds (`curl http://localhost:8080/health`)
- [ ] Can register user
- [ ] Can login user
- [ ] Can get user profile
- [ ] Can create meeting
- [ ] Can list meetings

### ✅ Frontend
- [ ] Login page loads
- [ ] Can register new user
- [ ] Can login
- [ ] Dashboard shows data
- [ ] Can view profile
- [ ] Can edit profile
- [ ] Users page shows list
- [ ] Search works
- [ ] Can create meeting
- [ ] Can view meeting details
- [ ] Join meeting opens Jitsi

### ✅ Integration
- [ ] Frontend connects to backend
- [ ] JWT authentication works
- [ ] API calls succeed
- [ ] Toasts show success/errors
- [ ] Navigation works
- [ ] Logout works

---

## 🚀 Production Deployment

### Backend

```bash
# Build binary
make build

# Run binary
./bin/evtaarpro

# Or with Docker
docker build -t evtaarpro-backend -f deploy/Dockerfile .
docker run -p 8080:8080 evtaarpro-backend
```

### Frontend

```bash
cd frontend

# Build for production
npm run build

# Serve dist folder
npm run preview

# Or deploy to Netlify/Vercel
netlify deploy --prod --dir=dist
```

---

## 📦 Full Stack in Docker

To run everything in Docker:

```bash
# Add frontend to docker-compose.local.yml
# Then run:
docker-compose -f deploy/docker-compose.local.yml up -d

# Access:
# Frontend: http://localhost:3000
# Backend: http://localhost:8080
# PostgreSQL: localhost:5432
# Redis: localhost:6379
```

---

## 🎓 Learning Resources

- **Backend**: Go, Gin, PostgreSQL, Redis, JWT
- **Frontend**: React, TailwindCSS, React Router, Axios
- **Video**: Jitsi Meet API
- **Architecture**: Clean Architecture, DDD

---

## 🎉 Success!

If you can:
1. ✅ Open `http://localhost:3000`
2. ✅ Register and login
3. ✅ View dashboard
4. ✅ Create a meeting
5. ✅ Join meeting (Jitsi opens)

**Congratulations! Your full-stack application is working perfectly! 🚀**

---

## 📝 Next Steps

1. **Customize**: Change colors, add your logo
2. **Extend**: Add CRM, Payroll, Notifications modules
3. **Deploy**: Push to production on AWS EKS
4. **Test**: Write unit and integration tests
5. **Monitor**: Set up Prometheus and Grafana

---

**Enjoy building with EvtaarPro! 🎯**
