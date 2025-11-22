# Simple Bank Statement Viewer Simulation

## Front End

### Project Structure

```
project-root/
└── src/
    ├── app
    ├── components/
    │   └── <component-name>/
    │       ├── index.tsx
    │       ├── index.test.tsx
    │       └── styles.module.css
    ├── features/
    │   └── <feature-name>/
    │        └── <components>/
    │            ├── index.tsx
    │            ├── index.test.tsx
    │            └── styles.module.css
    └── services/
        └── <service-name>/
            ├── api-sdk.ts
            └── api.types.ts
```

### Prerequisites

- Node.js 22 or higher
- npm 10 or higher
- pnpm 10 or higher

### Quick Start

#### Local Development

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <clone directory>
   ```

2. **Install dependencies**:
   ```bash
   pnpm install
   ```

3. **Setup environment**:
   ```bash
   cp .env.example .env
   # Edit values on the .env with the real value if needed
   ```

5. **Start the development**:
   ```bash
   pnpm run dev
   ```


### Testing

```bash
# Run all tests
pnpm test
```

### Build the Application

#### With Docker

```bash
docker compose up --build
```

#### Without Docker

```bash
pnpm build
```

## Back End

A 3-tier Go application built with net/http framework and structured logging using slog.

### Project Structure

```
project-root/
├── cmd/
│   ├── http_server/     # HTTP server entry point
│   │   └── main.go      # Main application with Swagger annotations
├── config/              # Configuration management
│   └── config.go        # Config loading with .env file support
├── feature/             # Feature-based modules
│   └── <feature_name>/  # Example feature
│       ├── repository/  # Data access layer
│       │   ├── repository.go
│       │   ├── repository_interface.go
│       │   ├── <method>.go
│       │   ├── <method>_dto.go
│       │   └── <method>_test.go
│       ├── service/     # Business logic layer
│       │   ├── service.go
│       │   ├── service_interface.go
│       │   ├── <method>.go
│       │   └── <method>_dto.go
│       │   └── <method>_test.go
│       └── handler/     # Presentation layer
│           ├── http.go  # Main HTTP handlers with Swagger docs
├── internal/            # Shared packages
│   ├── database/        # Database connection management
│   ├── logger/          # Structured logging utilities
│   └── http_server/     # Generic HTTP server with Swagger middleware
├── docs/                # Generated Swagger documentation
└── Makefile             # Build and development commands
```

### Prerequisites

- Go 1.25 or higher
- Make (optional, for build automation)

### Quick Start

#### Local Development

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <clone directory>
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Setup environment**:
   ```bash
   cp .env.example .env
   # Edit values on the .env with the real value if needed
   ```

4. **Generate API documentation**:
   ```bash
   make swagger
   ```

5. **Start the server**:
   ```bash
   make run
   ```

### Testing

```bash
# Run all tests
make test
```

### Build the Application

#### With Docker

```bash
docker compose up --build
```

#### Without Docker

```bash
make build-production
```


