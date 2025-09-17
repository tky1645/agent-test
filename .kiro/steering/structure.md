---
inclusion: always
---

# Architecture & Project Structure

## DDD/CQRS Architecture

### Layer Organization
- **Domain** (`/entities/`): Core entities (`User`, `Plant`, `WateringRecord`) with business rules, no external dependencies
- **Commands** (`/command/`): Write operations, business logic, HTTP handlers, repository interfaces
- **Queries** (`/query/`): Read operations optimized for UI, view models, data access implementations
- **Infrastructure**: Database, HTTP routing, external services

### Directory Structure
```
Backend:
├── entities/          # Domain models
├── command/{aggregate}/  # Write operations (user/, plant/)
├── query/{usecase}/   # Read operations (plant/)
├── migrations/        # Versioned SQL schema
└── main.go

Frontend (src/):
├── components/        # Reusable UI components
│   └── ui/           # Radix UI primitives (shadcn/ui)
├── contexts/         # React Context providers
├── hooks/            # Custom hooks for data/state
├── services/         # Typed API clients
├── types/            # TypeScript definitions
└── lib/              # Utilities
```

## Naming Conventions

### Go
- Files: `camelCase.go` (userService.go)
- Types: `PascalCase` (UserService)
- Interfaces: `I` prefix (IUserRepository)
- Methods: `PascalCase` public, `camelCase` private
- Constants: `UPPER_SNAKE_CASE`

### TypeScript
- Components: `PascalCase.tsx` (PlantCard.tsx)
- Hooks: `use-kebab-case.ts` (use-plant-data.ts)
- Services: `camelCase.ts` (plantApi.ts)
- Types: `PascalCase` interfaces matching backend

## Key Patterns
- **Repository**: Interfaces in application layer, separate command/query repositories
- **CQRS**: Commands for business logic, queries for UI optimization
- **Error Handling**: Wrapped errors with context (Go), typed responses (TypeScript)
- **Dependency Injection**: For testability and loose coupling