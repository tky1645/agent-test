# agent-test
for ai agent. ex .Devin

# DDD Project

This project is a basic implementation of Domain-Driven Design (DDD) principles in Go.

## Project Structure

- `command/user`: Contains the user command handlers and service.
- `entities`: Defines the core entities (e.g., User).
- `rdb`: Contains the database related files.

## Dependencies

- `github.com/gin-gonic/gin`: Web framework.
- `github.com/go-sql-driver/mysql`: MySQL driver.

## How to Run

1.  Install dependencies: `go mod tidy`
2.  Run the application: `go run main.go`

## Sequence Diagrams

### GET /users/{id}

```mermaid
sequenceDiagram
    participant Client
    participant Handler
    participant UserService
    participant UserRepository
    participant Database

    Client->>Handler: GET /users/{id}
    Handler->>UserService: GetByID(id)
    UserService->>UserRepository: GetByID(id)
    UserRepository->>Database: SELECT * FROM users WHERE id = {id}
    Database-->>UserRepository: User data
    UserRepository-->>UserService: User data
    UserService-->>Handler: User data
    Handler-->>Client: User data
```

### POST /users

```mermaid
sequenceDiagram
    participant Client
    participant Handler
    participant UserService
    participant UserRepository
    participant Database

    Client->>Handler: POST /users
    Handler->>UserService: Create(id, name)
    UserService->>UserRepository: Save(user)
    UserRepository->>Database: INSERT INTO users (id, name) VALUES (?, ?)
    Database-->>UserRepository: OK
    UserRepository-->>UserService: OK
    UserService-->>Handler: OK
    Handler-->>Client: OK
```

### PUT /users/{id}

```mermaid
sequenceDiagram
    participant Client
    participant Handler
    participant UserService
    participant UserRepository
    participant Database

    Client->>Handler: PUT /users/{id}
    Handler->>UserService: Update(id, name)
    UserService->>UserRepository: GetByID(id)
    UserRepository->>Database: SELECT * FROM users WHERE id = {id}
    Database-->>UserRepository: User data
    UserRepository-->>UserService: User data
    UserService->>UserRepository: Save(user)
    UserRepository->>Database: UPDATE users SET name = ? WHERE id = ?
    Database-->>UserRepository: OK
    UserRepository-->>UserService: OK
    UserService-->>Handler: OK
    Handler-->>Client: OK
```

### DELETE /users/{id}

```mermaid
sequenceDiagram
    participant Client
    participant Handler
    participant UserService
    participant UserRepository
    participant Database

    Client->>Handler: DELETE /users/{id}
    Handler->>UserService: Delete(id)
    UserService->>UserRepository: Delete(id)
    UserRepository->>Database: DELETE FROM users WHERE id = {id}
    Database-->>UserRepository: OK
    UserRepository-->>UserService: OK
    UserService-->>Handler: OK
    Handler-->>Client: OK
```
>>>>>>> master
