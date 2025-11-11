# swap-connector

A Go API service with PostgreSQL database, containerized with Docker.

## 🚀 Quick Start

### Prerequisites
- Go 1.25+
- Docker & Docker Compose
- (Optional) Trivy for security scanning

### Setup

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd MERCEDES
   ```
1. **Structure of the project**
cmd/
└── server/
    └── main.go                    ← Entry point

internal/
├── domain/                        ← Core business logic (no dependencies!)
│   ├── person.go                  ← type Person struct
│   ├── planet.go                  ← type Planet struct
│   └── pagination.go              ← type PaginatedResponse[T]
│
├── ports/                         ← Interfaces (Dependency Inversion)
│   ├── repository.go              ← Repository interface
│   └── service.go                 ← Service interface
│
├── adapters/                      ← Implementations
│   ├── http/                      ← HTTP adapter (Gin)
│   │   ├── handlers/
│   │   │   ├── people.go          ← GET /people handlers
│   │   │   └── planets.go         ← GET /planets handlers
│   │   ├── middleware/
│   │   │   └── cors.go
│   │   └── response.go            ← Success/Error helpers
│   │
│   └── swapi/                     ← SWAPI adapter
│       ├── client.go              ← HTTP client
│       ├── mapper.go              ← DTO → Domain mapping
│       └── dto.go                 ← SWAPI response structs
│
├── services/                      ← Business logic (uses ports)
│   ├── people_service.go          ← Implements ports.PeopleService
│   └── planet_service.go          ← Implements ports.PlanetService
│
├── sorting/                       ← Strategy pattern (Open/Closed)
│   ├── sorter.go                  ← type Sorter interface
│   ├── by_name.go                 ← ByName sorter
│   ├── by_created.go              ← ByCreated sorter
│   └── factory.go                 ← Creates sorter from query param
│
├── pagination/                    ← Pagination logic
│   └── paginator.go               ← Paginate() function
│
├── search/                        ← Search logic
│   └── filter.go                  ← FilterByName() function
│
└── errors/                        ← Custom errors
    └── errors.go                  ← APIError type

config/
└── config.go                      ← Configuration

docker-compose.yml
Dockerfile
go.mod
Makefile                           ← make run, make test, make docker
README.md