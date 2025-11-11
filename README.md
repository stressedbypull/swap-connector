# SWAPI Connector

Go API service that connects to the Star Wars API (SWAPI) with pagination, search, sorting, and comprehensive documentation.

## Features

- Clean Architecture with Hexagonal/Ports and Adapters pattern
- SOLID Principles implementation
- OpenAPI 3.1 Swagger documentation
- Unit tests with mocks and integration tests against real SWAPI
- Search and sort capabilities
- Pagination support
- Security scanning with Trivy
- Docker support

## Quick Start

### Prerequisites

- Go 1.25 or higher
- Docker and Docker Compose
- swag CLI (optional, for regenerating docs): `go install github.com/swaggo/swag/cmd/swag@latest`

### Run Locally

```bash
git clone <your-repo-url>
cd MERCEDES

go mod download
make swagger
make run
```

The server will start on port 6969:
- API: http://localhost:6969/api/people
- Swagger UI: http://localhost:6969/swagger/index.html
- Health check: http://localhost:6969/ping
## API Documentation

### Endpoints

#### List People

```
GET /api/people?page=1&search=luke&sortBy=name&sortOrder=asc
```

Query Parameters:
- `page` (optional): Page number, default is 1
- `search` (optional): Search by name, case-insensitive
- `sortBy` (optional): Sort field - name, created, or mass
- `sortOrder` (optional): Sort order - asc or desc, default is asc

Examples:
```bash
curl http://localhost:6969/api/people
curl http://localhost:6969/api/people?search=luke
curl http://localhost:6969/api/people?sortBy=mass&sortOrder=desc
curl http://localhost:6969/api/people?page=2&sortBy=name
```

Response:
```json
{
  "count": 82,
  "page": 1,
  "pageSize": 15,
  "results": [
    {
      "name": "Luke Skywalker",
      "mass": 77,
      "created": "2014-12-09T13:50:51.644000Z",
      "films": ["https://swapi.dev/api/films/1/"]
    }
  ]
}
```

## Testing

```bash
go test ./... -short -v          # Unit tests with mocks
go test ./... -v                 # All tests including integration
make test-coverage               # Coverage report
make coverage-html               # HTML coverage report
```

Test structure:
- Unit tests in `*_test.go` files use mocks
- Integration tests in `integrational/` directory test against real SWAPI API

## Architecture

Project follows Clean Architecture and Hexagonal pattern:

```
cmd/server/main.go              - Entry point with Swagger annotations
internal/
  domain/                       - Core entities (Person, Planet, Pagination)
  ports/                        - Interfaces for Dependency Inversion
  adapters/
    http/                       - HTTP handlers, middleware, responses
    swapi/                      - SWAPI client implementation
  services/                     - Business logic layer
  sorting/                      - Sorting strategies (Strategy pattern)
  search/                       - Search and filtering logic
  errors/                       - Domain errors
  mocks/                        - Test mocks
docs/                           - Generated Swagger documentation
```

### SOLID Principles

- Single Responsibility: Each package has one clear purpose
- Open/Closed: New sorters can be added without modifying existing code
- Liskov Substitution: All sorters implement the same interface
- Interface Segregation: Small, focused interfaces
- Dependency Inversion: Services depend on interfaces, not implementations

## Available Commands

```bash
make run              # Run the server
make build            # Build binary
make test             # Run all tests
make test-coverage    # Run tests with coverage
make swagger          # Generate Swagger docs
make clean            # Clean build artifacts
make tidy             # Tidy dependencies
```

## Docker

```bash
docker compose up -d           # Start services
docker compose down            # Stop services
docker compose logs -f         # View logs
```

## CI/CD

The project includes GitHub Actions workflow with separate jobs:
- Test: Unit tests with coverage
- Build: Go binary and Docker image
- Scan: Security scanning with Trivy
- Deploy: Push to GitHub Container Registry (on version tags)