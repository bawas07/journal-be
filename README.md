# Journaling Backend

A backend service for journaling application.

## Project Structure

```
mindscribe-be/
├── cmd/                  # Main application commands
│   ├── api/              # API server entry point
│   └── migrate/          # Database migration commands
├── internal/             # Internal application code
│   ├── handler/          # HTTP request handlers
│   ├── models/           # Database models
│   ├── repository/       # Database repository layer
│   └── service/          # Business logic services
├── migrations/           # Database migration files
├── pkg/                  # Shared packages
│   ├── config/           # Configuration management
│   ├── logger/           # Logging utilities
│   ├── response/         # API response handling
│   ├── route/            # HTTP routing
│   └── server/           # HTTP server setup
├── .air.toml             # Live reload configuration
├── .env.example          # Environment variables template
├── dockerfile            # Docker configuration
├── go.mod                # Go module dependencies
└── go.sum                # Go dependency checksums
```

## Getting Started

### Prerequisites

- Go 1.22+
- PostgreSQL

### Installation

1. Clone the repository
2. Run `go mod download` to install dependencies
3. Create `.env` file from `.env.example` and update the values
4. Start PostgreSQL server

## Tutorial

### Setting Up the Project

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/mindscribe-be.git
   cd mindscribe-be
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run database migrations:
   ```bash
   go run cmd/migrate/main.go -cmd up
   ```

### Running the Application

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

### Testing the API

1. Create a new user:
   ```bash
   curl -X POST http://localhost:8080/api/v1/users \
   -H "Content-Type: application/json" \
   -d '{"name":"John Doe","email":"john@example.com","password":"password"}'
   ```

2. Login to get access token:
   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/login \
   -H "Content-Type: application/json" \
   -d '{"email":"john@example.com","password":"password"}'
   ```

3. Create a journal entry:
   ```bash
   curl -X POST http://localhost:8080/api/v1/journals \
   -H "Content-Type: application/json" \
   -H "Authorization: Bearer <your_token>" \
   -d '{"title":"My First Entry","content":"This is my first journal entry"}'
   ```

### Database Migrations

Migrations are managed using the `migrate` CLI tool.

#### Applying Migrations

```bash
go run cmd/migrate/main.go -cmd up
```

#### Rolling Back Migrations

```bash
go run cmd/migrate/main.go -cmd down
```

#### Applying/Rolling Back Specific Number of Migrations

```bash
go run cmd/migrate/main.go -cmd up -steps 2
go run cmd/migrate/main.go -cmd down -steps 1
```

#### Creating New Migrations

To create a new migration:

```bash
go run cmd/migrate/main.go -cmd create -name <migration_name>
```

This will create two files in the migrations directory:

-   `<timestamp>_<migration_name>.up.sql`
-   `<timestamp>_<migration_name>.down.sql`

Edit these files to add your migration SQL statements.

### Migration File Naming Convention

-   Files must follow the pattern: `<timestamp>_<description>.up.sql` and `<timestamp>_<description>.down.sql`
-   Timestamp format: YYYYMMDDHHMMSS
-   Description should be short and descriptive using snake_case
