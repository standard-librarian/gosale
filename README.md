# GoSale - Enterprise Go Service Foundation

<div align="center">
  <img src=".images/doddd.png" alt="GoSale Logo" width="200"/>
  
  [![Go Version](https://img.shields.io/badge/Go-1.25.5-blue.svg)](https://golang.org)
  [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
</div>

## 📖 Overview

GoSale is a production-ready, enterprise-grade foundation for building scalable Go services. It demonstrates best practices for building robust, observable, and secure web services in Go, featuring custom web framework implementation, comprehensive middleware support, JWT authentication with OPA (Open Policy Agent), and Kubernetes-native deployment.

This project follows a **deploy-first** philosophy, emphasizing containerization and Kubernetes deployment from day one, making it ideal for cloud-native development workflows.

## ✨ Features

### Core Capabilities

- **Custom Web Framework**: Built on top of `httptreemux` with middleware support for maximum flexibility
- **Authentication & Authorization**: 
  - JWT (JSON Web Token) based authentication using RSA key pairs
  - Policy-based authorization using Open Policy Agent (OPA) with embedded Rego policies
  - Secure key management with filesystem-based keystore
- **Comprehensive Middleware Stack**:
  - Request logging with trace IDs
  - Error handling and panic recovery
  - Metrics collection (requests, errors, panics, goroutines)
  - Authentication and authorization enforcement
- **Observability**:
  - Structured JSON logging with custom formatter
  - Built-in metrics endpoint (expvar)
  - Request tracing with unique trace IDs
  - Performance profiling via pprof
- **Production Ready**:
  - Graceful shutdown with load shedding
  - Configuration management with environment variables and flags
  - Health check endpoints (liveness/readiness)
  - CPU quota awareness with GOMAXPROCS tuning

### Kubernetes & Deployment

- **Container-First Design**: Multi-stage Docker builds with minimal Alpine-based images
- **Kubernetes Manifests**: 
  - Kustomize-based configuration management
  - Resource quotas and limits
  - Development and production overlays
  - Service and deployment definitions
- **Local Development**: Kind (Kubernetes in Docker) cluster configuration for local testing
- **Build Metadata**: Build version and date baked into binaries and container labels

## 🏗️ Architecture

### Project Structure

```
gosale/
├── app/
│   ├── services/
│   │   └── sales-api/          # Main HTTP API service
│   └── tooling/
│       ├── logfmt/             # Log formatting utility
│       └── sales-admin/        # Admin CLI for JWT generation
├── business/
│   └── web/                    # Business-level web framework
│       └── v1/
│           ├── auth/           # Authentication & authorization
│           ├── mid/            # HTTP middleware
│           └── response/       # HTTP response helpers
├── foundation/
│   ├── keystore/               # Cryptographic key management
│   ├── logger/                 # Structured logging
│   └── web/                    # Core web framework
├── zarf/
│   ├── docker/                 # Dockerfiles
│   ├── k8s/                    # Kubernetes manifests
│   │   ├── base/              # Base configurations
│   │   └── dev/               # Development overlays
│   └── keys/                   # Authentication keys
└── vendor/                     # Vendored dependencies
```

### Key Components

1. **Foundation Layer**: Core utilities (logging, web framework, key management)
2. **Business Layer**: Business logic, authentication, and middleware
3. **Application Layer**: Executable services and tools
4. **Infrastructure Layer**: Deployment configurations and scripts

## 🚀 Quick Start

### Prerequisites

- Go 1.25.5 or higher
- Docker
- kubectl
- Kind (for local Kubernetes)
- Make

### Local Development (Standalone)

```bash
# Run the service locally
make run

# View help and configuration options
make run-help

# Test the API
make curl

# Test authenticated endpoint (requires TOKEN env var)
export TOKEN="your-jwt-token"
make curl-auth

# Run load tests
make load
```

### Generate Authentication Tokens

```bash
# Run the admin tool to generate JWT tokens
make admin

# This will:
# - Generate RSA key pairs
# - Create JWT tokens with claims
# - Validate signatures
```

### Local Kubernetes Development

```bash
# Create a local Kind cluster
make dev-up

# Build the Docker image
make service

# Load image into Kind cluster
make dev-load

# Deploy to Kubernetes
make dev-apply

# View logs
make dev-logs

# Check deployment status
make dev-status

# Update after code changes
make dev-update

# Tear down the cluster
make dev-down
```

### View Metrics

```bash
# Run expvarmon to visualize service metrics
make metrics-view-sc
```

## 🔐 Authentication & Authorization

### JWT Token Generation

The project uses RSA key pairs for JWT signing and verification:

```bash
# Generate new RSA key pair
openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in private.pem -out public.pem
```

### OPA Policy Enforcement

Authorization is handled via Open Policy Agent with embedded Rego policies. The auth system supports:

- Token signature validation
- Claims verification
- Role-based access control
- Policy-based authorization rules

## 📊 Monitoring & Observability

### Structured Logging

All logs are output in structured JSON format with:
- Trace IDs for request correlation
- Timestamps
- Log levels
- Contextual information

Use the included `logfmt` tool to make logs human-readable:

```bash
go run app/services/sales-api/main.go | go run app/tooling/logfmt/main.go
```

### Metrics

The service exposes metrics via `/debug/vars` endpoint:
- Request counts
- Error counts
- Panic counts
- Goroutine counts
- Memory statistics

Access the debug server (default: `localhost:4000`) for:
- Metrics: `http://localhost:4000/debug/vars`
- CPU profile: `http://localhost:4000/debug/pprof/profile`
- Heap profile: `http://localhost:4000/debug/pprof/heap`

## 🛠️ Configuration

The service is configured via environment variables and command-line flags:

| Variable | Default | Description |
|----------|---------|-------------|
| `WEB_API_HOST` | `0.0.0.0:3000` | API server address |
| `WEB_DEBUG_HOST` | `0.0.0.0:4000` | Debug/metrics server address |
| `WEB_READ_TIMEOUT` | `5s` | HTTP read timeout |
| `WEB_WRITE_TIMEOUT` | `10s` | HTTP write timeout |
| `WEB_IDLE_TIMEOUT` | `120s` | HTTP idle timeout |
| `WEB_SHUTDOWN_TIMEOUT` | `20s` | Graceful shutdown timeout |
| `AUTH_KEYS_FOLDER` | `zarf/keys/` | Path to RSA keys |
| `AUTH_ACTIVE_KID` | | Active key ID for JWT signing |

## 🧪 Testing

```bash
# Run tests (when test infrastructure is added)
go test ./...

# Run with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...
```

## 📦 Dependencies

Key dependencies managed in `go.mod`:

- `github.com/ardanlabs/conf/v3` - Configuration management
- `github.com/dimfeld/httptreemux/v5` - HTTP router
- `github.com/golang-jwt/jwt/v4` - JWT implementation
- `github.com/google/uuid` - UUID generation
- `github.com/open-policy-agent/opa` - Policy engine

All dependencies are vendored for reproducible builds.

## 🗺️ Development Roadmap

Based on the project's evolution through issues and PRs:

### ✅ Completed

- [x] Kubernetes and Kind setup for deploy-first workflow (#1)
- [x] Docker containerization with multi-stage builds (#2)
- [x] Kubernetes deployment manifests and resource quotas (#3)
- [x] Service startup, shutdown, and configuration management (#5)
- [x] Custom HTTP router with middleware support (#7)
- [x] Comprehensive middleware for logging, errors, panics, and metrics (#9)
- [x] JWT and OPA integration for authentication and authorization (#11)
- [x] Secure keystore implementation (#13)

### 🎯 Future Enhancements

- Database integration (PostgreSQL)
- API versioning strategy
- Rate limiting and throttling
- Distributed tracing (OpenTelemetry)
- CI/CD pipeline configuration
- API documentation (OpenAPI/Swagger)
- End-to-end tests
- Performance benchmarks

## 👥 Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes with descriptive messages
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built with inspiration from Go best practices
- OPA integration for flexible authorization
- Kubernetes-native design principles
- Community-driven development

## 📚 Additional Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Open Policy Agent Documentation](https://www.openpolicyagent.org/docs/latest/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [JWT Best Practices](https://datatracker.ietf.org/doc/html/rfc8725)

---

**Built with ❤️ using Go**