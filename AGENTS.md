# user-service

> AI Agent context for understanding this repository

## 📋 Overview

User management microservice. Handles user profiles and account operations.

## 🏗️ Architecture

```
user-service/
├── cmd/main.go
├── config/config.go
├── db/migrations/sql/
├── internal/
│   ├── core/
│   │   ├── database.go
│   │   └── domain/
│   ├── logic/v1/service.go
│   └── web/v1/handler.go
├── middleware/
└── Dockerfile
```

## 🔌 API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/users/:id` | Get user by ID |
| `GET` | `/api/v1/users/profile` | Get user profile |
| `PUT` | `/api/v1/users/profile` | Update user profile |
| `POST` | `/api/v1/users` | Create new user (internal) |

## 📐 3-Layer Architecture

| Layer | Location | Responsibility |
|-------|----------|----------------|
| **Web** | `internal/web/v1/handler.go` | HTTP handling, validation, error translation |
| **Logic** | `internal/logic/v1/service.go` | Business rules (❌ NO SQL) |
| **Core** | `internal/core/` | Domain models, repositories, database |

## 🗄️ Database

| Component | Value |
|-----------|-------|
| **Cluster** | supporting-db (Zalando Postgres Operator) |
| **PostgreSQL** | 16 |
| **HA** | Single instance |
| **Pooler** | PgBouncer Sidecar |
| **Endpoint** | `supporting-db-pooler.user.svc.cluster.local:5432` |
| **Pool Mode** | Transaction |
| **Shared DB** | Yes (with notification, shipping services) |

## 🚀 Graceful Shutdown

**VictoriaMetrics Pattern:**
1. `/ready` → 503 when `isShuttingDown = true`
2. Sleep `READINESS_DRAIN_DELAY` (5s)
3. Sequential: HTTP → Database → Tracer

## 🔧 Tech Stack

| Component | Technology |
|-----------|------------|
| **Framework** | Gin |
| **Database** | PostgreSQL 16 via pgx/v5 |
| **Logging** | Zap |
| **Tracing** | OpenTelemetry |
| **Metrics** | Prometheus |

## 🛠️ Development

```bash
go mod download && go test ./... && go build ./cmd/main.go
```
