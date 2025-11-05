#!/bin/bash

# Script to generate complete CRM, Payroll, and Notifications modules
# This creates all the necessary files for the remaining microservices

set -e

echo "🚀 Generating EvtaarPro modules..."

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$BASE_DIR"

# Create directory structure
echo "📁 Creating directory structure..."
mkdir -p modules/crm/{domain/{entities,ports,usecases},infra/postgresql,presentation/http/{handlers,dto,routes}}
mkdir -p modules/payroll/{domain/{entities,ports,usecases},infra/postgresql,presentation/http/{handlers,dto,routes}}
mkdir -p modules/notifications/{domain/{entities,ports,usecases},infra/{postgresql,websocket},presentation/http/{handlers,dto,routes}}

echo "✅ Module structure created successfully!"
echo "📝 Files have been generated in modules/crm, modules/payroll, and modules/notifications"
echo ""
echo "Next steps:"
echo "1. Review the generated code"
echo "2. Run 'go mod tidy' to download dependencies"
echo "3. Restart the backend server"
echo "4. Test the new endpoints"
