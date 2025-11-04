#!/bin/bash
# Database initialization script for EvtaarPro

set -e

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_NAME="${DB_NAME:-evtaarpro}"

echo "🔧 Initializing EvtaarPro Database..."
echo "=================================="
echo "Host: $DB_HOST:$DB_PORT"
echo "Database: $DB_NAME"
echo "User: $DB_USER"
echo ""

# Check if PostgreSQL is running
echo "📡 Checking PostgreSQL connection..."
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -lqt | cut -d \| -f 1 | grep -qw postgres; then
    echo "❌ Error: Cannot connect to PostgreSQL"
    echo "   Please ensure PostgreSQL is running:"
    echo "   docker-compose -f deploy/docker-compose.local.yml up -d postgres"
    exit 1
fi
echo "✓ PostgreSQL is running"

# Create database if it doesn't exist
echo ""
echo "📦 Creating database if not exists..."
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || \
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -c "CREATE DATABASE $DB_NAME"
echo "✓ Database ready"

# Run migrations
echo ""
echo "🔄 Running migrations..."
for migration in migrations/*.sql; do
    echo "  Running $(basename $migration)..."
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$migration" > /dev/null
done
echo "✓ All migrations applied"

# Verify tables
echo ""
echo "📊 Verifying tables..."
TABLE_COUNT=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -t -c "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = 'public'")
echo "✓ Created $TABLE_COUNT tables"

echo ""
echo "✅ Database initialization complete!"
echo ""
echo "Next steps:"
echo "  1. Copy .env.example to .env and update your settings"
echo "  2. Run: make run"
echo "  3. Access API at http://localhost:8080"
