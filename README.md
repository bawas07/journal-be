# Journaling Backend

A backend service for journaling application.

## Getting Started

### Prerequisites

- Go 1.22+
- PostgreSQL

### Installation

1. Clone the repository
2. Run `go mod download` to install dependencies
3. Create `.env` file from `.env.example` and update the values
4. Start PostgreSQL server

### Running the Application

```bash
go run cmd/api/main.go
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
