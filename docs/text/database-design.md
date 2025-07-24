# ミニマム植物水やり管理アプリ - データベース設計

## テーブル構成

### 1. users テーブル
ユーザー情報を管理するテーブル。認証はAWS Cognitoを使用。

| カラム名 | データ型 | 説明 | 制約 |
|---------|----------|------|------|
| id | UUID | ユーザーID | PRIMARY KEY |
| cognito_id | VARCHAR(255) | Cognito User ID | NOT NULL, UNIQUE |
| email | VARCHAR(255) | メールアドレス | NOT NULL, UNIQUE |
| name | VARCHAR(100) | ユーザー名 | NOT NULL |
| created_at | TIMESTAMP | 作成日時 | NOT NULL, DEFAULT CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | 更新日時 | NOT NULL, DEFAULT CURRENT_TIMESTAMP |

### 2. plants テーブル
ユーザーが登録した植物の情報を管理するテーブル。

| カラム名 | データ型 | 説明 | 制約 |
|---------|----------|------|------|
| id | UUID | 植物ID | PRIMARY KEY |
| user_id | UUID | 所有ユーザーID | FOREIGN KEY (users.id) |
| name | VARCHAR(100) | 植物の名前 | NOT NULL |
| description | TEXT | 説明・メモ | NULL |
| image_url | VARCHAR(255) | 植物の画像URL | NULL |
| created_at | TIMESTAMP | 作成日時 | NOT NULL, DEFAULT CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | 更新日時 | NOT NULL, DEFAULT CURRENT_TIMESTAMP |

### 3. watering_records テーブル
植物の水やり記録を管理するテーブル。前回の水やり日と間隔を計算するための履歴データ。

| カラム名 | データ型 | 説明 | 制約 |
|---------|----------|------|------|
| id | UUID | 記録ID | PRIMARY KEY |
| plant_id | UUID | 植物ID | FOREIGN KEY (plants.id) |
| watered_at | TIMESTAMP | 水やり実施日時 | NOT NULL |
| notes | TEXT | 特記事項 | NULL |
| created_at | TIMESTAMP | 作成日時 | NOT NULL, DEFAULT CURRENT_TIMESTAMP |

## ERD (Entity Relationship Diagram)

```
users 1--* plants (ユーザーは複数の植物を持つ)
plants 1--* watering_records (植物は複数の水やり記録を持つ)
```

## インデックス設計

### users テーブル
- `idx_users_cognito_id` - Cognitoでの認証に使用

### plants テーブル
- `idx_plants_user_id` - ユーザー別の植物リスト取得に使用

### watering_records テーブル
- `idx_wr_plant_id` - 植物IDからの水やり記録検索に使用
- `idx_wr_watered_at` - 日付順の水やり記録検索に使用

## 計算値（アプリケーション側で計算）

### 前回の水やり間隔（日）
最新の2つの水やり記録の日数差を計算
```sql
-- 例：植物IDが {plant_id} の前回の水やり間隔を計算するクエリ
SELECT
  EXTRACT(EPOCH FROM (w1.watered_at - w2.watered_at)) / 86400 AS days_between
FROM watering_records w1
JOIN watering_records w2 ON w1.plant_id = w2.plant_id
WHERE w1.plant_id = '{plant_id}'
AND w1.watered_at > w2.watered_at
ORDER BY w1.watered_at DESC, w2.watered_at DESC
LIMIT 1;
```

### 前回の水やりからの経過日数
最新の水やり記録と現在の日付の差を計算
```sql
-- 例：植物IDが {plant_id} の前回の水やりからの経過日数を計算するクエリ
SELECT
  EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - watered_at)) / 86400 AS days_since_last_watering
FROM watering_records
WHERE plant_id = '{plant_id}'
ORDER BY watered_at DESC
LIMIT 1;
```
