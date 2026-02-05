# user-service

User management microservice for profiles and account operations.

## Features

- User profile management
- Account operations
- User search

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/users/:id` | Get user by ID |
| `GET` | `/api/v1/users/profile` | Get user profile |
| `PUT` | `/api/v1/users/profile` | Update profile |

## Tech Stack

- Go + Gin framework
- PostgreSQL 16 (supporting-db cluster)
- PgBouncer connection pooling
- OpenTelemetry tracing

## Development

```bash
go mod download
go test ./...
go run cmd/main.go
```

## License

MIT
