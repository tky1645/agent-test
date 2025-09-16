
# DDDプロジェクト

このプロジェクトは、GoでDomain-Driven Design（DDD）の原則を基本的に実装したものです。

## プロジェクト構造

- `command/user`: ユーザーのコマンドハンドラーとサービスを含む
- `entities`: コアエンティティ（例：User、Plant）を定義
- `query/plant`: 植物のクエリハンドラーとリポジトリを含む
- `rdb`: データベース関連ファイルを含む
- `migrations`: データベースマイグレーションファイル

## 依存関係

- `github.com/gin-gonic/gin`: Webフレームワーク
- `github.com/go-sql-driver/mysql`: MySQLドライバー

## 実行方法

1.  依存関係をインストール: `go mod tidy`
2.  アプリケーションを実行: `go run main.go`

## クラス図

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

## ユースケース一覧

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

## シーケンス図

### GET /users/{id} (ユーザー取得)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant ハンドラー as Handler
    participant ユーザーサービス as UserService
    participant ユーザーリポジトリ as UserRepository
    participant データベース as Database

    クライアント->>ハンドラー: GET /users/{id}
    ハンドラー->>ユーザーサービス: GetByID(id)
    ユーザーサービス->>ユーザーリポジトリ: GetByID(id)
    ユーザーリポジトリ->>データベース: SELECT * FROM users WHERE id = {id}
    データベース-->>ユーザーリポジトリ: ユーザーデータ
    ユーザーリポジトリ-->>ユーザーサービス: ユーザーデータ
    ユーザーサービス-->>ハンドラー: ユーザーデータ
    ハンドラー-->>クライアント: ユーザーデータ
```

### POST /users (ユーザー作成)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant ハンドラー as Handler
    participant ユーザーサービス as UserService
    participant ユーザーリポジトリ as UserRepository
    participant データベース as Database

    クライアント->>ハンドラー: POST /users
    ハンドラー->>ユーザーサービス: Create(id, name)
    ユーザーサービス->>ユーザーリポジトリ: Save(user)
    ユーザーリポジトリ->>データベース: INSERT INTO users (id, name) VALUES (?, ?)
    データベース-->>ユーザーリポジトリ: OK
    ユーザーリポジトリ-->>ユーザーサービス: OK
    ユーザーサービス-->>ハンドラー: OK
    ハンドラー-->>クライアント: OK
```

### PUT /users/{id} (ユーザー更新)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant ハンドラー as Handler
    participant ユーザーサービス as UserService
    participant ユーザーリポジトリ as UserRepository
    participant データベース as Database

    クライアント->>ハンドラー: PUT /users/{id}
    ハンドラー->>ユーザーサービス: Update(id, name)
    ユーザーサービス->>ユーザーリポジトリ: GetByID(id)
    ユーザーリポジトリ->>データベース: SELECT * FROM users WHERE id = {id}
    データベース-->>ユーザーリポジトリ: ユーザーデータ
    ユーザーリポジトリ-->>ユーザーサービス: ユーザーデータ
    ユーザーサービス->>ユーザーリポジトリ: Save(user)
    ユーザーリポジトリ->>データベース: UPDATE users SET name = ? WHERE id = ?
    データベース-->>ユーザーリポジトリ: OK
    ユーザーリポジトリ-->>ユーザーサービス: OK
    ユーザーサービス-->>ハンドラー: OK
    ハンドラー-->>クライアント: OK
```

### DELETE /users/{id} (ユーザー削除)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant ハンドラー as Handler
    participant ユーザーサービス as UserService
    participant ユーザーリポジトリ as UserRepository
    participant データベース as Database

    クライアント->>ハンドラー: DELETE /users/{id}
    ハンドラー->>ユーザーサービス: Delete(id)
    ユーザーサービス->>ユーザーリポジトリ: Delete(id)
    ユーザーリポジトリ->>データベース: DELETE FROM users WHERE id = {id}
    データベース-->>ユーザーリポジトリ: OK
    ユーザーリポジトリ-->>ユーザーサービス: OK
    ユーザーサービス-->>ハンドラー: OK
    ハンドラー-->>クライアント: OK
```

## 植物管理のシーケンス図

