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

2. **Configure environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your values
   ```

   **Required variables:**
   - `DB_USER` - PostgreSQL username
   - `DB_PASSWORD` - PostgreSQL password
   - `DB_NAME` - Database name
   - `DB_PORT` - Database port (default: 5432)

3. **Run with Docker Compose**
   ```bash
   make compose-up
   ```

4. **Test the API**
   ```bash
   curl http://localhost:6969/ping
   # Response: {"status":"healthy","message":"pong"}
   ```

## 📋 Available Commands

### Development
```bash
make run              # Run locally (macOS ARM64)
make debug            # Run with race detector
```

### Docker
```bash
make compose-up       # Start all services
make compose-down     # Stop all services
make compose-logs     # View logs
make compose-rebuild  # Rebuild and restart
```

### Security Scanning
```bash
make security-scan    # Scan filesystem for vulnerabilities
make security-docker  # Scan Docker image
make security-secrets # Scan for secrets/credentials
```

### Maintenance
```bash
make tidy             # Update dependencies
make clean            # Remove build artifacts
```

## 🏗️ Architecture

This project follows Clean Architecture principles with:
- **API Layer** - HTTP handlers (Gin framework)
- **Service Layer** - Business logic
- **Repository Layer** - Data access
- **Domain Layer** - Core entities

## 🔒 Security

- **Trivy** scans for vulnerabilities (configured in `trivy.yaml`)
- Secrets detection enabled
- `.env` files are git-ignored

## 🚢 Deployment (Coming Soon)

- **Kubernetes** - Container orchestration
- **Terraform** - Infrastructure as Code
- **CI/CD** - GitHub Actions

---

## 📖 Infrastructure Overview

```
Developer
    ↓
  git push
    ↓
┌─────────────────────────────────────┐
│      GitHub Actions (CI/CD)         │
│  1. Security scan (Trivy)           │
│  2. Build Docker image              │
│  3. Push to registry                │
└─────────────────┬───────────────────┘
                  ↓
┌─────────────────────────────────────┐
│         Terraform (IaC)             │
│  Creates:                           │
│  • Kubernetes cluster               │
│  • VPC/Network                      │
│  • RDS PostgreSQL                   │
│  • Load Balancers                   │
└─────────────────┬───────────────────┘
                  ↓
┌─────────────────────────────────────┐
│      Kubernetes Cluster             │
│                                     │
│  ┌───────────────────────────────┐ │
│  │  Deployment (3 pods)          │ │
│  │  ┌─────┐ ┌─────┐ ┌─────┐     │ │
│  │  │ API │ │ API │ │ API │     │ │
│  │  └─────┘ └─────┘ └─────┘     │ │
│  └──────────────┬────────────────┘ │
│                 │                  │
│  ┌──────────────┴────────────────┐ │
│  │  Service (Load Balancer)      │ │
│  └──────────────┬────────────────┘ │
│                 │                  │
│  ┌──────────────┴────────────────┐ │
│  │  Ingress (External Access)    │ │
│  └──────────────┬────────────────┘ │
└─────────────────┼───────────────────┘
                  ↓
            Internet Users
```

### Component Breakdown

**Docker** 🐳
- Packages app + dependencies → Container image
- Ensures consistency across environments

**Kubernetes** ☸️
- **Pod**: Runs your container
- **Deployment**: Manages 3 replicas for high availability
- **Service**: Internal load balancer (stable IP)
- **Ingress**: Routes external traffic to services

**Terraform** 🏗️
- Declarative infrastructure as code
- Creates cloud resources (cluster, database, networking)
- Version controlled, reproducible

**Flow:**
```
Code → CI (build + scan) → Registry → K8s pulls image → Pods serve traffic
```

## 🤝 Contributing

1. Create a feature branch
2. Make changes
3. Run security scans
4. Submit pull request

## 📄 License

[Your License Here]
