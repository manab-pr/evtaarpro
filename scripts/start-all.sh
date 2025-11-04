#!/bin/bash
# Start both backend and frontend for EvtaarPro

set -e

echo "🚀 Starting EvtaarPro Full Stack Application"
echo "==========================================="
echo ""

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if backend dependencies are ready
echo -e "${BLUE}📦 Checking backend...${NC}"
if [ ! -f ".env" ]; then
    echo "⚠️  No .env file found. Running setup..."
    bash scripts/setup.sh
else
    echo "✓ Backend configured"
fi

# Check if frontend dependencies are installed
echo ""
echo -e "${BLUE}📦 Checking frontend...${NC}"
if [ ! -d "frontend/node_modules" ]; then
    echo "Installing frontend dependencies..."
    cd frontend && npm install && cd ..
    echo "✓ Frontend dependencies installed"
else
    echo "✓ Frontend dependencies ready"
fi

# Start Docker services
echo ""
echo -e "${BLUE}🐳 Starting database services...${NC}"
docker-compose -f deploy/docker-compose.local.yml up -d postgres redis
echo "✓ PostgreSQL and Redis started"

# Wait for services
echo ""
echo -e "${BLUE}⏳ Waiting for services to be ready...${NC}"
sleep 5

# Start backend in background
echo ""
echo -e "${BLUE}🔧 Starting backend server...${NC}"
make run > logs/backend.log 2>&1 &
BACKEND_PID=$!
echo "✓ Backend starting (PID: $BACKEND_PID)"

# Wait a bit for backend to start
sleep 3

# Check if backend is running
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✓ Backend is healthy"
else
    echo "⏳ Backend is starting up..."
fi

# Start frontend in background
echo ""
echo -e "${BLUE}💻 Starting frontend server...${NC}"
cd frontend
npm run dev > ../logs/frontend.log 2>&1 &
FRONTEND_PID=$!
cd ..
echo "✓ Frontend starting (PID: $FRONTEND_PID)"

echo ""
echo -e "${GREEN}✅ EvtaarPro is starting up!${NC}"
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "🌐 Application URLs:"
echo "   Frontend:  http://localhost:3000"
echo "   Backend:   http://localhost:8080"
echo "   API Docs:  http://localhost:8080/swagger/index.html"
echo ""
echo "📊 Services:"
echo "   PostgreSQL: localhost:5432"
echo "   Redis:      localhost:6379"
echo ""
echo "📝 Process IDs:"
echo "   Backend:  $BACKEND_PID"
echo "   Frontend: $FRONTEND_PID"
echo ""
echo "📋 Logs:"
echo "   Backend:  tail -f logs/backend.log"
echo "   Frontend: tail -f logs/frontend.log"
echo ""
echo "🛑 To stop:"
echo "   kill $BACKEND_PID $FRONTEND_PID"
echo "   docker-compose -f deploy/docker-compose.local.yml down"
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Opening browser in 3 seconds..."
sleep 3

# Open browser (works on macOS, Linux, Windows with WSL)
if command -v open > /dev/null 2>&1; then
    open http://localhost:3000
elif command -v xdg-open > /dev/null 2>&1; then
    xdg-open http://localhost:3000
elif command -v start > /dev/null 2>&1; then
    start http://localhost:3000
fi

echo ""
echo "Press Ctrl+C to stop all services..."
echo ""

# Wait for user interrupt
trap "echo ''; echo 'Stopping services...'; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; docker-compose -f deploy/docker-compose.local.yml down; echo 'All services stopped.'; exit 0" INT

# Keep script running
wait
