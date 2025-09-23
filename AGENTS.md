# Repository Guidelines

agent-test combines a DDD-style Go API with a Vite/React frontend for plant-watering operations. Use this guide to keep contributions consistent and easy to review.

## Project Structure & Module Organization
- `main.go` wires Gin routes with command/query services.
- `command/` hosts application handlers, `entities/` the aggregates/value objects, and `query/` read-side adapters.
- `rdb/` abstracts MySQL access; `migrations/` tracks schema SQL applied to `sampledb`.
- `frontend/plant-watering-app/` contains the UI (components in `src/`, configs and diagrams nearby).
- Architecture notes, sequence diagrams, and Docker layouts live in `docs/` and `README.md`; refresh them when behaviour changes.
- Root `docker-compose.yml` runs backend+DB; the frontend folder ships a full-stack compose file.

## Build, Test, and Development Commands
- `go mod tidy` refreshes backend modules after dependency changes.
- `go run main.go` serves the API on `:8080` (exposed as 18080 in the backend compose stack).
- `go test ./...` covers all Go packages using Testify.
- `npm install` then `npm run dev` in `frontend/plant-watering-app/` starts Vite on `:3000`.
- `npm run build` creates the production bundle; `npm run lint` enforces ESLint rules.
- `docker compose up` at the root or inside the frontend folder provisions backend-only or full-stack environments; add `-d` for detached mode.

## Coding Style & Naming Conventions
- Format Go code with `gofmt`; exported items use PascalCase, internals camelCase, packages lower_snake.
- Keep interface naming aligned with the existing `IUserRepository`/`IPlantRepository` pattern.
- TypeScript components and hooks stay in PascalCase and camelCase respectively; keep styling declarative with Tailwind utilities.
- Run `npm run lint` or configure your editor to apply ESLint + Prettier on save.

## Testing Guidelines
- Store backend tests beside implementations in `_test.go`, favouring table-driven cases and `require` assertions.
- Mock repositories for unit coverage; use the compose MySQL plus `migrations/` for integration runs when needed.
- Reuse use-case identifiers (e.g., `UC-W1`) in test names or descriptions to show traceability.

## Commit & Pull Request Guidelines
- Follow the observed `type: short imperative summary` style (e.g., `feat: expose watering metrics`) and reference issue numbers or use-case codes.
- Keep commits small, rebased, and free of generated or vendored artefacts.
- PRs should include a short summary, test evidence (`go test`, `npm run lint`, relevant docker steps), and screenshots for UI updates.
- Confirm compose stacks still start when touching Docker, migrations, or startup scripts.

## Communication Guidelines
- 回答は日本語で行うこと。
- reasoningイベントも日本語で出力すること