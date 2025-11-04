# EvtaarPro - Architecture Documentation

## System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client Layer                             │
│              (Web App, Mobile App, Third-party APIs)             │
└────────────────────────────┬────────────────────────────────────┘
                             │ HTTPS
┌────────────────────────────▼────────────────────────────────────┐
│                      Load Balancer (AWS ELB)                     │
└────────────────────────────┬────────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────────┐
│                      API Gateway (Ingress)                       │
│                     nginx-ingress-controller                     │
└────────────────────────────┬────────────────────────────────────┘
                             │
┌────────────────────────────▼────────────────────────────────────┐
│                    EvtaarPro API Server (Go)                     │
│                                                                  │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │                  Middleware Layer                        │  │
│  │  • CORS • Authentication • Logging • Recovery            │  │
│  └──────────────────────────┬───────────────────────────────┘  │
│                             │                                   │
│  ┌──────────────────────────▼───────────────────────────────┐  │
│  │                    Router (Gin)                          │  │
│  │          /api/v1/{auth,users,meetings,crm,...}           │  │
│  └──────────────────────────┬───────────────────────────────┘  │
│                             │                                   │
│  ┌──────────────────────────▼───────────────────────────────┐  │
│  │                  Business Modules                        │  │
│  │  ┌──────┐ ┌──────────┐ ┌─────┐ ┌─────────┐ ┌──────────┐│  │
│  │  │ Auth │ │ Meetings │ │ CRM │ │ Payroll │ │  Notify  ││  │
│  │  └──┬───┘ └────┬─────┘ └──┬──┘ └────┬────┘ └────┬─────┘│  │
│  │     │          │           │         │            │      │  │
│  │  ┌──▼──────────▼───────────▼─────────▼────────────▼───┐ │  │
│  │  │          Domain Layer (Business Logic)            │ │  │
│  │  │  • Entities • Use Cases • Domain Services         │ │  │
│  │  └──────────────────────┬─────────────────────────────┘ │  │
│  │                         │                                │  │
│  │  ┌──────────────────────▼─────────────────────────────┐ │  │
│  │  │    Infrastructure Layer (Adapters)                 │ │  │
│  │  │  • Repositories • External Services • Clients      │ │  │
│  │  └────────────────────────────────────────────────────┘ │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────┬────────────────────────────────────┘
                          │
        ┌─────────────────┼──────────────────┐
        │                 │                  │
┌───────▼────────┐ ┌──────▼──────┐ ┌────────▼────────┐
│   PostgreSQL   │ │    Redis    │ │  External APIs   │
│    (RDS)       │ │ (ElastiCache)│ │ • Jitsi         │
│                │ │              │ │ • AWS S3        │
│ • Users        │ │ • Sessions   │ │ • Google OAuth  │
│ • Meetings     │ │ • Cache      │ │ • Firebase FCM  │
│ • CRM          │ │ • Rate Limit │ └─────────────────┘
│ • Payroll      │ │              │
│ • Notifications│ └──────────────┘
└────────────────┘
```

## Clean Architecture Layers

### 1. Presentation Layer (HTTP/WebSocket)
**Location**: `modules/{module}/presentation/`

**Responsibilities**:
- Handle HTTP requests and responses
- Validate input
- Map DTOs to domain entities
- Return standardized responses

**Example**:
```
modules/auth/presentation/http/
├── dto/                    # Data Transfer Objects
├── handlers/              # HTTP request handlers
└── routes/                # Route registration
```

### 2. Domain Layer (Business Logic)
**Location**: `modules/{module}/domain/`

**Responsibilities**:
- Define business entities
- Implement business rules
- Define interfaces (ports)
- Implement use cases

**Example**:
```
modules/auth/domain/
├── entities/              # Domain entities (User, Meeting, etc.)
├── ports/                 # Interfaces for external dependencies
├── services/              # Domain services
└── usecases/             # Application use cases
```

### 3. Infrastructure Layer
**Location**: `modules/{module}/infra/`

**Responsibilities**:
- Implement port interfaces
- Database access
- External API clients
- File storage
- Caching

**Example**:
```
modules/auth/infra/
├── postgresql/            # Database implementations
├── redis/                 # Cache implementations
└── security/              # Password hashing, JWT generation
```

### 4. Dependency Injection
**Location**: `modules/{module}/container.go`

**Responsibilities**:
- Wire dependencies
- Initialize use cases with implementations
- Register routes

## Module Structure

Each module follows this structure:

```
modules/{module}/
├── container.go           # DI container and route registration
├── domain/
│   ├── entities/          # Business entities
│   ├── ports/             # Interfaces
│   ├── services/          # Domain services (optional)
│   └── usecases/          # Use cases
├── infra/                 # Infrastructure implementations
│   ├── postgresql/
│   ├── redis/
│   └── {external-service}/
└── presentation/
    └── http/
        ├── dto/           # Request/Response DTOs
        ├── handlers/      # HTTP handlers
        └── routes/        # Route definitions
```

## Data Flow

### Request Flow (Example: User Login)

```
1. HTTP Request
   ↓
2. Middleware Chain
   - Request Logger
   - CORS
   - Recovery
   ↓
3. Router
   - Match route: POST /api/v1/auth/login
   ↓
4. Handler (Presentation Layer)
   - Validate input (DTO)
   - Call use case
   ↓
5. Use Case (Domain Layer)
   - Execute business logic
   - Call repositories through ports
   ↓
