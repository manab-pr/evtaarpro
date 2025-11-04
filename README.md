# EvtaarPro - AI-Powered Collaboration & Payroll Platform

## Overview

EvtaarPro is a unified workspace platform that combines video conferencing, CRM, and payroll management into a single scalable solution. Built with Go microservices and deployed on AWS EKS.

## Problem Statement

Startups and small companies often juggle multiple tools — one for meetings, one for CRM, another for payroll — leading to context switching, data duplication, and poor integration. EvtaarPro unifies these into a single workspace where teams can chat, meet, track clients, and manage salaries.

## Architecture

### Core Backend (Golang Microservices)

- **Tech Stack**: Go (Gin), gRPC, GraphQL, Docker, AWS EKS
- **Microservices**:
  - **User Service**: Authentication (JWT + OAuth2 Google), profile, roles
  - **Meet Service**: Integration with Jitsi APIs for meetings, recordings, live events
  - **CRM Service**: Manage leads, customers, pipelines, communication logs
  - **Payroll Service**: Salary calculations, attendance sync, payout logs
  - **Notification Service**: Real-time WebSocket events for reminders and updates

### Database Layer

- **PostgreSQL**: Relational data (users, CRM records, payroll)
- **Redis**: Caching JWT sessions and meeting tokens
- **S3**: Meeting recordings and file storage

### Infrastructure

- Dockerized microservices pushed to ECR
- Kubernetes/Helm charts for EKS deployment
- Nginx Ingress Controller
- Prometheus + Grafana for monitoring
- GitHub Actions for CI/CD

## Features

| Module | Key Features |
|--------|-------------|
| **Meetings** | Create/join Jitsi meetings, store recordings on S3 |
| **CRM** | Manage leads, customers, and communications |
| **Payroll** | Salary sheets, attendance sync, auto-generation of payslips |
| **Notifications** | WebSocket-based reminders and updates |
| **Infrastructure** | Go microservices, Docker, EKS, Prometheus/Grafana monitoring |

## Project Structure

```
evtaarpro/
├── cmd/                    # Application entry points
│   ├── server/            # Main API server
│   └── seeder/            # Database seeder
├── internal/              # Internal shared packages
│   ├── config/           # Configuration management
│   ├── middleware/       # HTTP middlewares
│   ├── datastore/        # Database connections
│   └── ...
├── modules/              # Business modules
│   ├── auth/            # Authentication & authorization
│   ├── users/           # User management
│   ├── meetings/        # Video conferencing
│   ├── crm/             # Customer relationship management
│   ├── payroll/         # Payroll management
│   └── notifications/   # Real-time notifications
├── pkg/                 # Reusable packages
│   ├── clients/         # External API clients
│   ├── jwt/             # JWT utilities
│   └── logger/          # Logging utilities
├── deploy/              # Deployment configurations
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── k8s/            # Kubernetes manifests
└── testing/            # Test utilities

```

## Getting Started

### Prerequisites

- Go 1.22+
- Docker & Docker Compose
- PostgreSQL 15+
- Redis 7+
- AWS Account (for EKS deployment)

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/evtaarpro.git
cd evtaarpro
```

2. Copy configuration files:
```bash
cp config/app.yaml.example config/app.yaml
cp config/postgres.yaml.example config/postgres.yaml
cp config/redis.yaml.example config/redis.yaml
```

3. Start dependencies with Docker Compose:
```bash
docker-compose -f deploy/docker-compose.local.yml up -d
```

4. Run database migrations:
```bash
make migrate-up
```

5. Start the server:
```bash
make run
```

The API will be available at `http://localhost:8080`

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run integration tests
make test-integration

# Run E2E tests
make test-e2e
```

### Building for Production

```bash
# Build binary
make build

# Build Docker image
make docker-build

# Push to registry
make docker-push
```

## API Documentation

API documentation is available via Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

## Deployment

### AWS EKS Deployment

1. Build and push Docker image:
```bash
make docker-build docker-push
```

2. Deploy to EKS:
```bash
kubectl apply -f deploy/k8s/
```

3. Access the application via Load Balancer URL

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Environment (dev/staging/prod) | `dev` |
| `APP_PORT` | Server port | `8080` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_NAME` | Database name | `evtaarpro` |
| `REDIS_HOST` | Redis host | `localhost` |
| `REDIS_PORT` | Redis port | `6379` |
| `JWT_SECRET` | JWT signing secret | - |
| `GOOGLE_CLIENT_ID` | Google OAuth client ID | - |
| `GOOGLE_CLIENT_SECRET` | Google OAuth secret | - |
| `JITSI_API_URL` | Jitsi server URL | - |
| `AWS_S3_BUCKET` | S3 bucket for recordings | - |

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

Your Name - your.email@example.com

Project Link: [https://github.com/yourusername/evtaarpro](https://github.com/yourusername/evtaarpro)
