# Northwind API

A RESTful API for the classic Northwind database, built with Go and [Gin](https://github.com/gin-gonic/gin). This project demonstrates clean architecture, modular design, JWT authentication, and structured logging using [zerolog](https://github.com/rs/zerolog).

## Features

- Customer and Employee endpoints (CRUD)
- JWT-based authentication (middleware)
- Swagger/OpenAPI documentation (docs/swagger.yaml)
- Structured logging to [`app.log`](app.log )
- Environment-based configuration
- Graceful shutdown
- SQLite database ([`northwind.db`](northwind.db ))

## Project Structure

```
.
├── main.go
├── go.mod
├── internal/
│   ├── config/         # App config & DB setup
│   ├── handlers/       # HTTP handlers
│   ├── logging/        # Logger setup & middleware
│   ├── middleware/     # Custom Gin middleware
│   ├── models/         # Data models & responses
│   ├── repositories/   # Data access layer
│   ├── routes/         # Route registration
│   ├── server/         # Gin engine setup
│   └── utils/          # Utilities (e.g., JWT)
├── docs/               # Swagger docs
├── northwind.db        # SQLite database
├── .env                # Environment variables
├── .air.toml           # Live reload config
└── tmp/                # Build artifacts
```

## Getting Started

### Prerequisites

- Go 1.18+
- [swag](https://github.com/swaggo/swag) (for Swagger docs)
- SQLite3

### Setup

1. **Clone the repo:**
   ```sh
   git clone https://github.com/yourusername/northwind-api.git
   cd northwind-api
   ```

2. **Copy and edit [`.env`](.env ):**
   ```sh
   cp .env.example .env
   # Edit as needed (JWT_SECRET, PORT, etc.)
   ```

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Generate Swagger docs:**
   ```sh
   swag init -g main.go -o ./docs
   ```

5. **Run the server:**
   ```sh
   go run main.go
   # or with live reload (requires air)
   air
   ```

## API Documentation

- Swagger UI: [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html)
- Health check: `GET /healthz`
- Versioned API: `GET /api/v1/...`

## Configuration

See [`internal/config/config.go`](internal/config/config.go ) for config loading. Main environment variables:

- `JWT_SECRET` (required)
- `PORT` (default: 8080)
- `DB_PATH` (default: northwind.db)
- `GO_ENV` (default: development)
- `API_VERSION` (default: v1)

## Logging

- Logs are written to [`app.log`](app.log ) using zerolog.
- HTTP requests are logged via middleware ([`internal/logging/logger.go`](internal/logging/logger.go )).

## License

MIT

---

**See also:**  
- [`main.go`](main.go ) – Entry point  
- [`internal/routes/routes.go`](internal/routes/routes.go ) – Route registration  
- [`internal/config/config.go`](internal/config/config.go ) – Configuration  
- [`docs/swagger.yaml`](docs/swagger.yaml ) – OpenAPI spec
