---
inclusion: always
---

# Technology Stack & Development Guidelines

## Tech Stack
- **Backend**: Go 1.23+ + Gin + MySQL + DDD/CQRS architecture
- **Frontend**: React 18 + TypeScript + Vite + Radix UI + Tailwind CSS
- **Testing**: testify/assert (Go), table-driven tests, 80% coverage minimum
- **Development**: Air (hot reload), golangci-lint, Docker Compose

## Code Standards

### Go
- Always check errors, use wrapped errors with context
- JSON/DB fields: snake_case, struct tags consistent
- Interfaces: small, single responsibility, `I` prefix
- Files: `camelCase.go`, Types: `PascalCase`

### TypeScript/React
- Strict mode enabled, avoid `any` type
- Functional components with hooks only
- Props interfaces explicitly defined
- Types must match backend entities exactly

## Development Commands
```bash
# Backend
air                    # Hot reload development
go test -cover ./...   # Run tests with coverage
golangci-lint run      # Lint code

# Frontend (in frontend/plant-watering-app/)
npm run dev           # Development server (port 5173)
npm run type-check    # TypeScript validation
npm run build         # Production build

# Docker
docker-compose up db backend              # Backend + DB only
cd frontend/plant-watering-app && docker-compose up  # Full stack
```

## Environment Variables
- `DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME`: MySQL connection
- `GIN_MODE`: "release" for production
- `PORT`: Server port (default: 8080)