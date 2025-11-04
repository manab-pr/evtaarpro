#!/bin/bash
# Complete setup script for EvtaarPro

set -e

echo "🚀 EvtaarPro Setup Script"
echo "========================="
echo ""

# Check prerequisites
echo "📋 Checking prerequisites..."

# Check Go
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.22+"
    exit 1
fi
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✓ Go $GO_VERSION installed"

# Check Docker
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker"
    exit 1
fi
echo "✓ Docker installed"

# Check Docker Compose
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed"
    exit 1
fi
echo "✓ Docker Compose installed"

echo ""
echo "📦 Installing Go dependencies..."
go mod download
echo "✓ Dependencies installed"

echo ""
echo "🔧 Setting up environment..."
if [ ! -f .env ]; then
    cp .env.example .env
    echo "✓ Created .env file from template"
    echo "⚠️  Please update .env with your actual credentials"
else
    echo "✓ .env file already exists"
fi

echo ""
echo "🐳 Starting Docker services..."
docker-compose -f deploy/docker-compose.local.yml up -d postgres redis
echo "✓ PostgreSQL and Redis started"

echo ""
echo "⏳ Waiting for services to be ready..."
sleep 5

echo ""
echo "💾 Initializing database..."
bash scripts/init_db.sh

echo ""
echo "✅ Setup complete!"
echo ""
echo "To start the application:"
echo "  make run"
echo ""
echo "To view API documentation:"
echo "  Open http://localhost:8080/swagger/index.html"
echo ""
echo "To test the API:"
echo "  curl http://localhost:8080/health"
echo ""