### POST /plants (植物登録)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant 植物ハンドラー as PlantHandler
    participant 植物リポジトリ as PlantRepository
    participant データベース as Database

    クライアント->>植物ハンドラー: POST /plants
    植物ハンドラー->>植物ハンドラー: fetchPost(c)
    植物ハンドラー->>植物ハンドラー: validatePost(param)
    植物ハンドラー->>植物ハンドラー: entities.NewPlant(name)
    植物ハンドラー->>植物リポジトリ: create(plant)
    植物リポジトリ->>データベース: INSERT INTO plant (name, watering_date, created_at, updated_at)
    データベース-->>植物リポジトリ: OK
    植物リポジトリ-->>植物ハンドラー: OK
    植物ハンドラー-->>クライアント: 植物データ
```

### PATCH /plants/{id} (水やり記録)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant 植物ハンドラー as PlantHandler
    participant 植物リポジトリ as PlantRepository
    participant データベース as Database

    クライアント->>植物ハンドラー: PATCH /plants/{id}
    植物ハンドラー->>植物ハンドラー: fetchPatch(c)
    植物ハンドラー->>植物ハンドラー: validatePatch(param)
    植物ハンドラー->>植物リポジトリ: findByID(id)
    植物リポジトリ->>データベース: SELECT * FROM plant WHERE id = {id}
    データベース-->>植物リポジトリ: 植物データ
    植物リポジトリ-->>植物ハンドラー: 植物データ
    植物ハンドラー->>植物ハンドラー: plant.UpdateWatering()
    植物ハンドラー->>植物リポジトリ: save(plant)
    植物リポジトリ->>データベース: UPDATE plant SET watering_date = ? WHERE id = ?
    データベース-->>植物リポジトリ: OK
    植物リポジトリ-->>植物ハンドラー: OK
    植物ハンドラー-->>クライアント: 更新された植物データ
```

### GET /plants (植物一覧取得)

```mermaid
sequenceDiagram
    participant クライアント as Client
    participant 植物ハンドラー as PlantHandler
    participant 植物リポジトリ as PlantRepository
    participant データベース as Database

    クライアント->>植物ハンドラー: GET /plants
    植物ハンドラー->>植物ハンドラー: ShouldBindJSON(req)
    植物ハンドラー->>植物リポジトリ: FindAll(limit, offset)
    植物リポジトリ->>データベース: SELECT * FROM plant LIMIT ? OFFSET ?
    データベース-->>植物リポジトリ: 植物データ
    植物リポジトリ-->>植物ハンドラー: 植物データ
    植物ハンドラー-->>クライアント: 植物一覧
```
=======
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

## ローカルDockerアーキテクチャ構成

### パターン1: バックエンド + データベース構成

ルートディレクトリの `docker-compose.yml` を使用した構成です。

```bash
cd ~/repos/agent-test
docker-compose up
```

```mermaid
graph TB
    subgraph "Docker Compose - Backend Only"
        subgraph "ddd_app Container"
            A[Go Backend Application<br/>Port: 18080:8080]
        end
        
        subgraph "ddd_rdb Container"
            B[MySQL Database<br/>Port: 13306:3306<br/>DB: sampledb]
        end
        
        A --> B
    end
    
    C[External Client] --> A
    
    style A fill:#e1f5fe
    style B fill:#f3e5f5
```

**起動するコンテナ:**
- `agent-test_ddd_app_1`: Goバックエンドアプリケーション
- `agent-test_ddd_rdb_1`: MySQLデータベース

### パターン2: フルスタック構成

フロントエンドディレクトリの `docker-compose.yml` を使用した構成です。

```bash
cd ~/repos/agent-test/frontend/plant-watering-app
docker-compose up
```

```mermaid
graph TB
    subgraph "Docker Compose - Full Stack"
        subgraph "frontend Container"
            D[React TypeScript App<br/>Port: 3000:80<br/>Nginx Server]
        end
        
        subgraph "backend Container"
            E[Go Backend Application<br/>Port: 8080:8080]
        end
        
        subgraph "db Container"
            F[MySQL Database<br/>Port: 3306:3306<br/>DB: plantdb]
        end
        
        D --> E
        E --> F
    end
    
    G[External Client] --> D
    
    style D fill:#e8f5e8
    style E fill:#e1f5fe
    style F fill:#f3e5f5
```

**起動するコンテナ:**
- `plant-watering-app_frontend_1`: React TypeScriptフロントエンド
- `plant-watering-app_backend_1`: Goバックエンドアプリケーション
- `plant-watering-app_db_1`: MySQLデータベース

### コンテナ確認コマンド

```bash
# 起動中のコンテナを確認
docker ps

# コンテナのログを確認
docker-compose logs

# 特定のサービスのログを確認
docker-compose logs [service_name]
```

### 推奨構成

フルスタック開発には**パターン2**を使用することを推奨します。フロントエンドとバックエンドの連携テストが可能になります。
