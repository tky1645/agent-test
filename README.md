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
