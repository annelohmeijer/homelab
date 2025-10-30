# Go CRUD API

A simple CRUD API built with Go following Domain-Driven Design (DDD) and the service repository pattern.

## Architecture

The project follows clean architecture principles with clear separation of concerns:

```
internal/
├── domain/          # Domain models and repository interfaces
├── repository/      # Repository implementations (PostgreSQL)
├── service/         # Business logic layer
├── handler/         # HTTP handlers (adapters)
└── container/       # Dependency injection
```

## Domain Model

**Project**:
- ID (UUID)
- Name (string)
- Description (string)
- ExpirationDate (timestamp)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)

## Setup

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- goose (for migrations)

### Installation

1. Install goose CLI:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

2. Start the database:
```bash
make db-up
```

3. Run migrations:
```bash
make migrate-up
```

4. Run the server:
```bash
make run
```

The API will be available at `http://localhost:8080`

## API Documentation

Swagger documentation is available at `http://localhost:8080/swagger/index.html`

## API Endpoints

### Health Check
- `GET /health` - Health check endpoint

### Projects
- `POST /api/v1/projects` - Create a new project
- `GET /api/v1/projects` - Get all projects
- `GET /api/v1/projects/:id` - Get a project by ID
- `PUT /api/v1/projects/:id` - Update a project
- `DELETE /api/v1/projects/:id` - Delete a project

### Example Requests

**Create Project**:
```bash
curl -X POST http://localhost:8080/api/v1/projects \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Project",
    "description": "A test project",
    "expiration_date": "2026-12-31T23:59:59Z"
  }'
```

**Get All Projects**:
```bash
curl http://localhost:8080/api/v1/projects
```

**Get Project by ID**:
```bash
curl http://localhost:8080/api/v1/projects/{id}
```

**Update Project**:
```bash
curl -X PUT http://localhost:8080/api/v1/projects/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Project",
    "description": "Updated description",
    "expiration_date": "2027-12-31T23:59:59Z"
  }'
```

**Delete Project**:
```bash
curl -X DELETE http://localhost:8080/api/v1/projects/{id}
```

## Database

PostgreSQL is used as the database. Connection details:
- Host: localhost
- Port: 5432
- Database: homelab
- User: postgres
- Password: postgres

PgAdmin is available at `http://localhost:5050`:
- Email: admin@admin.com
- Password: admin

## Migrations

Migrations are managed using goose:

```bash
make migrate-up      # Run all migrations
make migrate-down    # Rollback last migration
make migrate-status  # Check migration status
make migrate-create  # Create a new migration
```

## Development

Run the server interactively with hot reload:
```bash
make run
```

Build the binary:
```bash
make build
```

Generate swagger docs (after changing API):
```bash
make swagger
```
