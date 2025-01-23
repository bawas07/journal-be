# Journaling BE

A REST API server built with Go and Fiber

## Development Setup

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Install Air for hot reload:
   ```bash
   go install github.com/air-verse/air@latest
   ```

3. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

4. Start the development server:
   ```bash
   air
   ```

The server will automatically reload when you make changes to .env or Go files.

## Configuration

Edit the `.env` file to configure the application.

## API Documentation

Coming soon...