6. Repository (Infrastructure Layer)
   - Query database
   - Return domain entity
   ↓
7. Use Case
   - Process result
   - Return output
   ↓
8. Handler
   - Map to response DTO
   - Return HTTP response
```

## Database Schema

### Users & Auth
```sql
users
- id (PK)
- email (unique)
- password_hash
- first_name
- last_name
- role (admin|employee|client|hr)
- is_active
- email_verified
- created_at
- updated_at
```

### Meetings
```sql
meetings                    meeting_participants
- id (PK)                  - meeting_id (FK)
- room_id (unique)         - user_id (FK)
- title                    - role
- organizer_id (FK)        - joined_at
- start_time               - left_at
- end_time
- status
- jitsi_room_url
- recording_url
```

### CRM
```sql
customers                   customer_interactions
- id (PK)                  - id (PK)
- company_name             - customer_id (FK)
- email                    - user_id (FK)
- status                   - interaction_type
- assigned_to (FK)         - subject
- created_by (FK)          - notes
                           - interaction_date
```

### Payroll
```sql
employees                   payroll_records           attendance
- id (PK, FK)              - id (PK)                - id (PK)
- employee_code            - employee_id (FK)       - employee_id (FK)
- department               - period_start           - date
- designation              - period_end             - check_in
- salary_amount            - gross_salary           - check_out
- date_of_joining          - net_salary             - status
- bank_account             - status
```

### Notifications
```sql
notifications
- id (PK)
- user_id (FK)
- type
- title
- message
- data (jsonb)
- is_read
- priority
- created_at
```

## Security

### Authentication Flow

```
1. User registers → Password hashed with bcrypt → Stored in DB
2. User logs in → Password verified → JWT tokens generated
3. Access Token (15 min) + Refresh Token (7 days)
4. Session stored in Redis
5. Subsequent requests → Bearer token in header → JWT validated
6. User context injected into request
```

### Authorization

- **Role-Based Access Control (RBAC)**
  - Admin: Full access
  - HR: Payroll, employee management
  - Employee: Own profile, meetings, CRM
  - Client: Limited CRM access

- **Middleware**: `RequireRole(roles...)`

### Security Measures

- ✅ Password hashing with bcrypt
- ✅ JWT with expiration
- ✅ Session management with Redis
- ✅ CORS configuration
- ✅ Rate limiting (configured)
- ✅ SQL injection prevention (parameterized queries)
- ✅ XSS prevention (input validation)

## API Design

### RESTful Conventions

```
GET    /api/v1/resource        # List resources
GET    /api/v1/resource/:id    # Get single resource
POST   /api/v1/resource        # Create resource
PUT    /api/v1/resource/:id    # Update resource (full)
PATCH  /api/v1/resource/:id    # Update resource (partial)
DELETE /api/v1/resource/:id    # Delete resource
```

### Response Format

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Human readable message",
    "details": "Additional details"
  }
}
```

### Pagination

```json
{
  "success": true,
  "data": [...],
  "pagination": {
    "page": 1,
    "page_size": 20,
    "total_pages": 5,
    "total_items": 100
  }
}
```

## Error Handling

### Error Types

1. **Validation Errors** (400 Bad Request)
2. **Authentication Errors** (401 Unauthorized)
3. **Authorization Errors** (403 Forbidden)
4. **Not Found** (404 Not Found)
5. **Conflict** (409 Conflict)
6. **Server Errors** (500 Internal Server Error)

### Error Flow

```
Domain Error → Use Case → Handler → Response Helper → JSON Response
```

## Monitoring & Observability

### Metrics (Prometheus)

- Request count
- Request duration
- Error rate
- Database connection pool stats
- Redis cache hit/miss ratio

### Logging

- Structured JSON logging
- Request ID for tracing
- Log levels: DEBUG, INFO, WARN, ERROR
- Contextual information

### Health Checks

- `/health` - Basic health check
- `/ready` - Readiness check (DB + Redis)

## Deployment

### Local Development
```
Docker Compose → PostgreSQL + Redis + App
```

### Production (AWS EKS)
```
GitHub → GitHub Actions → ECR → EKS → Application
                ↓
         RDS (PostgreSQL)
         ElastiCache (Redis)
         S3 (File Storage)
         ALB (Load Balancer)
```

## Performance Considerations

1. **Connection Pooling**: PostgreSQL and Redis pools configured
2. **Caching**: Redis for sessions and frequently accessed data
3. **Indexes**: Database indexes on foreign keys and search columns
4. **Pagination**: All list endpoints support pagination
5. **Rate Limiting**: Prevent abuse
6. **Graceful Shutdown**: Clean connection closure

## Testing Strategy

1. **Unit Tests**: Test use cases in isolation with mocks
2. **Integration Tests**: Test with real PostgreSQL and Redis
3. **E2E Tests**: Test complete API flows
4. **Load Tests**: Performance testing with real workload

## Next Implementation Steps

### Priority 1: Core Modules
1. Users module (profile management)
2. Meetings module (Jitsi integration)

### Priority 2: Business Modules
3. CRM module (customer management)
4. Payroll module (salary management)

### Priority 3: Real-time Features
5. Notifications module (WebSocket)
6. Meeting real-time updates

### Priority 4: Advanced Features
7. Google OAuth integration
8. File uploads to S3
9. Email notifications
10. Analytics and reporting
