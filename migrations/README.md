# go-migration
- インストール
  - `go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- マイグレーションファイルの生成
  - `migrate create -ext sql -dir migrations create_users_table`
- マイグレーションファイルの編集
  - up,down両方
- マイグレーションの実行方法
  - `migrate -database "mysql://user:password@tcp(host:port)/dbname" -path migrations up`
  - `migrate -database "mysql://root:password@tcp(localhost:13306)/mydb" -path DDD/migrations up`
- 特定のバージョンまでmigrateする場合
  - `migrate -database "mysql://user:password@tcp(host:port)/dbname" -path migrations goto バージョン`
- バージョンを指定してdownする場合
  - `migrate -database "mysql://user:password@tcp(host:port)/dbname" -path migrations down バージョン`


# mysqlクエリサンプル
```
-- 単一のレコードを挿入
INSERT INTO users (id, cognito_id, email, name) 
VALUES (
    UUID(), -- MySQL 8.0以降ではUUID()関数が使用可能
    'cognito_xxxxxxxxxxxxx',
    'test1@example.com',
    'テストユーザー1'
);

-- 複数レコードを一括挿入
INSERT INTO users (id, cognito_id, email, name) 
VALUES 
    (UUID(), 'cognito_111111111', 'user1@example.com', '山田太郎'),
    (UUID(), 'cognito_222222222', 'user2@example.com', '鈴木花子'),
    (UUID(), 'cognito_333333333', 'user3@example.com', '佐藤次郎');

-- テスト用データの作成（10件）
INSERT INTO users (id, cognito_id, email, name) 
VALUES 
    (UUID(), 'cognito_001', 'test1@example.com', 'テストユーザー1'),
    (UUID(), 'cognito_002', 'test2@example.com', 'テストユーザー2'),
    (UUID(), 'cognito_003', 'test3@example.com', 'テストユーザー3'),
    (UUID(), 'cognito_004', 'test4@example.com', 'テストユーザー4'),
    (UUID(), 'cognito_005', 'test5@example.com', 'テストユーザー5'),
    (UUID(), 'cognito_006', 'test6@example.com', 'テストユーザー6'),
    (UUID(), 'cognito_007', 'test7@example.com', 'テストユーザー7'),
    (UUID(), 'cognito_008', 'test8@example.com', 'テストユーザー8'),
    (UUID(), 'cognito_009', 'test9@example.com', 'テストユーザー9'),
    (UUID(), 'cognito_010', 'test10@example.com', 'テストユーザー10');
```