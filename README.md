# DDD Project

This project is a basic implementation of Domain-Driven Design (DDD) principles in Go.

## Project Structure

- `command/user`: Contains the user command handlers and service.
- `entities`: Defines the core entities (e.g., User, Plant).
- `query/plant`: Contains the plant query handlers and repository.
- `rdb`: Contains the database related files.
- `migrations`: Database migration files.

## Dependencies

- `github.com/gin-gonic/gin`: Web framework.
- `github.com/go-sql-driver/mysql`: MySQL driver.

## How to Run

1.  Install dependencies: `go mod tidy`
2.  Run the application: `go run main.go`

## クラス図 (Class Diagram)

```mermaid
classDiagram
    class User {
        +int ID
        +userName Name
        +NewUser(id: int, name: string) User
    }
    
    class userName {
        +NewUserName(name: string) userName
    }
    
    class Plant {
        +int ID
        +PlantName Name
        +time.Time WateringDate
        +time.Time CreatedAt
        +time.Time UpdatedAt
        +NewPlant(name: string) Plant
        +UpdateWatering()
    }
    
    class PlantName {
        +NewPlantName(name: string) PlantName
    }
    
    class UserService {
        -IUserRepository userRepository
        +Create(id: int, name: string) error
        +Update(id: string, name: string) error
        +Delete(id: string) error
        +GetByID(id: string) User
    }
    
    class UserRepository {
        -sql.DB db
        +Save(user: User) error
        +GetByID(id: string) User
        +GetAll() []User
        +Delete(id: uint) error
        +Create(id: int) User
        +Update(id: string, name: string) error
    }
    
    class IUserRepository {
        <<interface>>
        +Create(id: int) User
        +Save(user: User) error
        +Update(id: string, name: string) error
        +GetByID(id: string) User
        +Delete(id: uint) error
    }
    
    class PlantRepository {
        -sql.DB db
        +create(plant: Plant) error
        +save(plant: Plant) error
        +findByID(id: int) Plant
        +FindAll(limit: int, offset: int) []Plant
    }
    
    class IPlantRepository {
        <<interface>>
        +create(Plant) error
        +save(Plant) error
        +findByID(int) Plant
        +FindAll(int, int) []Plant
    }
    
    class UserHandler {
        +HandlerGET(c: gin.Context)
        +HandlerPOST(c: gin.Context)
        +HandlerPUT(c: gin.Context)
        +HandlerGetByID(c: gin.Context)
        +HandlerDelete(c: gin.Context)
        +InitHandlers(db: sql.DB)
    }
    
    class PlantHandler {
        +HandlerPOST(c: gin.Context)
        +HandlerPATCH(c: gin.Context)
        +HandlerGETPlants(c: gin.Context)
    }
    
    User --> userName
    Plant --> PlantName
    UserService --> IUserRepository
    UserRepository ..|> IUserRepository
    PlantRepository ..|> IPlantRepository
    UserHandler --> UserService
    PlantHandler --> IPlantRepository
    UserRepository --> User
    PlantRepository --> Plant
```

## ユースケース一覧 (Use Cases)

### ユーザー管理
- **UC-U1**: ユーザー登録 - 新しいユーザーを作成する
- **UC-U2**: ユーザー情報取得 - 特定のユーザー情報を取得する
- **UC-U3**: ユーザー一覧取得 - 全ユーザーの一覧を取得する
- **UC-U4**: ユーザー情報更新 - ユーザーの名前を更新する
- **UC-U5**: ユーザー削除 - ユーザーを削除する

### 植物管理
- **UC-P1**: 植物登録 - 新しい植物を登録する
- **UC-P2**: 植物一覧取得 - ユーザーの植物一覧を取得する
- **UC-P3**: 植物詳細取得 - 特定の植物の詳細情報を取得する
- **UC-P4**: 植物情報更新 - 植物の情報を更新する
- **UC-P5**: 植物削除 - 植物を削除する

### 水やり管理
- **UC-W1**: 水やり記録 - 植物に水やりを行い記録する
- **UC-W2**: 水やり履歴取得 - 植物の水やり履歴を取得する
- **UC-W3**: 水やり状態確認 - 前回の水やりからの経過日数を確認する
- **UC-W4**: 水やり記録削除 - 水やり記録を削除する

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

## 植物管理のシーケンス図 (Plant Management Sequence Diagrams)

### POST /plants (植物登録)

```mermaid
sequenceDiagram
    participant Client
    participant PlantHandler
    participant PlantRepository
    participant Database

    Client->>PlantHandler: POST /plants
    PlantHandler->>PlantHandler: fetchPost(c)
    PlantHandler->>PlantHandler: validatePost(param)
    PlantHandler->>PlantHandler: entities.NewPlant(name)
    PlantHandler->>PlantRepository: create(plant)
    PlantRepository->>Database: INSERT INTO plant (name, watering_date, created_at, updated_at)
    Database-->>PlantRepository: OK
    PlantRepository-->>PlantHandler: OK
    PlantHandler-->>Client: Plant data
```

### PATCH /plants/{id} (水やり記録)

```mermaid
sequenceDiagram
    participant Client
    participant PlantHandler
    participant PlantRepository
    participant Database

    Client->>PlantHandler: PATCH /plants/{id}
    PlantHandler->>PlantHandler: fetchPatch(c)
    PlantHandler->>PlantHandler: validatePatch(param)
    PlantHandler->>PlantRepository: findByID(id)
    PlantRepository->>Database: SELECT * FROM plant WHERE id = {id}
    Database-->>PlantRepository: Plant data
    PlantRepository-->>PlantHandler: Plant data
    PlantHandler->>PlantHandler: plant.UpdateWatering()
    PlantHandler->>PlantRepository: save(plant)
    PlantRepository->>Database: UPDATE plant SET watering_date = ? WHERE id = ?
    Database-->>PlantRepository: OK
    PlantRepository-->>PlantHandler: OK
    PlantHandler-->>Client: Updated plant data
```

### GET /plants (植物一覧取得)

```mermaid
sequenceDiagram
    participant Client
    participant PlantHandler
    participant PlantRepository
    participant Database

    Client->>PlantHandler: GET /plants
    PlantHandler->>PlantHandler: ShouldBindJSON(req)
    PlantHandler->>PlantRepository: FindAll(limit, offset)
    PlantRepository->>Database: SELECT * FROM plant LIMIT ? OFFSET ?
    Database-->>PlantRepository: Plants data
    PlantRepository-->>PlantHandler: Plants data
    PlantHandler-->>Client: Plants list
```
